package pkg

import (
	"github.com/google/wire"
	"helloword/pkg/logger"
	"helloword/pkg/orm"
	"helloword/pkg/xcasbin"
	"helloword/pkg/xjwt"
)

var ProviderSet = wire.NewSet(
	orm.NewDB,
	xcasbin.NewCasbin,
	xjwt.NewJwtHelper,
	logger.NewLogger,
)
