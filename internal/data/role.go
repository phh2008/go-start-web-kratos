package data

import (
    "context"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/jinzhu/copier"
    "gorm.io/gorm"
    "helloword/internal/biz"
    "helloword/internal/model"
    "helloword/pkg/orm"
)

type RoleEntity struct {
    BaseEntity
    RoleCode string // 角色编码
    RoleName string // 角色名称
}

func (RoleEntity) TableName() string {
    return "sys_role"
}

var _ biz.RoleRepo = (*roleRepo)(nil)

type roleRepo struct {
    *BaseRepo[RoleEntity]
}

// NewRoleRepo 创建 dao
func NewRoleRepo(db *gorm.DB) biz.RoleRepo {
    return &roleRepo{
        NewBaseRepo[RoleEntity](db),
    }
}

func (a *roleRepo) ListPage(ctx context.Context, req model.RoleListReq) (*model.PageData[*model.RoleModel], error) {
    db := a.getDb(ctx)
    db = db.Model(&RoleEntity{})
    if req.RoleCode != "" {
        db = db.Where("role_code like ?", "%"+req.RoleCode+"%")
    }
    if req.RoleName != "" {
        db = db.Where("role_name like ?", "%"+req.RoleName+"%")
    }
    pageData, db := orm.QueryPageData[*model.RoleModel](db, req.GetPageNo(), req.GetPageSize())
    err := db.Error
    if err != nil {
        return nil, err
    }
    return &pageData, nil
}

// Add 添加角色
func (a *roleRepo) Add(ctx context.Context, req model.RoleModel) (*model.RoleModel, error) {
    // 检查角色是否存在
    role, err := a.GetByCode(ctx, req.RoleCode)
    if err != nil {
        return nil, err
    }
    if role.Id > 0 {
        return nil, errors.New(500, "role_exist", "角色已存在")
    }
    var entity RoleEntity
    _ = copier.Copy(&entity, req)
    err = a.insert(ctx, &entity)
    if err != nil {
        return nil, err
    }
    _ = copier.Copy(&req, &entity)
    return &req, nil
}

func (a *roleRepo) GetById(ctx context.Context, id int64) (*model.RoleModel, error) {
    entity, err := a.getById(ctx, id)
    if err != nil {
        return nil, err
    }
    var role model.RoleModel
    _ = copier.Copy(&role, entity)
    return &role, nil
}

// GetByCode 根据角色编号获取角色
func (a *roleRepo) GetByCode(ctx context.Context, code string) (*model.RoleModel, error) {
    var entity RoleEntity
    err := a.getDb(ctx).Where("role_code=?", code).Limit(1).Find(&entity).Error
    if err != nil {
        return nil, err
    }
    var role model.RoleModel
    _ = copier.Copy(&role, &entity)
    return &role, nil
}

// DeleteById 删除角色
func (a *roleRepo) DeleteById(ctx context.Context, id int64) error {
    return a.deleteById(ctx, id)
}

// ListByIds 根据角色ID集合查询角色列表
func (a *roleRepo) ListByIds(ctx context.Context, ids []int64) ([]model.RoleModel, error) {
    list, err := a.listByIds(ctx, ids)
    if err != nil {
        return nil, err
    }
    var roleList []model.RoleModel
    _ = copier.Copy(&roleList, list)
    return roleList, nil
}
