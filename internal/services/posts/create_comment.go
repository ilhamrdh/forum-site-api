package posts

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateCommnet(ctx context.Context, postId, userId int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.Comment{
		PostID:         postId,
		UserID:         userId,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userId, 10),
		UpdatedBy:      strconv.FormatInt(userId, 10),
	}
	fmt.Println("Log Model ", postId)
	err := s.postRepo.CreateCommnet(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment to repository")
	}
	return nil
}
