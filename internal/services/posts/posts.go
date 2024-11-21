package posts

import (
	"context"

	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/ilhamrdh/situs-forum/internal/models/web"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (web.ListResponse[posts.PostResponse], error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)
	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error get all post from database")
		return response, err
	}
	return response, nil
}
