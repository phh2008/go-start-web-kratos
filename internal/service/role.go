package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
	"helloword/internal/biz"
	"helloword/internal/model"

	pb "helloword/api/helloworld/v1"
)

type RoleService struct {
	pb.UnimplementedRoleServer
	roleUseCase *biz.RoleUseCase
}

func NewRoleService(roleUseCase *biz.RoleUseCase) *RoleService {
	return &RoleService{
		roleUseCase: roleUseCase,
	}
}

func (s *RoleService) ListPage(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListReply, error) {
	var request model.RoleListReq
	_ = copier.Copy(&request, &req)
	result := s.roleUseCase.ListPage(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.RoleListReply
	_ = copier.Copy(&reply, &result.Data)
	_ = copier.Copy(&reply.RoleList, &result.Data.Data)
	return &reply, nil
}

func (s *RoleService) Add(ctx context.Context, req *pb.RoleSaveRequest) (*pb.RoleReply, error) {
	var request model.RoleModel
	_ = copier.Copy(&request, &req)
	result := s.roleUseCase.Add(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.RoleReply
	_ = copier.Copy(&reply, &result.Data)
	return &reply, nil
}
func (s *RoleService) GetByCode(ctx context.Context, req *pb.RoleCodeRequest) (*pb.RoleReply, error) {
	result := s.roleUseCase.GetByCode(ctx, req.RoleCode)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.RoleReply
	_ = copier.Copy(&reply, &result.Data)
	return &reply, nil
}

func (s *RoleService) AssignPermission(ctx context.Context, req *pb.RoleAssignPermRequest) (*pb.RoleOk, error) {
	var request model.RoleAssignPermModel
	_ = copier.Copy(&request, &req)
	result := s.roleUseCase.AssignPermission(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	return &pb.RoleOk{Success: true}, nil
}

func (s *RoleService) DeleteById(ctx context.Context, req *pb.RoleDeleteRequest) (*pb.RoleOk, error) {
	result := s.roleUseCase.DeleteById(ctx, req.Id)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	return &pb.RoleOk{Success: true}, nil
}
