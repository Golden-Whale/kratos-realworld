package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) CreateUser(ctx *context.Context, user *biz.User) error {
	return nil
}
