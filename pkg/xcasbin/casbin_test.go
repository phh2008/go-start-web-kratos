package xcasbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"helloword/internal/conf"
	"helloword/pkg/orm"
	"os"
	"testing"
)

var enforcer *casbin.Enforcer

func init() {
	var c = &conf.Bootstrap{}
	db := orm.NewDB(c.Data)
	enforcer = NewCasbin(db, c, log.NewStdLogger(os.Stdout))
}

func Test1(t *testing.T) {
	m := enforcer.GetModel()
	fmt.Println(m)
}

// 增加策略
func Test2(t *testing.T) {
	//if ok, _ := enforcer.AddPolicy("dev", "/api/v1/hello2", "GET"); !ok {
	//	fmt.Println("Policy已经存在")
	//} else {
	//	fmt.Println("增加成功")
	//}
	//批量添加 (要是有期中一条存在，就整批不会添加成功)
	//enforcer.AddPolicies()
	// 批量添加 (要是有期中一条存在，就整批不会添加成功)
	ok, err := enforcer.AddPermissionsForUser("xxx", []string{"/api/v1/hello5", "GET"}, []string{"/api/v1/hello6", "GET"})
	fmt.Println(ok, err)
}

// 删除策略
func Test3(t *testing.T) {
	fmt.Println("删除Policy")
	if ok, _ := enforcer.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
		fmt.Println("Policy不存在")
	} else {
		fmt.Println("删除成功")
	}
}

// 获取策略
func Test4(t *testing.T) {
	fmt.Println("查看policy")
	list := enforcer.GetPolicy()
	for _, vlist := range list {
		for _, v := range vlist {
			fmt.Printf("value: %s, ", v)
		}
		fmt.Printf("\n")
	}
}

// 更新
func Test5(t *testing.T) {
	enforcer.UpdatePolicy([]string{"admin", "/api/v1/test/query", "VIEW"}, []string{"admin", "/api/v1/test/query", "look"})
}

// 查询，条件过滤
func Test6(t *testing.T) {
	// 更新策略，根据条件(fieldIndex：表示参数开始字段索引)
	// 生成 sql如下，
	// DELETE FROM `casbin_rule` WHERE ptype = 'p' and v0 = 'bob' and v1 = 'data2'
	// INSERT INTO `casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`) VALUES ('p','admin','/api/v3/hello','POST','','','','','')
	enforcer.UpdateFilteredPolicies([][]string{{"admin", "/api/v3/hello", "POST"}}, 0, "bob", "data2")
}

// 给用户添加角色
func Test7(t *testing.T) {
	// 单个添加
	//ret, _ := enforcer.AddGroupingPolicy("tom", "admin")
	//ret, _ = enforcer.AddGroupingPolicy("lili", "admin")
	// 批量添加
	ret, _ := enforcer.AddGroupingPolicies([][]string{[]string{"1000", "root"}})
	//enforcer.SavePolicy()
	fmt.Println(ret)
}

// 验证是否具有权限（添加了用户与角色关系时，可以用角色或用户来检查）
func Test8(t *testing.T) {
	//ret, err := enforcer.Enforce("tom", "/api/v1/hello", "GET")
	ret, err := enforcer.Enforce("root", "/api/v1/test/:idx", "GET")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}

// 查询用户角色
func Test9(t *testing.T) {
	ret, _ := enforcer.GetRolesForUser("1000")
	fmt.Println(ret)
}

// 查询角色权限
func Test10(t *testing.T) {
	ret := enforcer.GetPermissionsForUser("root")
	fmt.Println(ret)
}

// 删除g，（可用于删除用户的角色）
func Test11(t *testing.T) {
	ok, err := enforcer.RemoveGroupingPolicy("role:area_ad_admin")
	fmt.Println(ok, err)
}

// 删除角色下的权限(根据角色删除所有权限，可以用来更新角色的权限)
// DELETE FROM `casbin_rule` WHERE ptype = 'p' and v0 = 'role:area_ad_admin'
func Test12(t *testing.T) {
	ok, err := enforcer.DeletePermissionsForUser("dev")
	fmt.Println(ok, err)
}

// 删除用户角色或权限
// DELETE FROM `casbin_rule` WHERE ptype = 'g' and v0 = 'guest'
// DELETE FROM `casbin_rule` WHERE ptype = 'p' and v0 = 'guest'
func Test13(t *testing.T) {
	ok, err := enforcer.DeleteUser("guest")
	fmt.Println(ok, err)
}

// 删除用户拥有的角色
func Test14(t *testing.T) {
	ok, err := enforcer.DeleteRolesForUser("lili")
	fmt.Println(ok, err)
}

// 更新策略
func Test15(t *testing.T) {
	// 先根据条件删除，再添加
	ok, err := enforcer.UpdateFilteredPolicies([][]string{[]string{"xxx", "/api/v1/hello8", "GET"}, []string{"xxx", "/api/v1/hello9", "PULL"}, []string{"xxx", "/api/v1/hello10", "GET"}}, 0, "xxx")
	fmt.Println(ok, err)
}

// 删除角色(删除角色关联用户，删除角色资源)
// DELETE FROM `casbin_rule` WHERE ptype = 'g' and v1 = 'xxx'
// DELETE FROM `casbin_rule` WHERE ptype = 'p' and v0 = 'xxx'
func Test16(t *testing.T) {
	ok, err := enforcer.DeleteRole("xxx")
	fmt.Println(ok, err)
}

func Test17(t *testing.T) {
	permList := enforcer.GetFilteredPolicy(3, "4")
	fmt.Println(permList)
}
