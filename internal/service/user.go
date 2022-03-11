package service

import (
	"context"
	v1 "kratos-realworld/api/helloworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{Username: "foo"}}, nil
}

func (s *RealWorldService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.UserReply, error) {
	loginUser, err := s.uc.Register(ctx, in.User.Username, in.User.Email, in.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: &v1.UserReply_User{Username: loginUser.Username, Email: loginUser.Email, Bio: loginUser.Bio, Image: loginUser.Bio, Token: loginUser.Token}}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, in *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{Username: "foo"}}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{Username: "foo"}}, nil
}
