package biz

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cristalhq/jwt/v5"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"helloword/internal/model"
	"helloword/internal/model/result"
	"helloword/pkg/exception"
	"helloword/pkg/logger"
	"helloword/pkg/xjwt"
	"strconv"
	"time"
)

type UserEntity struct {
	BaseEntity
	RealName string `json:"realName"`                // 姓名
	UserName string `json:"userName"`                // 用户名
	Email    string `json:"email"`                   // 邮箱
	Password string `json:"password"`                // 密码
	Status   int    `gorm:"default:1" json:"status"` //状态: 1-启用，2-禁用
	RoleCode string `json:"roleCode"`                // 角色编号
}

func (UserEntity) TableName() string {
	return "sys_user"
}

type UserRepo interface {
	IBaseRepo[UserEntity]

	ListPage(ctx context.Context, req model.UserListReq) model.PageData[model.UserModel]
	// GetByEmail 根据 email 查询
	GetByEmail(ctx context.Context, email string) UserEntity
	// Add 添加用户
	Add(ctx context.Context, user UserEntity) (UserEntity, error)
	// SetRole 设置角色
	SetRole(ctx context.Context, userId int64, role string) error
	// DeleteById 删除用户
	DeleteById(ctx context.Context, id int64) error
	// CancelRole 撤销用户角色
	CancelRole(ctx context.Context, roleCode string) error
}

// UserUseCase 用户业务逻辑封装
type UserUseCase struct {
	userRepo UserRepo
	jwt      *xjwt.JwtHelper
	enforcer *casbin.Enforcer
}

// NewUserUseCase 构造业务结构体
func NewUserUseCase(userRepo UserRepo, jwt *xjwt.JwtHelper, enforcer *casbin.Enforcer) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		jwt:      jwt,
		enforcer: enforcer,
	}
}

// ListPage 用户列表
func (a *UserUseCase) ListPage(ctx context.Context, req model.UserListReq) *result.Result[model.PageData[model.UserModel]] {
	data := a.userRepo.ListPage(ctx, req)
	return result.Ok[model.PageData[model.UserModel]](data)
}

// CreateByEmail 根据邮箱创建用户
func (a *UserUseCase) CreateByEmail(ctx context.Context, email model.UserEmailRegister) *result.Result[model.UserModel] {
	user := a.userRepo.GetByEmail(ctx, email.Email)
	if user.Id > 0 {
		return result.Failure[model.UserModel]("email 已存在")
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(email.Password), 1)
	if err != nil {
		logger.Errorf("生成密码出错：%s", err.Error())
		return result.Error[model.UserModel](err)
	}

	user = UserEntity{
		Email:    email.Email,
		RealName: email.Email,
		UserName: email.Email,
		Password: string(pwd),
		Status:   1,
		RoleCode: "",
	}
	user, err = a.userRepo.Add(ctx, user)
	if err != nil {
		logger.Errorf("创建用户出错：%s", err.Error())
		return result.Failure[model.UserModel]("创建用户出错")
	}
	var userModel model.UserModel
	copier.Copy(&userModel, &user)
	return result.Ok[model.UserModel](userModel)
}

// LoginByEmail 邮箱登录
func (a *UserUseCase) LoginByEmail(ctx context.Context, loginModel model.UserLoginModel) *result.Result[string] {
	user := a.userRepo.GetByEmail(ctx, loginModel.Email)
	if user.Id == 0 {
		return result.Failure[string]("用户或密码错误")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password))
	if err != nil {
		return result.Failure[string]("用户或密码错误")
	}
	// 生成token
	userClaims := xjwt.UserClaims{}
	userClaims.ID = strconv.FormatInt(user.Id, 10)
	userClaims.Role = user.RoleCode
	userClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7))
	token, err := a.jwt.CreateToken(userClaims)
	if err != nil {
		logger.Errorf("生成token错误：%s", err.Error())
		return result.Error[string](exception.SysError)
	}
	return result.Ok[string](token.String())
}

// AssignRole 给用户分配角色
func (a *UserUseCase) AssignRole(ctx context.Context, userRole model.AssignRoleModel) *result.Result[any] {
	err := a.userRepo.SetRole(ctx, userRole.UserId, userRole.RoleCode)
	if err != nil {
		logger.Errorf("db update error: %s", err.Error())
		return result.Failure[any]("分配角色出错")
	}
	// 更新casbin中的用户与角色关系
	uid := strconv.FormatInt(userRole.UserId, 10)
	_, _ = a.enforcer.DeleteRolesForUser(uid)
	// 角色为空，表示清除此用户的角色,无需添加
	if userRole.RoleCode != "" {
		_, _ = a.enforcer.AddGroupingPolicy(uid, userRole.RoleCode)
	}
	return result.Success[any]()
}

// DeleteById 根据ID删除
func (a *UserUseCase) DeleteById(ctx context.Context, id int64) *result.Result[any] {
	err := a.userRepo.DeleteById(ctx, id)
	if err != nil {
		logger.Errorf("delete error: %s", err.Error())
		return result.Failure[any]("刪除出错")
	}
	// 清除 casbin 中用户信息
	_, err = a.enforcer.DeleteRolesForUser(strconv.FormatInt(id, 10))
	if err != nil {
		logger.Errorf("Enforcer.DeleteRolesForUser error: %s", err)
	}
	return result.Success[any]()
}
