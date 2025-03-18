package biz

import (
    "context"
    "gorm.io/plugin/soft_delete"
)

func init() {
    soft_delete.FlagDeleted = 2
    soft_delete.FlagActived = 1
}

type IBaseRepo interface {
    // Transaction 开启事务
    Transaction(c context.Context, handler func(tx context.Context) error) error
}
