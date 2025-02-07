// Code generated by hertz generator.

package main

import (
	"context"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/biz/utils"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/app/frontend/middleware"
	frontendutils "github.com/Yzc216/gomall/app/frontend/utils"
	"github.com/Yzc216/gomall/common/mtl"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	oteltrace "go.opentelemetry.io/otel/trace"
	"os"
	"time"

	"github.com/Yzc216/gomall/app/frontend/biz/router"
	"github.com/Yzc216/gomall/app/frontend/conf"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ServiceName = frontendutils.ServiceName

func main() {

	_ = godotenv.Load()

	consul, registryInfo := mtl.InitMetric(ServiceName, conf.GetConf().Hertz.MetricsPort, conf.GetConf().Hertz.RegistryAddr)
	defer consul.Deregister(registryInfo)

	mtl.InitTracing(ServiceName)
	//hertzmtl.InitMtl()
	rpc.InitClient()

	address := conf.GetConf().Hertz.Address

	p := hertzotelprovider.NewOpenTelemetryProvider(
		hertzotelprovider.WithSdkTracerProvider(mtl.TracerProvider),
		hertzotelprovider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := hertzoteltracing.NewServerTracer(hertzoteltracing.WithCustomResponseHandler(func(ctx context.Context, c *app.RequestContext) {
		c.Header("shop-trace-id", oteltrace.SpanFromContext(ctx).SpanContext().TraceID().String())
	}))

	h := server.New(server.WithHostPorts(address), server.WithTracer(hertzprom.NewServerTracer(
		"",
		"",
		hertzprom.WithDisableServer(true),
		hertzprom.WithRegistry(mtl.Registry),
	)), tracer)

	h.Use(hertzoteltracing.ServerMiddleware(cfg))
	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	h.Static("/static", "./")

	h.GET("/about", middleware.Auth(), func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about", frontendUtils.WarpResponse(c, ctx, utils.H{"title": "About"}))
	})

	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		data := utils.H{
			"title": "Sign In",
			"next":  ctx.Query("next"),
		}
		ctx.HTML(consts.StatusOK, "sign-in", data)
	})

	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up", utils.H{"title": "Sign Up"})
	})

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	//session
	store, _ := redis.NewStore(10, "tcp",
		conf.GetConf().Redis.Address,
		"",
		[]byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("gomall-shop", store))

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	middleware.Register(h)
}
