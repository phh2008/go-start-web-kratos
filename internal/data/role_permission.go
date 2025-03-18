package data

import (
    "context"
    "github.com/jinzhu/copier"
    "gorm.io/gorm"
    "helloword/internal/biz"
    "helloword/internal/model"
)

type RolePermissionEntity struct {
    Id     int64 // 主键id
    RoleId int64 // 角色id
    PermId int64 // 权限id
}

func (RolePermissionEntity) TableName() string {
    return "sys_role_permission"
}

var _ biz.RolePermissionRepo = (*rolePermissionRepo)(nil)

type rolePermissionRepo struct {
    *BaseRepo[RolePermissionEntity]
}

// NewRolePermissionRepo 创建 repo
func NewRolePermissionRepo(db *gorm.DB) biz.RolePermissionRepo {
    return &rolePermissionRepo{
        NewBaseRepo[RolePermissionEntity](db),
    }
}

func (a *rolePermissionRepo) DeleteByRoleId(ctx context.Context, roleId int64) error {
    db := a.getDb(ctx).Where("role_id=?", roleId).Delete(&RolePermissionEntity{})
    return db.Error
}

func (a *rolePermissionRepo) BatchAdd(ctx context.Context, list []model.RolePermission) error {
    if len(list) == 0 {
        return nil
    }
    var listEntity []RolePermissionEntity
    _ = copier.Copy(&listEntity, list)
    err := a.getDb(ctx).Create(&listEntity).Error
    return err
}

func (a *rolePermissionRepo) ListRoleIdByPermId(ctx context.Context, permId int64) []int64 {
    var roleIds []int64
    a.getDb(ctx).Model(&RolePermissionEntity{}).
        Where("perm_id=?", permId).
        Pluck("role_id", &roleIds)
    return roleIds
}
