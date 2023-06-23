package data

import (
	"context"
	"gorm.io/gorm"
	"helloword/internal/biz"
)

type rolePermissionRepo struct {
	BaseRepo[biz.RolePermissionEntity]
}

// NewRolePermissionRepo 创建 repo
func NewRolePermissionRepo(db *gorm.DB) biz.RolePermissionRepo {
	return &rolePermissionRepo{
		NewBaseRepo[biz.RolePermissionEntity](db),
	}
}

func (a *rolePermissionRepo) DeleteByRoleId(ctx context.Context, roleId int64) error {
	db := a.GetDb(ctx).Where("role_id=?", roleId).Delete(&biz.RolePermissionEntity{})
	return db.Error
}

func (a *rolePermissionRepo) BatchAdd(ctx context.Context, list []*biz.RolePermissionEntity) error {
	if len(list) == 0 {
		return nil
	}
	db := a.GetDb(ctx).Create(list)
	return db.Error
}

func (a *rolePermissionRepo) ListRoleIdByPermId(ctx context.Context, permId int64) []int64 {
	var roleIds []int64
	a.GetDb(ctx).Model(&biz.RolePermissionEntity{}).
		Where("perm_id=?", permId).
		Pluck("role_id", &roleIds)
	return roleIds
}
