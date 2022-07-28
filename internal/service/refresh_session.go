package service

import (
	"context"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/repository"
	"github.com/nekruzrabiev/simple-app/pkg/jwt"
	"strconv"
	"time"
)

type refreshSessionService struct {
	repo repository.RefreshSession

	jwtManager      jwt.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func newRefreshSessionService(repo repository.RefreshSession, jwtManager jwt.TokenManager,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *refreshSessionService {
	return &refreshSessionService{
		repo:            repo,
		jwtManager:      jwtManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL}
}

func (s *refreshSessionService) Create(ctx context.Context, userId int) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.jwtManager.NewJWT(strconv.Itoa(userId), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.jwtManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domain.RefreshSession{
		Token:     res.RefreshToken,
		ExpiresAt: time.Now().Add(s.refreshTokenTTL),
		UserId:    userId,
	}

	err = s.repo.Create(ctx, session)

	return res, err
}

func (s *refreshSessionService) Update(ctx context.Context, refreshToken string) (Tokens, error) {
	refreshSession, err := s.repo.GetByToken(ctx, refreshToken, time.Now())
	if err != nil {
		return Tokens{}, err
	}
	return s.Create(ctx, refreshSession.UserId)
}
