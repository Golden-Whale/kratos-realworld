package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "kratos-realworld/api/helloworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{Username: "foo"}}, nil
}
func (s *RealWorldService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, in *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (s *RealWorldService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *RealWorldService) GetProfile(ctx context.Context, in *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (s *RealWorldService) FollowUser(ctx context.Context, in *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, in *v1.UnFollowUserRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (s *RealWorldService) ListArticles(ctx context.Context, in *v1.ListArticlesRequest) (*v1.MultipleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (s *RealWorldService) FeedArticles(ctx context.Context, in *v1.FeedArticlesRequest) (*v1.MultipleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (s *RealWorldService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (s *RealWorldService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (s *RealWorldService) AddComment(ctx context.Context, in *v1.AddCommentRequest) (*v1.SingleComment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (s *RealWorldService) GetComments(ctx context.Context, in *v1.GetCommentsRequest) (*v1.MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (s *RealWorldService) DeleteComments(ctx context.Context, in *v1.DeleteCommentsRequest) (*v1.SingleComment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComments not implemented")
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, in *v1.FollowArticleRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, in *v1.UnFollowArticleRequest) (*v1.ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}
func (s *RealWorldService) GetTag(ctx context.Context, in *v1.GetTagRequest) (*v1.TagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
