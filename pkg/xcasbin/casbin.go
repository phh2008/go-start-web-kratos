package xcasbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"helloword/internal/conf"
	"helloword/pkg/logger"
	"path/filepath"
)

func NewCasbin(db *gorm.DB, c *conf.Bootstrap) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		logger.Errorf("casbin gorm 适配器创建失败,error:%s", err.Error())
		panic(err)
	}
	configFile := filepath.Join(c.Folder, "rbac_model.conf")
	rbacEnforcer, err := casbin.NewEnforcer(configFile, adapter)
	if err != nil {
		logger.Errorf("casbin.NewEnforcer 错误,error:%s", err.Error())
		panic(err)
	}
	rbacEnforcer.EnableAutoSave(true)
	// Load the policy from DB.
	_ = rbacEnforcer.LoadPolicy()
	return rbacEnforcer
}
