// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/7qing/gomall/api/frontend/infra/rpc"
	frontendUtils "github.com/7qing/gomall/api/frontend/utils"
	"github.com/7qing/gomall/common/mtl"
	"github.com/hertz-contrib/casbin"
	prometheus "github.com/hertz-contrib/monitor-prometheus"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis" //   引进中间件session登录验证
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/7qing/gomall/api/frontend/biz/router"
	"github.com/7qing/gomall/api/frontend/conf"
	"github.com/7qing/gomall/api/frontend/middleware"
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
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName     = frontendUtils.ServiceName
	MetricsPort     = conf.GetConf().Hertz.MetricsPort
	RegistryAddress = conf.GetConf().Hertz.RegistryAddress[0]
)

func main() {
	_ = godotenv.Load() // 环境变量加载

	// init dal
	// dal.Init()
	consul, registerInfo := mtl.InitMetrics(ServiceName, MetricsPort, RegistryAddress)
	defer consul.Deregister(registerInfo)
	rpc.Init()
	address := conf.GetConf().Hertz.Address

	tracer, cfg := hertzoteltracing.NewServerTracer(hertzoteltracing.WithCustomResponseHandler(func(ctx context.Context, c *app.RequestContext) {
		c.Header("shop-trace-id", oteltrace.SpanFromContext(ctx).SpanContext().TraceID().String())
	}))

	h := server.New(server.WithHostPorts(address),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry),
		),
		),
		tracer,
	)
	h.Use(hertzoteltracing.ServerMiddleware(cfg))
	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)

	//加载资源
	h.LoadHTMLGlob("template/*")
	h.Delims("{{", "}}")
	h.Static("/static", "./")

	// 登录界面
	h.GET("sign-in", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-in", utils.H{
			"title": "Sign in",
			"next":  c.Query("next"),
		})
	})

	// 注册界面
	h.GET("sign-up", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "sign-up", utils.H{
			"title": "Sign up",
		})
	})

	h.GET("create-product", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "create_product", utils.H{
			"title": "create-product",
		})
	})

	h.GET("del-product", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "del_product", utils.H{
			"title": "del_product",
		})
	})
	// 重定向
	h.GET("/about", middleware.Casbinauth.RequiresRoles("admin", casbin.WithLogic(casbin.AND)), middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "about", utils.H{
			"title": "Error",
		})
	})

	h.GET("/about2", middleware.Casbinauth.RequiresRoles("admin", casbin.WithLogic(casbin.AND)), middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
		hlog.CtxInfof(ctx, "about page")
		c.HTML(consts.StatusOK, "about", utils.H{
			"title": "Error",
		})
	})

	// 添加中间件的方法
	//h.GET("/redirect", middleware.Auth(), func(ctx context.Context, c *app.RequestContext) {
	//	c.HTML(consts.StatusOK, "about", utils.H{
	//		"title": "Error",
	//	})
	//})
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret")) // 引进中间件session登录验证

	h.Use(sessions.New("cloudwego-shop", store))
	// log
	logger := hertzobslogrus.NewLogger(hertzobslogrus.WithLogger(hertzlogrus.NewLogger().Logger()))
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	var flushInterval time.Duration
	if os.Getenv("GO_ENV") == "online" {
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

	// cores
	h.Use(cors.Default())

	middleware.Register(h)

	middleware.CasbinRegister(h)
}
