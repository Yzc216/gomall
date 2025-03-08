package mallagent

import (
	"context"

	"github.com/cloudwego/eino-ext/components/retriever/volc_vikingdb"
	"github.com/cloudwego/eino/schema"
)

func NewRetriever(ctx context.Context, config *RetrieverConfig) (retriever.Retriever, error) {
	if config == nil {
		config = defaultRetrieverConfig()
	}

	return vikingdb.NewRetriever(ctx, &vikingdb.Config{
		Endpoint:  config.Endpoint,
		TopK:      config.TopK,
		MinScore:  config.MinScore,
		IndexName: "mall_knowledge",
	})
}

func RetrieveKnowledge(ctx context.Context, query string) ([]*schema.Document, error) {
	retriever, err := NewRetriever(ctx, nil)
	if err != nil {
		return nil, err
	}

	return retriever.Retrieve(ctx, query)
}
