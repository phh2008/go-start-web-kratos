package data

import (
	"context"
	"gorm.io/gorm"
	"helloword/internal/biz"
	"helloword/internal/model"
	"helloword/pkg/orm"
)

type permissionRepo struct {
	*BaseRepo[biz.PermissionEntity]
}

// NewPermissionRepo 创建dao
func NewPermissionRepo(db *gorm.DB) biz.PermissionRepo {
	return &permissionRepo{
		NewBaseRepo[biz.PermissionEntity](db),
	}
}

func (a *permissionRepo) ListPage(ctx context.Context, req model.PermissionListReq) model.PageData[model.PermissionModel] {
	db := a.GetDb(ctx)
	db = db.Model(&biz.PermissionEntity{})
	if req.PermName != "" {
		db = db.Where("perm_name like ?", "%"+req.PermName+"%")
	}
	if req.Url != "" {
		db = db.Where("url=?", req.Url)
	}
	if req.Action != "" {
		db = db.Where("action=?", req.Action)
	}
	if req.PermType != 0 {
		db = db.Where("perm_type=?", req.PermType)
	}
	pageData, db := orm.QueryPageData[model.PermissionModel](db, req.GetPageNo(), req.GetPageSize())
	return pageData
}

func (a *permissionRepo) Add(ctx context.Context, permission biz.PermissionEntity) (biz.PermissionEntity, error) {
	db := a.GetDb(ctx).Create(&permission)
	return permission, db.Error
}

func (a *permissionRepo) Update(ctx context.Context, permission biz.PermissionEntity) (biz.PermissionEntity, error) {
	db := a.GetDb(ctx).Model(&permission).Updates(permission)
	return permission, db.Error
}

func (a *permissionRepo) FindByIdList(idList []int64) []biz.PermissionEntity {
	var list []biz.PermissionEntity
	if len(idList) == 0 {
		return list
	}
	db := a.GetDb(context.TODO())
	db.Find(&list, idList)
	return list
}
