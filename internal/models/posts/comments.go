package posts

import "time"

type Comment struct {
	ID             int64     `db:"id"`
	PostID         int64     `db:"post_id"`
	UserID         int64     `db:"user_id"`
	CommentContent string    `db:"comment_content"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	CreatedBy      string    `db:"created_by"`
	UpdatedBy      string    `db:"updated_by"`
}

type (
	CreateCommentRequest struct {
		CommentContent string `json:"comment_content"`
	}

	CommentResponse struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"user_id"`
		Username       string `json:"username"`
		CommentContent string `json:"comment_content"`
	}
)