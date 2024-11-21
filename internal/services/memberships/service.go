package memberships

import (
	"context"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/configs"
	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.User, error)
	Save(ctx context.Context, model memberships.User) error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshToken) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshToken, error)
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepository
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		membershipRepo: membershipRepo,
		cfg:            cfg,
	}
}
