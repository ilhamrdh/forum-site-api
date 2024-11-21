package posts

import "time"

type Post struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	PostTitle    string    `db:"post_title"`
	PostContent  string    `db:"post_content"`
	PostHashtags string    `db:"post_hashtags"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedBy    string    `db:"updated_by"`
}

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"post_title"`
		PostContent  string   `json:"post_content"`
		PostHashtags []string `json:"post_hashtags"`
	}

	PostResponse struct {
		ID           int64    `json:"id"`
		UserID       int64    `json:"user_id"`
		Username     string   `json:"username"`
		PostTitle    string   `json:"post_title"`
		PostContent  string   `json:"post_content"`
		PostHashtags []string `json:"post_hashtags"`
		IsLiked      bool     `json:"is_liked"`
	}

	DetailPostResponse struct {
		Post      PostResponse      `json:"post"`
		LikeCount int               `json:"like_count"`
		Comments  []CommentResponse `json:"comments"`
	}
)
