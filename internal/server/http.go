package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "helloword/api/helloworld/v1"
	"helloword/internal/conf"
	"helloword/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cf *conf.Bootstrap,
	logger log.Logger,
	userService *service.UserService,
	roleService *service.RoleService,
	permissionService *service.PermissionService,
) *http.Server {
	c := cf.Server
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	opts = append(opts, http.PathPrefix("/api/v1"))
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, userService)
	v1.RegisterRoleHTTPServer(srv, roleService)
	v1.RegisterPermissionHTTPServer(srv, permissionService)
	return srv
}
