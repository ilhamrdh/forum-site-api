package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
	"github.com/ilhamrdh/situs-forum/pkg/jwt"
	tokenUtil "github.com/ilhamrdh/situs-forum/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("gagal untuk mengambil data user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email tidak ada")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", errors.New("email atau password salah")
	}

	token, err := jwt.GenerateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, int64(user.ID), time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get latest refresh token from database")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	log.Print("Generated Refresh Token in Service: ", refreshToken)
	if refreshToken == "" {
		log.Error().Msg("Failed to generate refresh token, empty token")
		return token, "", errors.New("failed to generate refresh token")
	}

	now := time.Now()
	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshToken{
		UserID:       int64(user.ID),
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    strconv.FormatInt(int64(user.ID), 10),
		UpdatedBy:    strconv.FormatInt(int64(user.ID), 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("Error inserting refresh token to the database")
		return token, refreshToken, err
	}

	log.Print("Successfully inserted refresh token: ", refreshToken)

	return token, refreshToken, nil
}
