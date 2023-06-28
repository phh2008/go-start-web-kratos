package server

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "helloword/api/helloworld/v1"
	"helloword/internal/conf"
	"helloword/internal/middleware"
	"helloword/internal/service"
	"helloword/pkg/logger"
	"helloword/pkg/xjwt"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	cf *conf.Bootstrap,
	jwt *xjwt.JwtHelper,
	enforcer *casbin.Enforcer,
	userService *service.UserService,
	roleService *service.RoleService,
	permissionService *service.PermissionService,
) *grpc.Server {
	c := cf.Server
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			selector.Server(middleware.NewAuthenticate(jwt),
				middleware.NewAuthorization(enforcer)).
				Match(func(ctx context.Context, operation string) bool {
					logger.Infof("operation: %s", operation)
					if middleware.NoneAuthOperation.Contains(operation) {
						return false
					}
					return true
				}).
				Build(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, userService)
	v1.RegisterRoleServer(srv, roleService)
	v1.RegisterPermissionServer(srv, permissionService)
	return srv
}
