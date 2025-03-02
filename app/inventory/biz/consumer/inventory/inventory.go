package inventory

import (
	"context"
	"github.com/Yzc216/gomall/app/inventory/biz/dal/mysql"
	"github.com/Yzc216/gomall/app/inventory/biz/model"
	"github.com/Yzc216/gomall/app/inventory/infra/mq"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	tracer := otel.Tracer("shop-inventory-nats-consumer")
	sub, err := mq.Nc.Subscribe("inventory", func(m *nats.Msg) {
		var req inventory.ProductCreatedEvent
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
		}

		ctx := context.Background()
		inv := &model.Inventory{
			SkuID:     req.SkuId,
			Total:     req.InitialStock,
			Available: int32(req.InitialStock),
			Locked:    0,
			Version:   1,
		}
		err = model.InitStock(ctx, mysql.DB, inv)
		if err != nil {
			klog.Error(err)
		}

		// consumer otel
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(m.Header))
		_, span := tracer.Start(ctx, "inventory.init-stock")
		defer span.End()
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
