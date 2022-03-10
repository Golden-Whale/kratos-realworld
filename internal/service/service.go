package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos-realworld/api/helloworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnsafeRealWorldServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewRealWorldService new a realworld service.
func NewRealWorldService(uc *biz.UserUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}
