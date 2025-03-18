package biz

import (
    "context"
    "github.com/casbin/casbin/v2"
    "github.com/jinzhu/copier"
    "helloword/internal/model"
    "helloword/internal/model/result"
    "helloword/pkg/exception"
    "helloword/pkg/logger"
)

type PermissionRepo interface {
    IBaseRepo
    ListPage(ctx context.Context, req model.PermissionListReq) (*model.PageData[*model.PermissionModel], error)
    Add(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error)
    Update(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error)
    FindByIdList(ctx context.Context, idList []int64) ([]model.PermissionModel, error)
    GetById(ctx context.Context, id int64) (*model.PermissionModel, error)
}

type PermissionUseCase struct {
    permissionRepo PermissionRepo
    enforcer       *casbin.Enforcer
}

func NewPermissionUseCase(permissionRepo PermissionRepo, enforcer *casbin.Enforcer) *PermissionUseCase {
    return &PermissionUseCase{
        permissionRepo: permissionRepo,
        enforcer:       enforcer,
    }
}

// ListPage 权限资源列表
func (a *PermissionUseCase) ListPage(ctx context.Context, req model.PermissionListReq) *result.Result[*model.PageData[*model.PermissionModel]] {
    data, err := a.permissionRepo.ListPage(ctx, req)
    if err != nil {
        return result.Error[*model.PageData[*model.PermissionModel]](err)
    }
    return result.Ok(data)
}

// Add 添加权限资源
func (a *PermissionUseCase) Add(ctx context.Context, perm model.PermissionModel) *result.Result[*model.PermissionModel] {
    var permission model.PermissionModel
    _ = copier.Copy(&permission, &perm)
    res, err := a.permissionRepo.Add(ctx, permission)
    if err != nil {
        logger.Errorf("添加权限资源失败，%s", err.Error())
        return result.Failure[*model.PermissionModel]("添加权限资源失败")
    }
    return result.Ok(res)
}

// Update 更新权限资源
func (a *PermissionUseCase) Update(ctx context.Context, perm model.PermissionModel) *result.Result[*model.PermissionModel] {
    oldPerm, err := a.permissionRepo.GetById(ctx, perm.Id)
    if err != nil {
        return result.Error[*model.PermissionModel](err)
    }
    if oldPerm.Id == 0 {
        return result.Error[*model.PermissionModel](exception.NotFound)
    }
    var permission model.PermissionModel
    _ = copier.Copy(&permission, &perm)
    // 更新权限资源表
    res, err := a.permissionRepo.Update(ctx, permission)
    if err != nil {
        return result.Error[*model.PermissionModel](err)
    }
    // 获取角色与资源列表,比如：[[systemAdmin /api/v1/user/list get] [guest /api/v1/user/list get]]
    perms, _ := a.enforcer.GetFilteredPolicy(1, oldPerm.Url, oldPerm.Action)
    // 更新casbin中的数据
    if len(perms) > 0 {
        for i, v := range perms {
            item := v
            item[1] = res.Url
            item[2] = res.Action
            perms[i] = item
        }
        _, err = a.enforcer.UpdateFilteredPolicies(perms, 1, oldPerm.Url, oldPerm.Action)
        if err != nil {
            logger.Errorf("更新casbin中的权限错误: %s", err)
        }
    }
    return result.Ok(res)
}
