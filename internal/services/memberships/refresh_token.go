package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
	"github.com/ilhamrdh/situs-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("gagal untuk mengambil data user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user tidak ada")
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", nil
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}

	token, err := jwt.GenerateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
