package data

import (
	"context"
	"gorm.io/gorm"
)

var dbKey = "txDb"

type BaseRepo[T any] interface {
	GetDb(ctx context.Context) *gorm.DB
	Transaction(c context.Context, handler func(tx context.Context) error) error
	GetById(ctx context.Context, id int64) (T, error)
}

type baseRepo[T any] struct {
	db *gorm.DB
}

// NewBaseRepo 创建 baseRepo
func NewBaseRepo[T any](db *gorm.DB) BaseRepo[T] {
	return &baseRepo[T]{db: db}
}

// Transaction 开启事务
func (a *baseRepo[T]) Transaction(c context.Context, handler func(tx context.Context) error) error {
	db := a.db
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		return handler(context.WithValue(c, dbKey, tx))
	})
}

// GetDb 获取事务的db连接
func (a *baseRepo[T]) GetDb(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		db = a.db
		return db.WithContext(ctx)
	}
	return db
}

// GetById 根据ID查询
func (a *baseRepo[T]) GetById(ctx context.Context, id int64) (T, error) {
	var domain T
	err := a.GetDb(ctx).Limit(1).Find(&domain, id).Error
	return domain, err
}
