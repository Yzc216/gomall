package main

import (
	"github.com/Yzc216/gomall/app/inventory/biz/consumer"
	"github.com/Yzc216/gomall/app/inventory/biz/dal"
	"github.com/Yzc216/gomall/app/inventory/infra/mq"
	"github.com/Yzc216/gomall/common/mtl"
	"github.com/Yzc216/gomall/common/serversuite"
	"github.com/joho/godotenv"
	"net"
	"time"

	"github.com/Yzc216/gomall/app/inventory/conf"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/inventory/inventoryservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	//加载环境变量
	_ = godotenv.Load()
	//初始化指标
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	//初始化trace
	mtl.InitTracing(serviceName)

	dal.Init()
	//消息队列
	mq.Init()
	consumer.Init()

	opts := kitexInit()

	svr := inventoryservice.NewServer(new(InventoryServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr),
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0]},
		),
	)

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
