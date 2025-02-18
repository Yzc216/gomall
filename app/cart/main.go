package main

import (
	"github.com/Yzc216/gomall/app/cart/biz/dal"
	"github.com/Yzc216/gomall/app/cart/infra/rpc"
	"github.com/Yzc216/gomall/common/mtl"
	"github.com/Yzc216/gomall/common/serversuite"
	"github.com/joho/godotenv"
	"net"
	"time"

	"github.com/Yzc216/gomall/app/cart/conf"
	"github.com/Yzc216/gomall/rpc_gen/kitex_gen/cart/cartservice"
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

	//rpc客户端初始化
	rpc.InitClient()

	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

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
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0]}))

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
