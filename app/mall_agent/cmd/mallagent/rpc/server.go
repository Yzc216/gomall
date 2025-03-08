package rpc

import (
	"context"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/strings77wzq/gomall/app/mall_agent/config"
	"github.com/strings77wzq/gomall/app/mall_agent/discovery"
	"github.com/strings77wzq/gomall/app/mall_agent/eino/mallagent"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/mall_agent"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/mall_agent/mallagentservice"
)

// MallAgentServiceImpl 实现mall_agent服务接口
type MallAgentServiceImpl struct {
	agentManager *mallagent.AgentManager
}

// NewMallAgentServiceImpl 创建mall_agent服务实现
func NewMallAgentServiceImpl() *MallAgentServiceImpl {
	return &MallAgentServiceImpl{
		agentManager: mallagent.GetAgentManager(),
	}
}

// ProcessQuery 处理用户查询
func (s *MallAgentServiceImpl) ProcessQuery(ctx context.Context, req *mall_agent.QueryRequest) (*mall_agent.QueryResponse, error) {
	log.Printf("[RPC] 收到查询请求: %s, 会话ID: %s", req.Query, req.SessionId)
	
	// 调用Agent处理查询
	result, err := s.agentManager.ProcessQuery(ctx, req.SessionId, req.Query, req.UserId)
	if err != nil {
		log.Printf("[RPC] 处理查询失败: %v", err)
		return &mall_agent.QueryResponse{
			Success: false,
			Message: "处理查询失败: " + err.Error(),
		}, nil
	}
	
	// 构建响应
	actions := make([]*mall_agent.Action, 0)
	for _, act := range result.Actions {
		actions = append(actions, &mall_agent.Action{
			ActionType: act.Type,
			TargetId:   act.TargetID,
			Result:     act.Result,
		})
	}
	
	return &mall_agent.QueryResponse{
		Success: true,
		Message: result.Message,
		Actions: actions,
	}, nil
}

// ProcessQueryStream 流式处理用户查询
func (s *MallAgentServiceImpl) ProcessQueryStream(req *mall_agent.QueryRequest, stream mall_agent.MallAgentService_ProcessQueryStreamServer) error {
	log.Printf("[RPC] 收到流式查询请求: %s, 会话ID: %s", req.Query, req.SessionId)
	
	// 创建结果通道
	resultCh := make(chan mallagent.StreamChunk, 10)
	
	// 在后台处理查询
	go func() {
		defer close(resultCh)
		err := s.agentManager.ProcessQueryStream(stream.Context(), req.SessionId, req.Query, req.UserId, resultCh)
		if err != nil {
			log.Printf("[RPC] 流式处理查询失败: %v", err)
			resultCh <- mallagent.StreamChunk{
				Content: "处理查询失败: " + err.Error(),
				IsFinal: true,
			}
		}
	}()
	
	// 发送流式响应
	for chunk := range resultCh {
		if err := stream.Send(&mall_agent.QueryResponseChunk{
			Content:    chunk.Content,
			IsFinal:    chunk.IsFinal,
			ActionType: chunk.ActionType,
		}); err != nil {
			log.Printf("[RPC] 发送流式响应失败: %v", err)
			return err
		}
	}
	
	return nil
}

// GetSessionHistory 获取会话历史
func (s *MallAgentServiceImpl) GetSessionHistory(ctx context.Context, req *mall_agent.HistoryRequest) (*mall_agent.HistoryResponse, error) {
	log.Printf("[RPC] 获取会话历史: %s", req.SessionId)
	
	// 获取会话历史
	history, err := s.agentManager.GetSessionHistory(ctx, req.SessionId, int(req.Limit))
	if err != nil {
		log.Printf("[RPC] 获取会话历史失败: %v", err)
		return &mall_agent.HistoryResponse{}, nil
	}
	
	// 构建响应
	messages := make([]*mall_agent.Message, 0, len(history))
	for _, msg := range history {
		messages = append(messages, &mall_agent.Message{
			Role:      msg.Role,
			Content:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}
	
	return &mall_agent.HistoryResponse{
		Messages: messages,
	}, nil
}

// HealthCheck 健康检查
func (s *MallAgentServiceImpl) HealthCheck(ctx context.Context, req *mall_agent.HealthCheckRequest) (*mall_agent.HealthCheckResponse, error) {
	log.Printf("[RPC] 健康检查请求")
	
	status := true
	message := "服务正常"
	dependencies := make(map[string]*mall_agent.ServiceStatus)
	
	// 如果需要检查依赖服务
	if req.CheckDependencies {
		// 检查各个微服务的健康状态
		clients := mallagent.GetClients()
		if clients != nil {
			depStatus := clients.HealthCheck(ctx)
			for svc, st := range depStatus {
				dependencies[svc] = &mall_agent.ServiceStatus{
					Available: st.Available,
					Message:   st.Message,
				}
				
				if !st.Available {
					status = false
				}
			}
		}
	}
	
	return &mall_agent.HealthCheckResponse{
		Status:       status,
		Message:      message,
		Dependencies: dependencies,
	}, nil
}

// StartRPCServer 启动RPC服务器
func StartRPCServer(cfg *config.Config, registry discovery.Registry) error {
	// 创建服务实现
	impl := NewMallAgentServiceImpl()
	
	// 服务地址
	addr, err := net.ResolveTCPAddr("tcp", cfg.RPCServer.Address)
	if err != nil {
		return err
	}
	
	// 创建服务器
	svr := mallagentservice.NewServer(
		impl,
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "mall-agent-service",
		}),
	)
	
	// 启动服务器
	log.Printf("mall_agent RPC服务启动，监听地址 %s", cfg.RPCServer.Address)
	return svr.Run()
}