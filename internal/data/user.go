package data

import (
	"context"
	"gorm.io/gorm"
	"helloword/internal/biz"
	"helloword/internal/model"
	"helloword/pkg/orm"
)

type userRepo struct {
	*BaseRepo[biz.UserEntity]
}

// NewUserRepo 创建 userRepo
func NewUserRepo(db *gorm.DB) biz.UserRepo {
	return &userRepo{
		NewBaseRepo[biz.UserEntity](db),
	}
}

func (a *userRepo) ListPage(ctx context.Context, req model.UserListReq) model.PageData[model.UserModel] {
	db := a.GetDb(ctx)
	db = db.Model(&biz.UserEntity{})
	if req.RealName != "" {
		db = db.Where("real_name like ?", "%"+req.RealName+"%")
	}
	if req.Email != "" {
		db = db.Where("email like ?", "%"+req.Email+"%")
	}
	if req.Status != 0 {
		db = db.Where("status=?", req.Status)
	}
	var pageData model.PageData[model.UserModel]
	pageData.PageNo = req.GetPageNo()
	pageData.PageSize = req.GetPageSize()
	db.Count(&pageData.Count).
		Scopes(orm.OrderScope(req.Sort, req.Direction), orm.PageScope(req.GetPageNo(), req.GetPageSize())).
		Find(&pageData.Data)
	return pageData
}

// GetByEmail 根据 email 查询
func (a *userRepo) GetByEmail(ctx context.Context, email string) biz.UserEntity {
	var user biz.UserEntity
	a.GetDb(ctx).Model(biz.UserEntity{}).Where("email=?", email).First(&user)
	return user
}

// Add 添加用户
func (a *userRepo) Add(ctx context.Context, user biz.UserEntity) (biz.UserEntity, error) {
	ret := a.GetDb(ctx).Create(&user)
	return user, ret.Error
}

// SetRole 设置角色
func (a *userRepo) SetRole(ctx context.Context, userId int64, role string) error {
	db := a.GetDb(ctx).Model(&biz.UserEntity{}).Where("id=?", userId).Update("role_code", role)
	return db.Error
}

// DeleteById 删除用户
func (a *userRepo) DeleteById(ctx context.Context, id int64) error {
	db := a.GetDb(ctx).Delete(&biz.UserEntity{}, id)
	return db.Error
}

// CancelRole 撤销用户角色
func (a *userRepo) CancelRole(ctx context.Context, roleCode string) error {
	ret := a.GetDb(ctx).Model(&biz.UserEntity{}).Where("role_code=?", roleCode).Update("role_code", "")
	return ret.Error
}
