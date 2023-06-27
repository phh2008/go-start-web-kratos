package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
	"helloword/internal/biz"
	"helloword/internal/model"

	pb "helloword/api/helloworld/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	userUseCase *biz.UserUseCase
}

func NewUserService(userUseCase *biz.UserUseCase) *UserService {
	return &UserService{
		userUseCase: userUseCase,
	}
}

func (s *UserService) ListPage(ctx context.Context, req *pb.UserListRequest) (*pb.UserListReply, error) {
	var request model.UserListReq
	_ = copier.Copy(&request, &req)
	result := s.userUseCase.ListPage(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.UserListReply
	_ = copier.Copy(&reply, &result.Data)
	_ = copier.Copy(&reply.UserList, &result.Data.Data)
	return &reply, nil
}

func (s *UserService) CreateByEmail(ctx context.Context, req *pb.UserEmailRequest) (*pb.UserReply, error) {
	var request model.UserEmailRegister
	_ = copier.Copy(&request, &req)
	result := s.userUseCase.CreateByEmail(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.UserReply
	_ = copier.Copy(&reply, &result.Data)
	return &reply, nil
}

func (s *UserService) LoginByEmail(ctx context.Context, req *pb.UserEmailLoginRequest) (*pb.LoginReply, error) {
	var request model.UserLoginModel
	_ = copier.Copy(&request, &req)
	result := s.userUseCase.LoginByEmail(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	return &pb.LoginReply{Token: result.Data}, nil
}

func (s *UserService) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.UserOk, error) {
	var request model.AssignRoleModel
	_ = copier.Copy(&request, &req)
	result := s.userUseCase.AssignRole(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	return &pb.UserOk{Success: true}, nil
}

func (s *UserService) DeleteById(ctx context.Context, req *pb.UserDeleteRequest) (*pb.UserOk, error) {
	result := s.userUseCase.DeleteById(ctx, req.Id)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	return &pb.UserOk{Success: true}, nil
}
