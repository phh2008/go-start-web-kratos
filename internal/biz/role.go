package biz

import (
    "context"
    "github.com/casbin/casbin/v2"
    "helloword/internal/model"
    "helloword/internal/model/result"
    "helloword/pkg/exception"
    "helloword/pkg/logger"
    "strconv"
)

type RoleRepo interface {
    IBaseRepo
    ListPage(ctx context.Context, req model.RoleListReq) (*model.PageData[*model.RoleModel], error)
    // Add 添加角色
    Add(ctx context.Context, entity model.RoleModel) (*model.RoleModel, error)
    // GetByCode 根据角色编号获取角色
    GetByCode(ctx context.Context, code string) (*model.RoleModel, error)
    // DeleteById 删除角色
    DeleteById(ctx context.Context, id int64) error
    GetById(ctx context.Context, id int64) (*model.RoleModel, error)
}

type RoleUseCase struct {
    roleRepo           RoleRepo
    rolePermissionRepo RolePermissionRepo
    permissionRepo     PermissionRepo
    enforcer           *casbin.Enforcer
    userRepo           UserRepo
}

func NewRoleUseCase(
    roleRepo RoleRepo,
    rolePermissionRepo RolePermissionRepo,
    permissionRepo PermissionRepo,
    enforcer *casbin.Enforcer,
    userRepo UserRepo,
) *RoleUseCase {
    return &RoleUseCase{
        roleRepo:           roleRepo,
        rolePermissionRepo: rolePermissionRepo,
        permissionRepo:     permissionRepo,
        enforcer:           enforcer,
        userRepo:           userRepo,
    }
}

// ListPage 角色列表
func (a *RoleUseCase) ListPage(ctx context.Context, req model.RoleListReq) *result.Result[*model.PageData[*model.RoleModel]] {
    data, err := a.roleRepo.ListPage(ctx, req)
    if err != nil {
        return result.Error[*model.PageData[*model.RoleModel]](err)
    }
    return result.Ok(data)
}

// Add 添加角色
func (a *RoleUseCase) Add(ctx context.Context, role model.RoleModel) *result.Result[*model.RoleModel] {
    resp, err := a.roleRepo.Add(ctx, role)
    if err != nil {
        return result.Error[*model.RoleModel](err)
    }
    return result.Ok(resp)
}

// GetByCode 根据角色编号获取角色
func (a *RoleUseCase) GetByCode(ctx context.Context, roleCode string) *result.Result[*model.RoleModel] {
    role, err := a.roleRepo.GetByCode(ctx, roleCode)
    if err != nil {
        return result.Error[*model.RoleModel](err)
    }
    return result.Ok(role)
}

// AssignPermission 分配权限
func (a *RoleUseCase) AssignPermission(ctx context.Context, assign model.RoleAssignPermModel) *result.Result[any] {
    // 获取角色
    role, err := a.roleRepo.GetById(ctx, assign.RoleId)
    if err != nil {
        return result.Error[any](err)
    }
    if role == nil || role.Id == 0 {
        return result.Failure[any]("角色不存在")
    }
    // 获取权限列表(权限列表为空表示要清空权限)
    permList, err := a.permissionRepo.FindByIdList(ctx, assign.PermIdList)
    if err != nil {
        return result.Error[any](err)
    }
    // 构建角色与权限关系对象
    var rolePermList []model.RolePermission // 角色权限关系表数据
    var permissions [][]string              // casbin 中的角色与权限数据
    for _, v := range permList {
        rolePermList = append(rolePermList, model.RolePermission{PermId: v.Id, RoleId: assign.RoleId})
        permissions = append(permissions, []string{role.RoleCode, v.Url, v.Action, strconv.FormatInt(v.Id, 10)})
    }
    err = a.rolePermissionRepo.Transaction(ctx, func(tx context.Context) error {
        // 删除原来的角色与权限关系
        err := a.rolePermissionRepo.DeleteByRoleId(tx, assign.RoleId)
        if err != nil {
            logger.Errorf("删除操作失败，%s", err.Error())
            return err
        }
        // 新增角色与权限关系
        err = a.rolePermissionRepo.BatchAdd(tx, rolePermList)
        if err != nil {
            logger.Errorf("删除操作失败，%s", err.Error())
            return err
        }
        return err
    })
    if err != nil {
        return result.Error[any](err)
    }
    // 更新casbin中的角色与资源关系
    _, _ = a.enforcer.DeletePermissionsForUser(role.RoleCode)
    if len(permissions) > 0 {
        _, _ = a.enforcer.AddPolicies(permissions)
    }
    return result.Success[any]()
}

// DeleteById 删除角色
func (a *RoleUseCase) DeleteById(ctx context.Context, id int64) *result.Result[any] {
    role, _ := a.roleRepo.GetById(ctx, id)
    if role.Id == 0 {
        return result.Success[any]()
    }
    err := a.roleRepo.Transaction(ctx, func(tx context.Context) error {
        // 清空用户表上的角色字段
        err := a.userRepo.CancelRole(ctx, role.RoleCode)
        if err != nil {
            return err
        }
        // 删除角色
        err = a.roleRepo.DeleteById(ctx, id)
        if err != nil {
            return err
        }
        // 删除角色与资源关系
        err = a.rolePermissionRepo.DeleteByRoleId(ctx, id)
        if err != nil {
            return err
        }
        // 删除 casbin 中的角色与资源关系、角色与用户关系
        _, err = a.enforcer.DeleteRole(role.RoleCode)
        return err
    })
    if err != nil {
        logger.Errorf("删除角色操作失败，%s", err)
        return result.Error[any](exception.SysError)
    }
    return result.Success[any]()
}
