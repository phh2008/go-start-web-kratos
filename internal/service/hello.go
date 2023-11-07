package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "helloword/api/helloworld/v1"
	"helloword/pkg/logger"
	"time"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (a *HelloService) HelloTest(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	logger.Infof("req param: %s", req.GetName())
	date := time.Now()
	reply := &pb.HelloReply{}
	reply.Date = timestamppb.New(date)
	return reply, nil
}
