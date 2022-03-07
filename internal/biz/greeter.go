package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RealWorld struct {
	Hello string
}

type RealWorldRepo interface {
	CreateRealWorld(context.Context, *RealWorld) error
	UpdateRealWorld(context.Context, *RealWorld) error
}

type RealWorldUsecase struct {
	repo RealWorldRepo
	log  *log.Helper
}

func NewRealWorldUsecase(repo RealWorldRepo, logger log.Logger) *RealWorldUsecase {
	return &RealWorldUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RealWorldUsecase) Create(ctx context.Context, g *RealWorld) error {
	return uc.repo.CreateRealWorld(ctx, g)
}

func (uc *RealWorldUsecase) Update(ctx context.Context, g *RealWorld) error {
	return uc.repo.UpdateRealWorld(ctx, g)
}
