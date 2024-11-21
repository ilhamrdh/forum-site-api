package posts

import (
	"context"

	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostById(ctx context.Context, postID int64) (*posts.DetailPostResponse, error) {
	post, err := s.postRepo.GetPostById(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id to database")
		return nil, err
	}

	countLike, err := s.postRepo.CountLikeByPost(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like to database")
		return nil, err
	}

	comment, err := s.postRepo.GetCommentByPost(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment post to database")
		return nil, err
	}
	return &posts.DetailPostResponse{
		Post: posts.PostResponse{
			ID:           post.ID,
			UserID:       post.UserID,
			Username:     post.Username,
			PostTitle:    post.PostTitle,
			PostContent:  post.PostContent,
			PostHashtags: post.PostHashtags,
			IsLiked:      post.IsLiked,
		},
		LikeCount: countLike,
		Comments:  comment,
	}, nil
}
