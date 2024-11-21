package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")

	now := time.Now()
	model := posts.Post{
		UserID:       userId,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHashtags,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.FormatInt(userId, 10),
		UpdatedBy:    strconv.FormatInt(userId, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error create post to repository")
		return err
	}

	return nil
}
