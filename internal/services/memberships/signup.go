package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		if user.Username != "" {
			return errors.New("username already exists")
		}
		if user.Email != "" {
			return errors.New("email already exists")
		}
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Username,
		UpdatedBy: req.Username,
	}
	return s.membershipRepo.Save(ctx, model)
}
