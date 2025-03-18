package biz

import (
    "context"
    "helloword/internal/model"
)

type RolePermissionRepo interface {
    IBaseRepo
    DeleteByRoleId(ctx context.Context, roleId int64) error
    BatchAdd(ctx context.Context, list []model.RolePermission) error
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
