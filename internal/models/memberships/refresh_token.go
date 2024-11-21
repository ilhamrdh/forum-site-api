package memberships

import "time"

type RefreshToken struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	ExpiredAt    time.Time `db:"expired_at"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedBy    string    `db:"updated_by"`
}

type (
	RefreshTokenRequest struct {
		Token string `json:"token"`
	}

	RefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
	}
)
