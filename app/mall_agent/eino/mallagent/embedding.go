package mallagent

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding"
	"github.com/cloudwego/eino-ext/components/model/ark"
)

// EmbeddingConfig Embedding模型配置
type EmbeddingConfig struct {
	EndpointID string // ARK Embedding模型的Endpoint ID
	APIKey     string // ARK API Key
	Model      string // 模型名称
}

// NewEmbeddingModel 创建ARK Embedding模型实例
func NewEmbeddingModel(ctx context.Context, config *EmbeddingConfig) (embedding.Model, error) {
	if config == nil {
		config = &EmbeddingConfig{
			Model: "doubao-embedding-large",
		}
	}

	return ark.NewEmbeddingModel(ctx, &ark.EmbeddingModelConfig{
		Model:      config.Model,
		MaxRetries: 3,
		Timeout:    time.Second * 10,
	})
}
