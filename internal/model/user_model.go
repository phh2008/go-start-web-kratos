package model

import (
	"helloword/pkg/localtime"
)

type UserModel struct {
	Id       int64               `json:"id"`       // 主键id
	RealName string              `json:"realName"` // 姓名
	UserName string              `json:"userName"` // 用户名
	Email    string              `json:"email"`    // 邮箱
	Status   int                 `json:"status"`   // 状态: 1-启用，2-禁用
	RoleCode string              `json:"roleCode"` // 角色编号
	CreateAt localtime.LocalTime `json:"createAt"` // 创建时间
	UpdateAt localtime.LocalTime `json:"updateAt"` // 更新时间
	CreateBy int64               `json:"createBy"` // 创建人
	UpdateBy int64               `json:"updateBy"` // 更新人
	Deleted  uint8               `json:"deleted"`  // 是否删除 1-否，2-是
}

type UserListReq struct {
	QueryPage
	RealName string `json:"realName" form:"realName"` // 姓名
	Email    string `json:"email" form:"email"`       // 邮箱
	Status   int    `json:"status" form:"status"`     // 状态: 1-启用，2-禁用
}

type UserEmailRegister struct {
	Email    string `json:"email" binding:"required"`    // 邮箱
	Password string `json:"password" binding:"required"` // 密码
}

type UserLoginModel struct {
	Email    string `json:"email" binding:"required"`    // 邮箱
	Password string `json:"password" binding:"required"` // 密码
}

type AssignRoleModel struct {
	UserId   int64  `json:"userId" binding:"required"`   // 用户ID
	RoleCode string `json:"roleCode" binding:"required"` // 角色编码
}
