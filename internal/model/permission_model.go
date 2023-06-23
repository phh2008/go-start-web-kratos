package model

type PermissionModel struct {
	Id       int64  `json:"id"`                                      //权限ID
	PermName string `json:"permName" binding:"required"`             //权限名称
	Url      string `json:"url" binding:"required"`                  //URL路径
	Action   string `json:"action" binding:"required"`               //权限动作：比如get、post、delete等
	PermType uint8  `json:"permType" binding:"required,gte=1,lte=2"` //权限类型：1-菜单、2-按钮
	ParentId int64  `json:"parentId"`                                //父级ID：资源层级关系
}

type PermissionListReq struct {
	QueryPage
	PermName string `json:"permName" form:"permName"` //权限名称
	Url      string `json:"url" form:"url"`           //URL路径
	Action   string `json:"action" form:"action"`     //权限动作：比如get、post、delete等
	PermType uint8  `json:"permType" form:"permType"` //权限类型：1-菜单、2-按钮
}
