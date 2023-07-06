package biz

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"helloword/pkg/common"
	"helloword/pkg/localtime"
	"helloword/pkg/xjwt"
	"strconv"
)

func init() {
	soft_delete.FlagDeleted = 2
	soft_delete.FlagActived = 1
}

type IBaseRepo[T any] interface {
	// Transaction 开启事务
	Transaction(c context.Context, handler func(tx context.Context) error) error

	// GetDb 获取事务的db连接
	GetDb(ctx context.Context) *gorm.DB

	// GetById 根据ID查询
	GetById(ctx context.Context, id int64) (T, error)

	// Insert 新增
	Insert(ctx context.Context, entity T) (T, error)

	// Update 更新
	Update(ctx context.Context, entity T) (T, error)

	// DeleteById 根据ID删除
	DeleteById(ctx context.Context, id int64) error

	// ListByIds 根据ID集合查询
	ListByIds(ctx context.Context, ids []int64) ([]T, error)
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
