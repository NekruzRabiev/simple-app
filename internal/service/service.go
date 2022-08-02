package service

import (
	"context"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/repository"
	"github.com/nekruzrabiev/simple-app/pkg/jwt"
	"github.com/nekruzrabiev/simple-app/pkg/rnd"
	"regexp"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
//ALL services
type Services struct {
	User           User
	RefreshSession RefreshSession
}

func NewServices(deps Deps) *Services {
	refreshSession := newRefreshSessionService(deps.Repos.RefreshSession, deps.JwtManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	user := newUserService(deps.Repos.User, deps.RndGen, refreshSession, deps.ReValidPassword, deps.Repos.Transactor)
	return &Services{
		RefreshSession: refreshSession,
		User:           user,
	}
}

//Dependencies we need to inject to the service
type Deps struct {
	Repos           *repository.Repositories
	ReValidPassword *regexp.Regexp
	JwtManager      jwt.TokenManager
	RndGen          rnd.Generator
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type UserSignInInfo struct {
	AccessToken  string
	RefreshToken string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type UserUpdateInput struct {
	Id   int
	Name string
}

type UserSignInInput struct {
	Email    string
	Password string
}

// User service
type User interface {
	Create(ctx context.Context, staff domain.User) (int, error)
	SignIn(ctx context.Context, input UserSignInInput) (UserSignInInfo, error)
	UpdateName(ctx context.Context, input UserUpdateInput) error
	Get(ctx context.Context, id int) (*domain.User, error)
}

//Refresh tokens sessions
type RefreshSession interface {
	Create(ctx context.Context, userId int) (Tokens, error)
	Update(ctx context.Context, refreshToken string) (Tokens, error)
}
