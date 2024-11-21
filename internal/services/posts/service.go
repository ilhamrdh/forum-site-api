package posts

import (
	"context"

	"github.com/ilhamrdh/situs-forum/internal/configs"
	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/ilhamrdh/situs-forum/internal/models/web"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.Post) error
	GetAllPost(ctx context.Context, limit, offset int) (web.ListResponse[posts.PostResponse], error)
	GetPostById(ctx context.Context, id int64) (*posts.PostResponse, error)

	CreateCommnet(ctx context.Context, model posts.Comment) error
	GetUserActivity(ctx context.Context, model posts.UserActivity) (*posts.UserActivity, error)
	GetCommentByPost(ctx context.Context, postID int64) ([]posts.CommentResponse, error)

	CreateUserActivity(ctx context.Context, model posts.UserActivity) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivity) error
	CountLikeByPost(ctx context.Context, postID int64) (int, error)
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		postRepo: postRepo,
		cfg:      cfg,
	}
}
