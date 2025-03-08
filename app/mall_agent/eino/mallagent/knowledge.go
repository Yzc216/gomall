package mallagent

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/components/retriever/redis"
	"github.com/cloudwego/eino-ext/components/retriever/volc_vikingdb"
	"github.com/cloudwego/eino/schema"
)

// KnowledgeConfig 知识库配置
type KnowledgeConfig struct {
	RedisAddr    string
	IndexName    string
	TopK         int
	MinScore     float32
	EmbeddingDim int
	Endpoint     string
}

// NewKnowledgeRetriever 创建知识库检索器
func NewKnowledgeRetriever(ctx context.Context, config *KnowledgeConfig) (retriever.Retriever, error) {
	if config == nil {
		config = &KnowledgeConfig{
			RedisAddr:    os.Getenv("REDIS_ADDR"),
			IndexName:    "mall_knowledge",
			TopK:         5,
			MinScore:     0.7,
			EmbeddingDim: 1536,
			Endpoint:     os.Getenv("VIKINGDB_ENDPOINT"),
		}
	}

	// 创建Embedding模型
	embeddingModel, err := NewEmbeddingModel(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 创建Redis向量检索器
	redisRetriever, err := redis.NewRedisRetriever(ctx, &redis.Config{
		Addr:         config.RedisAddr,
		IndexName:    config.IndexName,
		TopK:         config.TopK,
		MinScore:     config.MinScore,
		EmbeddingDim: config.EmbeddingDim,
		Model:        embeddingModel,
	})
	if err != nil {
		return nil, fmt.Errorf("create redis retriever failed: %v", err)
	}

	return vikingdb.NewRetriever(ctx, &vikingdb.Config{
		Endpoint: config.Endpoint,
		TopK:     config.TopK,
		MinScore: config.MinScore,
	})
}

// IndexKnowledge 索引知识到Redis
func IndexKnowledge(ctx context.Context, docs []*schema.Document) error {
	// 创建Redis向量检索器
	retriever, err := NewKnowledgeRetriever(ctx, nil)
	if err != nil {
		return err
	}

	// 索引知识
	redisRetriever, ok := retriever.(*redis.RedisRetriever)
	if !ok {
		return fmt.Errorf("retriever is not redis retriever")
	}

	// 使用Redis检索器索引文档
	err = redisRetriever.Index(ctx, docs)
	if err != nil {
		return fmt.Errorf("index documents failed: %v", err)
	}

	return nil
}

// LoadAndIndexKnowledge 加载并索引电商知识库
func LoadAndIndexKnowledge(ctx context.Context) error {
	// 准备知识库文档
	docs := []*schema.Document{
		{ID: "faq-1", Content: "如何查询订单状态？您可以在"我的订单"页面查看所有订单状态，或者告诉客服您的订单号，我们可以帮您查询。"},
		{ID: "faq-2", Content: "如何申请退款？在订单详情页面，点击"申请退款"按钮，填写退款原因并提交申请。客服会在24小时内处理您的请求。"},
		{ID: "faq-3", Content: "忘记密码怎么办？请点击登录页面的"忘记密码"链接，通过绑定的手机号或邮箱进行身份验证后重置密码。"},
		{ID: "faq-4", Content: "如何修改收货地址？登录后进入"个人中心"，点击"地址管理"，您可以添加、编辑或删除收货地址。"},
		{ID: "policy-1", Content: "退换货政策：商品支持7天无理由退换，部分特殊商品（如鲜花、定制商品）除外。退换商品需保持原包装完好，不影响二次销售。"},
		{ID: "policy-2", Content: "配送说明：订单确认后通常1-3个工作日内发货，偏远地区可能需要更长时间。支持快递配送和到店自提两种方式。"},
		{ID: "policy-3", Content: "支付方式：支持微信支付、支付宝、银行卡支付和货到付款等多种支付方式。部分地区和特殊商品可能不支持货到付款。"},
		{ID: "product-1", Content: "商品保修服务：电子产品类商品享有一年保修服务，自收到商品之日起计算。保修期内出现非人为损坏，可联系客服安排维修或更换。"},
		{ID: "product-2", Content: "正品保证：本平台所有商品均为正品，如发现假冒商品，可获得双倍赔偿。"},
		{ID: "membership-1", Content: "会员等级：普通会员、银卡会员、金卡会员和钻石会员。不同等级会员享有不同的优惠和服务。"},
		{ID: "membership-2", Content: "积分规则：消费1元积1分，积分可用于抵扣订单金额或兑换礼品。积分有效期为一年，请及时使用。"},
	}

	// 索引知识库
	if err := IndexKnowledge(ctx, docs); err != nil {
		return fmt.Errorf("索引知识库失败: %v", err)
	}

	return nil
}

// SearchKnowledge 搜索知识库
func SearchKnowledge(ctx context.Context, query string) ([]*schema.Document, error) {
	// 创建检索器
	retriever, err := NewKnowledgeRetriever(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 搜索相关文档
	docs, err := retriever.Retrieve(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("检索知识库失败: %v", err)
	}

	return docs, nil
}
