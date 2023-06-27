package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
	"helloword/internal/biz"
	"helloword/internal/model"

	pb "helloword/api/helloworld/v1"
)

type PermissionService struct {
	pb.UnimplementedPermissionServer
	permissionUseCase *biz.PermissionUseCase
}

func NewPermissionService(permissionUseCase *biz.PermissionUseCase) *PermissionService {
	return &PermissionService{
		permissionUseCase: permissionUseCase,
	}
}

func (s *PermissionService) ListPage(ctx context.Context, req *pb.PermListRequest) (*pb.PermListReply, error) {
	var request model.PermissionListReq
	_ = copier.Copy(&request, &req)
	result := s.permissionUseCase.ListPage(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.PermListReply
	_ = copier.Copy(&reply, &result.Data)
	_ = copier.Copy(&reply.PermList, &result.Data.Data)
	return &reply, nil
}

func (s *PermissionService) Add(ctx context.Context, req *pb.PermSaveRequest) (*pb.PermReply, error) {
	var request model.PermissionModel
	_ = copier.Copy(&request, &req)
	result := s.permissionUseCase.Add(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.PermReply
	_ = copier.Copy(&reply, result.Data)
	return &reply, nil
}

func (s *PermissionService) Update(ctx context.Context, req *pb.PermSaveRequest) (*pb.PermReply, error) {
	var request model.PermissionModel
	_ = copier.Copy(&request, &req)
	result := s.permissionUseCase.Update(ctx, request)
	if !result.IsSuccess() {
		return nil, errors.New(500, result.Code, result.Message)
	}
	var reply pb.PermReply
	_ = copier.Copy(&reply, result.Data)
	return &reply, nil
}
