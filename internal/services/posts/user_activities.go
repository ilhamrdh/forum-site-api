package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpdateUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()

	model := posts.UserActivity{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	user, err := s.postRepo.GetUserActivity(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("error get user from database")
		return err
	}

	if user == nil {
		if !request.IsLiked {
			return errors.New("anda belum pernah like sebelum nya")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity to database")
		return err
	}

	return nil
}
