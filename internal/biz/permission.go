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

type PermissionEntity struct {
	BaseEntity
	PermName string // 权限名称
	Url      string // URL路径
	Action   string // 权限动作：比如get、post、delete等
	PermType uint8  `gorm:"default:1"` // 权限类型：1-菜单、2-按钮
	ParentId int64  `gorm:"default:0"` // 父级ID：资源层级关系（0表示顶级）
}

func (PermissionEntity) TableName() string {
	return "sys_permission"
}

type PermissionRepo interface {
	IBaseRepo[PermissionEntity]

	ListPage(ctx context.Context, req model.PermissionListReq) model.PageData[model.PermissionModel]
	Add(ctx context.Context, permission PermissionEntity) (PermissionEntity, error)
	Update(ctx context.Context, permission PermissionEntity) (PermissionEntity, error)
	FindByIdList(idList []int64) []PermissionEntity
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
func (a *PermissionUseCase) ListPage(ctx context.Context, req model.PermissionListReq) *result.Result[model.PageData[model.PermissionModel]] {
	data := a.permissionRepo.ListPage(ctx, req)
	return result.Ok[model.PageData[model.PermissionModel]](data)
}

// Add 添加权限资源
func (a *PermissionUseCase) Add(ctx context.Context, perm model.PermissionModel) *result.Result[PermissionEntity] {
	var permission PermissionEntity
	copier.Copy(&permission, &perm)
	res, err := a.permissionRepo.Add(ctx, permission)
	if err != nil {
		logger.Errorf("添加权限资源失败，%s", err.Error())
		return result.Failure[PermissionEntity]("添加权限资源失败")
	}
	return result.Ok(res)
}

// Update 更新权限资源
func (a *PermissionUseCase) Update(ctx context.Context, perm model.PermissionModel) *result.Result[*PermissionEntity] {
	oldPerm, _ := a.permissionRepo.GetById(ctx, perm.Id)
	if oldPerm.Id == 0 {
		return result.Error[*PermissionEntity](exception.NotFound)
	}
	var permission PermissionEntity
	copier.Copy(&permission, &perm)
	// 更新权限资源表
	res, err := a.permissionRepo.Update(ctx, permission)
	if err != nil {
		logger.Errorf("更新权限资源失败，%s", err.Error())
		return result.Failure[*PermissionEntity]("更新权限资源失败")
	}
	// 获取角色与资源列表,比如：[[systemAdmin /api/v1/user/list get] [guest /api/v1/user/list get]]
	perms := a.enforcer.GetFilteredPolicy(1, oldPerm.Url, oldPerm.Action)
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
	return result.Ok(&res)
}
