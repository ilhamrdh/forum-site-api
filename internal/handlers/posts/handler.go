package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/middleware"
	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/ilhamrdh/situs-forum/internal/models/web"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (web.ListResponse[posts.PostResponse], error)
	CreateCommnet(ctx context.Context, postId, userId int64, request posts.CreateCommentRequest) error
	UpdateUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error
	GetPostById(ctx context.Context, postID int64) (*posts.DetailPostResponse, error)
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) PostRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("", h.CreatePost)
	route.GET("", h.GetAllPost)
	route.GET("/:post_id", h.GetPostById)

	route.POST("/comment/:post_id", h.CreateCommnet)

	route.PUT("/user-activity/like/:post_id", h.UpsetUserActivity)

}
