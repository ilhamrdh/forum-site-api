package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/middleware"
	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := route.Group("/refresh")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/get", h.RefreshToken)
}
