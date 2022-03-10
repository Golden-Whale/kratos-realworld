package service

import (
	"context"
	v1 "kratos-realworld/api/helloworld/v1"
)

func (s *RealWorldService) GetProfile(ctx context.Context, in *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, in *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, in *v1.UnFollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, in *v1.FollowArticleRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, in *v1.UnFollowArticleRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, in *v1.FeedArticlesRequest) (*v1.MultipleArticleReply, error) {
	return &v1.MultipleArticleReply{}, nil
}
func (s *RealWorldService) CreateArticle(ctx context.Context, in *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) GetArticle(ctx context.Context, in *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) ListArticles(ctx context.Context, in *v1.ListArticlesRequest) (*v1.MultipleArticleReply, error) {
	return &v1.MultipleArticleReply{}, nil
}
func (s *RealWorldService) AddComment(ctx context.Context, in *v1.AddCommentRequest) (*v1.SingleComment, error) {
	return &v1.SingleComment{}, nil
}
func (s *RealWorldService) DeleteComments(ctx context.Context, in *v1.DeleteCommentsRequest) (*v1.SingleComment, error) {
	return &v1.SingleComment{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, in *v1.GetCommentsRequest) (*v1.MultipleCommentsReply, error) {
	return &v1.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) GetTag(ctx context.Context, in *v1.GetTagRequest) (*v1.TagListReply, error) {
	return &v1.TagListReply{}, nil
}
