package xcasbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloword/internal/conf"
	"path/filepath"
)

func NewCasbin(db *gorm.DB, c *conf.Bootstrap, logger log.Logger) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.NewHelper(logger).Errorf("casbin gorm 适配器创建失败,error:%s", err.Error())
		panic(err)
	}
	configFile := filepath.Join(c.Folder, "rbac_model.conf")
	rbacEnforcer, err := casbin.NewEnforcer(configFile, adapter)
	if err != nil {
		log.NewHelper(logger).Errorf("casbin.NewEnforcer 错误,error:%s", err.Error())
		panic(err)
	}
	rbacEnforcer.EnableAutoSave(true)
	// Load the policy from DB.
	_ = rbacEnforcer.LoadPolicy()
	return rbacEnforcer
}
