package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"helloword/internal/biz"
	"helloword/internal/model"
	"helloword/pkg/orm"
)

type roleRepo struct {
	*BaseRepo[biz.RoleEntity]
}

// NewRoleRepo 创建 dao
func NewRoleRepo(db *gorm.DB) biz.RoleRepo {
	return &roleRepo{
		NewBaseRepo[biz.RoleEntity](db),
	}
}

func (a *roleRepo) ListPage(ctx context.Context, req model.RoleListReq) model.PageData[model.RoleModel] {
	db := a.GetDb(ctx)
	db = db.Model(&biz.UserEntity{})
	if req.RoleCode != "" {
		db = db.Where("role_code like ?", "%"+req.RoleCode+"%")
	}
	if req.RoleName != "" {
		db = db.Where("role_name like ?", "%"+req.RoleName+"%")
	}
	pageData, db := orm.QueryPageData[model.RoleModel](db, req.GetPageNo(), req.GetPageSize())
	return pageData
}

// Add 添加角色
func (a *roleRepo) Add(ctx context.Context, entity biz.RoleEntity) (biz.RoleEntity, error) {
	// 检查角色是否存在
	role := a.GetByCode(ctx, entity.RoleCode)
	if role.Id > 0 {
		return entity, errors.New(500, "role_exist", "角色已存在")
	}
	db := a.GetDb(ctx).Create(&entity)
	return entity, db.Error
}

// GetByCode 根据角色编号获取角色
func (a *roleRepo) GetByCode(ctx context.Context, code string) biz.RoleEntity {
	var role biz.RoleEntity
	a.GetDb(ctx).Where("role_code=?", code).First(&role)
	return role
}

// DeleteById 删除角色
func (a *roleRepo) DeleteById(ctx context.Context, id int64) error {
	ret := a.GetDb(ctx).Delete(&biz.RoleEntity{}, id)
	return ret.Error
}

// ListByIds 根据角色ID集合查询角色列表
func (a *roleRepo) ListByIds(ctx context.Context, ids []int64) ([]biz.RoleEntity, error) {
	var list []biz.RoleEntity
	db := a.GetDb(ctx).Find(&list, ids)
	return list, db.Error
}
