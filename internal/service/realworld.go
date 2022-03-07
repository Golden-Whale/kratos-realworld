package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-realworld/api/helloworld/v1"
	"kratos-realworld/internal/biz"
)

// RealWorldService is a realworld service.
type RealWorldService struct {
	v1.UnsafeRealWorldServer

	uc  *biz.RealWorldUsecase
	log *log.Helper
}

// NewRealWorldService new a realworld service.
func NewRealWorldService(uc *biz.RealWorldUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.RealWorldServer
func (s *RealWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
