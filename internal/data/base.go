package data

import (
    "context"
    "gorm.io/gorm"
    "gorm.io/plugin/soft_delete"
    "helloword/internal/biz"
    "helloword/pkg/common"
    "helloword/pkg/localtime"
    "helloword/pkg/xjwt"
    "strconv"
)

var _ = biz.IBaseRepo(new(BaseRepo[any]))

var dbKey = "txDb"

type BaseRepo[T any] struct {
    database *gorm.DB
}

// NewBaseRepo 创建 baseRepo
func NewBaseRepo[T any](db *gorm.DB) *BaseRepo[T] {
    return &BaseRepo[T]{database: db}
}

// Transaction 开启事务
func (a *BaseRepo[T]) Transaction(c context.Context, handler func(tx context.Context) error) error {
    db := a.database
    return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
        return handler(context.WithValue(c, dbKey, tx))
    })
}

// GetDb 获取事务的db连接
func (a *BaseRepo[T]) getDb(ctx context.Context) *gorm.DB {
    db, ok := ctx.Value(dbKey).(*gorm.DB)
    if !ok {
        db = a.database
        return db.WithContext(ctx)
    }
    return db
}

// GetById 根据ID查询
func (a *BaseRepo[T]) getById(ctx context.Context, id int64) (*T, error) {
    var domain T
    err := a.getDb(ctx).Limit(1).Find(&domain, id).Error
    return &domain, err
}

// Insert 新增
func (a *BaseRepo[T]) insert(ctx context.Context, entity *T) error {
    return a.getDb(ctx).Create(entity).Error
}

// Update 更新
func (a *BaseRepo[T]) update(ctx context.Context, entity *T) error {
    return a.getDb(ctx).Model(entity).Updates(*entity).Error
}

// DeleteById 根据ID删除
func (a *BaseRepo[T]) deleteById(ctx context.Context, id int64) error {
    return a.getDb(ctx).Delete(new(T), id).Error
}

// ListByIds 根据ID集合查询
func (a *BaseRepo[T]) listByIds(ctx context.Context, ids []int64) ([]T, error) {
    var list []T
    db := a.getDb(ctx).Find(&list, ids)
    return list, db.Error
}

type BaseEntity struct {
    Id       int64                 `gorm:"primaryKey" json:"id"`                                             // 主键id
    CreateAt localtime.LocalTime   `gorm:"autoCreateTime" json:"createAt"`                                   // 创建时间
    UpdateAt localtime.LocalTime   `gorm:"autoUpdateTime" json:"updateAt"`                                   // 更新时间
    CreateBy int64                 `json:"createBy"`                                                         // 创建人
    UpdateBy int64                 `json:"updateBy"`                                                         // 更新人
    Deleted  soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:UpdateAt,default:1" json:"deleted"` // 是否删除 1-否，2-是
}

func (a *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
    a.Deleted = 1
    ctx := tx.Statement.Context
    user, ok := ctx.Value(common.UserKey).(xjwt.UserClaims)
    if ok {
        a.CreateBy, _ = strconv.ParseInt(user.ID, 10, 64)
        a.UpdateBy = a.CreateBy
    }
    return
}

func (a *BaseEntity) BeforeUpdate(tx *gorm.DB) (err error) {
    ctx := tx.Statement.Context
    user, ok := ctx.Value(common.UserKey).(xjwt.UserClaims)
    if ok {
        a.UpdateBy, _ = strconv.ParseInt(user.ID, 10, 64)
    }
    return
}

func (a *BaseEntity) BeforeDelete(tx *gorm.DB) (err error) {
    ctx := tx.Statement.Context
    user, ok := ctx.Value(common.UserKey).(xjwt.UserClaims)
    if ok {
        a.UpdateBy, _ = strconv.ParseInt(user.ID, 10, 64)
    }
    return
}
