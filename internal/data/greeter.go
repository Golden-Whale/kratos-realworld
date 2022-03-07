package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

type realworldRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealWorldRepo .
func NewRealWorldRepo(data *Data, logger log.Logger) biz.RealWorldRepo {
	return &realworldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *realworldRepo) CreateRealWorld(ctx context.Context, g *biz.RealWorld) error {
	return nil
}

func (r *realworldRepo) UpdateRealWorld(ctx context.Context, g *biz.RealWorld) error {
	return nil
}
