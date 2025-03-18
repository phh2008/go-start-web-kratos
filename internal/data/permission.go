package data

import (
    "context"
    "github.com/jinzhu/copier"
    "gorm.io/gorm"
    "helloword/internal/biz"
    "helloword/internal/model"
    "helloword/pkg/orm"
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

var _ biz.PermissionRepo = (*permissionRepo)(nil)

type permissionRepo struct {
    *BaseRepo[PermissionEntity]
}

// NewPermissionRepo 创建dao
func NewPermissionRepo(db *gorm.DB) biz.PermissionRepo {
    return &permissionRepo{
        NewBaseRepo[PermissionEntity](db),
    }
}

func (a *permissionRepo) ListPage(ctx context.Context, req model.PermissionListReq) (*model.PageData[*model.PermissionModel], error) {
    db := a.getDb(ctx)
    db = db.Model(&PermissionEntity{})
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
    pageData, db := orm.QueryPageData[*model.PermissionModel](db, req.GetPageNo(), req.GetPageSize())
    return &pageData, db.Error
}

func (a *permissionRepo) Add(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error) {
    var entity PermissionEntity
    _ = copier.Copy(&entity, permission)
    err := a.insert(ctx, &entity)
    if err != nil {
        return nil, err
    }
    _ = copier.Copy(&permission, &entity)
    return &permission, nil
}

func (a *permissionRepo) Update(ctx context.Context, permission model.PermissionModel) (*model.PermissionModel, error) {
    var entity PermissionEntity
    _ = copier.Copy(&entity, &permission)
    err := a.update(ctx, &entity)
    if err != nil {
        return nil, err
    }
    _ = copier.Copy(&permission, &entity)
    return &permission, nil
}

func (a *permissionRepo) FindByIdList(ctx context.Context, idList []int64) ([]model.PermissionModel, error) {
    if len(idList) == 0 {
        return nil, nil
    }
    list, err := a.listByIds(ctx, idList)
    if err != nil {
        return nil, err
    }
    var result []model.PermissionModel
    _ = copier.Copy(&result, &list)
    return result, nil
}

func (a *permissionRepo) GetById(ctx context.Context, id int64) (*model.PermissionModel, error) {
    entity, err := a.getById(ctx, id)
    if err != nil {
        return nil, err
    }
    var permission model.PermissionModel
    _ = copier.Copy(&permission, entity)
    return &permission, nil
}
