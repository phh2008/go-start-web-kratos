package main

import (
	"flag"
	"go.uber.org/zap"
	"os"

	"helloword/internal/conf"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()

	active string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf ./config")
	flag.StringVar(&active, "active", "dev", "active environment, eg: -active dev")
}

func newApp(gs *grpc.Server, hs *http.Server, zapLog *zap.Logger) *kratos.App {
	// 包装 zap logger
	zlog := kratoszap.NewLogger(zapLog.WithOptions(zap.AddCallerSkip(3)))
	// 添加 traceId 等字段
	logger := log.With(zlog)//
	//"service.name", Name,
	//"service.version", Version,
	//"trace.id", tracing.TraceID(),
	//"span.id", tracing.SpanID(),

	log.SetLogger(logger)
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	var bc = conf.NewConfig(flagconf)
	conf.Active = active
	app, cleanup, err := wireApp(&bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
