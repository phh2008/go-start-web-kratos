package biz

import (
	"context"
)

type RolePermissionEntity struct {
	Id     int64 // 主键id
	RoleId int64 // 角色id
	PermId int64 // 权限id
}

func (RolePermissionEntity) TableName() string {
	return "sys_role_permission"
}

type RolePermissionRepo interface {
	IBaseRepo[RolePermissionEntity]

	DeleteByRoleId(ctx context.Context, roleId int64) error
	BatchAdd(ctx context.Context, list []*RolePermissionEntity) error
	ListRoleIdByPermId(ctx context.Context, permId int64) []int64
}

type RolePermissionUseCase struct {
	rolePermissionRepo RolePermissionRepo
}

func NewRolePermissionUseCase(rolePermissionRepo RolePermissionRepo) *RolePermissionUseCase {
	return &RolePermissionUseCase{
		rolePermissionRepo: rolePermissionRepo,
	}
}
