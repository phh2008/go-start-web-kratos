package data

import (
	"gorm.io/gorm"
	"helloword/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewRolePermissionRepo,
	NewRoleRepo,
	NewPermissionRepo,
)

// Data .
type Data struct {
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Bootstrap, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}
