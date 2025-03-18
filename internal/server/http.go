package server

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "helloword/api/helloworld/v1"
	"helloword/internal/conf"
	"helloword/internal/middleware"
	"helloword/internal/service"
	"helloword/pkg/logger"
	"helloword/pkg/xjwt"
	"strings"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cf *conf.Bootstrap,
	jwt *xjwt.JwtHelper,
	enforcer *casbin.Enforcer,
	userService *service.UserService,
	roleService *service.RoleService,
	permissionService *service.PermissionService,
	helloServer *service.HelloService,
) *http.Server {
	c := cf.Server
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(middleware.NewAuthenticate(jwt),
				middleware.NewAuthorization(enforcer)).
				Match(func(ctx context.Context, operation string) bool {
					logger.Infof("operation: %s", operation)
					if middleware.NoneAuthOperation.Contains(operation) {
						return false
					}
                    // TODO 测试，放行所有API
					return false
				}).
				Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	corsCfg := cf.Cors
	// 跨域配置
	corsOption := []handlers.CORSOption{
		handlers.AllowedOrigins(corsCfg.AllowedOriginPatterns),
		handlers.AllowedMethods(strings.Split(corsCfg.AllowedMethods, ",")),
		handlers.AllowedHeaders(strings.Split(corsCfg.AllowedHeaders, ",")),
		handlers.ExposedHeaders(strings.Split(corsCfg.ExposeHeaders, ",")),
		handlers.MaxAge(int(corsCfg.MaxAge)),
	}
	if corsCfg.AllowCredentials {
		corsOption = append(corsOption, handlers.AllowCredentials())
	}
	opts = append(opts, http.Filter(handlers.CORS(corsOption...)))
	opts = append(opts, http.PathPrefix("/api/v1"))
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, userService)
	v1.RegisterRoleHTTPServer(srv, roleService)
	v1.RegisterPermissionHTTPServer(srv, permissionService)
	v1.RegisterHelloHTTPServer(srv, helloServer)
	return srv
}
