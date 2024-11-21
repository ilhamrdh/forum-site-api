package posts

import (
	"context"
	"strings"

	"github.com/ilhamrdh/situs-forum/internal/models/posts"
	"github.com/ilhamrdh/situs-forum/internal/models/web"
)

func (r *repository) CreatePost(ctx context.Context, model posts.Post) error {
	query := `INSERT INTO posts(user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by) VALUES (?,?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query,
		model.UserID,
		model.PostTitle,
		model.PostContent,
		model.PostHashtags,
		model.CreatedAt,
		model.UpdatedAt,
		model.CreatedBy,
		model.UpdatedBy,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (web.ListResponse[posts.PostResponse], error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags 
				FROM posts p JOIN users u ON p.user_id = u.id ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	response := web.ListResponse[posts.PostResponse]{}
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return response, err
	}
	defer rows.Close()

	data := make([]posts.PostResponse, 0)

	for rows.Next() {
		var (
			model    posts.Post
			username string
		)
		err = rows.Scan(
			&model.ID,
			&model.UserID,
			&username,
			&model.PostTitle,
			&model.PostContent,
			&model.PostHashtags,
		)

		if err != nil {
			return response, err
		}

		data = append(data, posts.PostResponse{
			ID:           model.ID,
			UserID:       model.UserID,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashtags: strings.Split(model.PostHashtags, ","),
		})
	}
	response.Data = data
	response.Pagination = web.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	return response, nil
}

func (r *repository) GetPostById(ctx context.Context, id int64) (*posts.PostResponse, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags, ua.is_liked
				FROM posts p 
				JOIN users u ON p.user_id = u.id 
				JOIN user_activities ua ON ua.post_id = p.id
				WHERE p.id = ?`

	var (
		model    posts.Post
		username string
		isLike   bool
	)
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&model.ID,
		&model.UserID,
		&username,
		&model.PostTitle,
		&model.PostContent,
		&model.PostHashtags,
		&isLike,
	)
	if err != nil {
		return nil, err
	}
	return &posts.PostResponse{
		ID:           model.ID,
		UserID:       model.UserID,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashtags: strings.Split(model.PostHashtags, ","),
		IsLiked:      isLike,
	}, nil
}
