package mallagent

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

// EmbeddingConfig Embedding模型配置
type EmbeddingConfig struct {
	EndpointID string // ARK Embedding模型的Endpoint ID
	APIKey     string // ARK API Key
}

// NewEmbeddingModel 创建ARK Embedding模型实例
func NewEmbeddingModel(ctx context.Context, config *EmbeddingConfig) (*ark.EmbeddingModel, error) {
	if config == nil {
		config = &EmbeddingConfig{
			EndpointID: os.Getenv("ARK_EMBEDDING_MODEL"), // 使用环境变量中的ARK配置
			APIKey:     os.Getenv("ARK_API_KEY"),
		}
	}

	// 创建ARK EmbeddingModel
	embeddingModel, err := ark.NewEmbeddingModel(ctx, &ark.EmbeddingModelConfig{
		EndpointID: config.EndpointID,
		APIKey:     config.APIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("create ark embedding model failed: %v", err)
	}

	return embeddingModel, nil
}
