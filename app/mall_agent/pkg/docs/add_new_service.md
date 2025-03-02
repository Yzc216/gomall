# 如何添加新的微服务节点

本文档说明如何向mall_agent中添加新的微服务节点，使其能够被智能客服系统调用。

## 步骤概述

1. 定义新服务的接口
2. 实现新服务的RPC客户端
3. 创建服务工具
4. 将工具添加到Agent中

## 详细步骤

### 1. 定义服务接口

在`app/mall_agent/eino/mallagent/tools.go`文件中添加新服务的接口定义：

```go
// 示例：物流服务接口
type LogisticsService interface {
    TrackOrder(ctx context.Context, orderID string) (*LogisticsInfo, error)
    GetShippingOptions(ctx context.Context, address string) ([]ShippingOption, error)
    CalculateShippingFee(ctx context.Context, addressID string, weight float64) (float64, error)
}
```

同时在`model.go`中添加对应的数据结构：

```go
// 物流信息
type LogisticsInfo struct {
    OrderID      string    `json:"order_id"`
    TrackingID   string    `json:"tracking_id"`
    Status       string    `json:"status"`
    Carrier      string    `json:"carrier"`
    ShippedAt    time.Time `json:"shipped_at"`
    EstimatedArr time.Time `json:"estimated_arrival"`
    Locations    []TrackingLocation `json:"locations"`
}

// 物流跟踪位置
type TrackingLocation struct {
    Time        time.Time `json:"time"`
    Location    string    `json:"location"`
    Description string    `json:"description"`
}

// 配送选项
type ShippingOption struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Fee         float64 `json:"fee"`
    EstDays     int     `json:"estimated_days"`
}
```

### 2. 实现RPC客户端

在`app/mall_agent/eino/mallagent/client.go`中添加新服务客户端：

```go
// 1. 在ServiceClients结构体中添加字段
type ServiceClients struct {
    ProductClient   product.Client
    CartClient      cart.Client
    OrderClient     order.Client
    UserClient      user.Client
    LogisticsClient logistics.Client // 新增物流服务客户端
    once            sync.Once
}

// 2. 添加初始化函数
func initLogisticsClient(ctx context.Context, registry discovery.Registry) (logistics.Client, error) {
    return logistics.NewClient("logistics-service",
        client.WithRegistry(registry),
        client.WithRPCTimeout(3*time.Second))
}

// 3. 在InitClients中添加初始化语句
func InitClients(ctx context.Context, registry discovery.Registry) error {
    // ... 现有代码 ...
    
    if defaultClients.LogisticsClient, err = initLogisticsClient(ctx, registry); err != nil {
        return
    }
    
    // ... 其余代码 ...
}

// 4. 添加健康检查方法
func (s *ServiceClients) checkLogisticsHealth(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    
    // 实际应调用LogisticsClient的健康检查方法
    return nil
}

// 5. 在HealthCheck方法中添加服务检查
func (s *ServiceClients) HealthCheck(ctx context.Context) error {
    // ... 现有代码 ...
    
    // 检查Logistics服务
    if s.LogisticsClient != nil {
        if err := s.checkLogisticsHealth(ctx); err != nil {
            return fmt.Errorf("logistics service unhealthy: %v", err)
        }
    }
    
    return nil
}
```

### 3. 创建服务工具

在`app/mall_agent/eino/mallagent/tools.go`中添加新服务的工具实现：

```go
// 物流服务工具
func NewLogisticsTools(svc LogisticsService) []tool.BaseTool {
    return []tool.BaseTool{
        // 订单跟踪工具
        tool.NewTool(&schema.ToolInfo{
            Name:        "track_order",
            Description: "跟踪订单物流状态",
            Parameters: []*schema.Parameter{
                {Name: "order_id", Type: "string", Required: true, Description: "订单ID"},
            },
        }, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
            return svc.TrackOrder(ctx, args["order_id"].(string))
        }),

        // 查询配送选项工具
        tool.NewTool(&schema.ToolInfo{
            Name:        "get_shipping_options",
            Description: "获取可用的配送选项",
            Parameters: []*schema.Parameter{
                {Name: "address", Type: "string", Required: true, Description: "配送地址"},
            },
        }, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
            return svc.GetShippingOptions(ctx, args["address"].(string))
        }),

        // 计算运费工具
        tool.NewTool(&schema.ToolInfo{
            Name:        "calculate_shipping_fee",
            Description: "计算商品配送费用",
            Parameters: []*schema.Parameter{
                {Name: "address_id", Type: "string", Required: true, Description: "地址ID"},
                {Name: "weight", Type: "number", Required: true, Description: "商品重量(kg)"},
            },
        }, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
            weight := args["weight"].(float64)
            return svc.CalculateShippingFee(ctx, args["address_id"].(string), weight)
        }),
    }
}
```

### 4. 将工具添加到Agent

在`app/mall_agent/eino/mallagent/tools.go`中的GetTools函数中添加新服务的工具：

```go
func GetTools(ctx context.Context) ([]tool.BaseTool, error) {
    // ... 现有代码 ...
    
    // 添加物流服务工具
    if clients.LogisticsClient != nil {
        logisticsTools := NewLogisticsTools(clients.LogisticsClient)
        tools = append(tools, logisticsTools...)
    }
    
    return tools, nil
}
```

### 5. 更新环境配置

在`.env`文件中添加新服务的地址配置：

```
# 物流服务地址
LOGISTICS_SERVICE_ADDR=logistics-service:9005
```

## 验证新服务

添加完成后，可以重启mall_agent服务，并通过以下方式验证新服务是否正常工作：

1. 调用健康检查API验证服务连接是否正常
2. 通过客服界面测试新工具，如跟踪订单、查询配送选项等
3. 查看服务日志，确认工具调用过程没有错误

## 常见问题

1. **服务连接失败**：检查服务地址是否正确，确保目标服务已启动
2. **工具未显示**：检查GetTools函数中是否正确添加了新工具
3. **参数错误**：确保工具参数定义与实际API参数一致
4. **结果解析错误**：检查返回类型是否与模型定义一致 