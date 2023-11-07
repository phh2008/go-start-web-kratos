package middleware

import (
	"context"
	"github.com/casbin/casbin/v2"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	v1 "helloword/api/helloworld/v1"
	"helloword/pkg/common"
	"helloword/pkg/logger"
	"helloword/pkg/xjwt"
	"path/filepath"
	"strings"
	"time"
)

var NoneAuthOperation = mapset.NewSet[string](
	v1.User_LoginByEmail_FullMethodName,
	v1.Hello_HelloTest_FullMethodName,
)

// NewAuthenticate 认证校验
func NewAuthenticate(jwt *xjwt.JwtHelper) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, v1.ErrorNoLogin("未登录")
			}
			var user xjwt.UserClaims
			token := tr.RequestHeader().Get(common.AuthTokenKey)
			if token == "" {
				return nil, v1.ErrorNoLogin("未登录")
			}
			jwtToken, err := jwt.VerifyToken(token)
			if err != nil {
				return nil, v1.ErrorNoLogin("未登录")
			}
			user, err = jwt.ParseToken(jwtToken)
			if !user.IsValidExpiresAt(time.Now()) {
				return nil, v1.ErrorNoLogin("未登录")
			}
			ctx = context.WithValue(ctx, common.UserKey, user)
			return handler(ctx, req)
		}
	}
}

// NewAuthorization 权限校验
func NewAuthorization(enforcer *casbin.Enforcer) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			user, ok := ctx.Value(common.UserKey).(xjwt.UserClaims)
			role := user.Role // 当前用户角色
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, v1.ErrorUnauthorized("无权限")
			}
			// 获取接口包名与方法名称
			dir, file := filepath.Split(tr.Operation())
			obj := strings.TrimSuffix(dir, "/")
			act := file
			logger.Infof("权限校验 role:%s, obj:%s, act:%s", role, obj, act)
			ok, err = enforcer.Enforce(role, obj, act)
			if err != nil {
				logger.Errorf("Enforcer.Enforce error:%s", err.Error())
				// 鉴权出错了
				return nil, v1.ErrorSysError("鉴权出错了,请联系管理员")
			}
			if !ok {
				// 无权限
				return nil, v1.ErrorUnauthorized("无权限")
			}
			return handler(ctx, req)
		}
	}
}
