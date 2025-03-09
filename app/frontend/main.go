// Code generated by hertz generator.

package main

import (
	"context"
	"encoding/json"
	frontendUtils "github.com/Yzc216/gomall/app/frontend/biz/utils"
	"github.com/Yzc216/gomall/app/frontend/infra/mtl"
	"github.com/Yzc216/gomall/app/frontend/infra/rpc"
	"github.com/Yzc216/gomall/app/frontend/middleware"
	frontendutils "github.com/Yzc216/gomall/app/frontend/utils"
	hertzprom "github.com/hertz-contrib/monitor-prometheus"
	hertzotelprovider "github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/joho/godotenv"
	oteltrace "go.opentelemetry.io/otel/trace"
	"html/template"
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
	hertzobslogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	hertzoteltracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ServiceName = frontendutils.ServiceName

var funcMap = template.FuncMap{
	"sub": func(a, b int) int { return a - b },
	"toJson": func(v interface{}) (template.JS, error) {
		b, err := json.Marshal(v)
		return template.JS(b), err
	},
}

func main() {

	_ = godotenv.Load()
	middleware.InitJWT()
	middleware.InitCasbin()

	mtl.InitMtl()
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
	// 初始化模板引擎
	t := template.New("")
	t.Funcs(funcMap) // 注册自定义函数

	// 加载模板文件（需要先注册函数再解析模板）
	t, err := t.ParseGlob("template/*.tmpl")
	if err != nil {
		panic(err)
	}
	// 设置模板引擎到 Hertz
	h.SetHTMLTemplate(t)

	//h.LoadHTMLGlob("template/*")
	//h.Delims("{{", "}}")
	h.Static("/static", "./")

	h.GET("/about", middleware.JwtOnlyParseMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		hlog.CtxInfof(c, "gomall about page")
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
	//store, _ := redis.NewStore(10, "tcp",
	//	conf.GetConf().Redis.Address,
	//	"",
	//	[]byte(os.Getenv("SESSION_SECRET")))
	//h.Use(sessions.New("gomall-shop", store))

	// log
	logger := hertzobslogrus.NewLogger(hertzobslogrus.WithLogger(hertzlogrus.NewLogger().Logger()))
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	var flushInterval time.Duration
	if os.Getenv("GO ENV") == "online" {
		flushInterval = time.Minute
	} else {
		flushInterval = time.Second
	}
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: flushInterval,
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

	h.OnShutdown = append(h.OnShutdown, mtl.Hooks...)

	// cores
	h.Use(cors.Default())
}
