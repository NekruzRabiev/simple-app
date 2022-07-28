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

//ALL services
type Services struct {
	User            User
	RefreshSessions RefreshSessions
}

//Dependencies we need to inject to the service
type Deps struct {
	Repos           *repository.Repositories
	ReValidEmail    *regexp.Regexp
	ReValidPassword *regexp.Regexp
	JwtManager      jwt.TokenManager
	RndGen          rnd.Generator
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type UserSignInInfo struct {
	AccessToken  string
	RefreshToken string
	IsUserExist  bool
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type UserUpdateInput struct {
	Id   int
	Name string
}

type UserGetInput struct {
	Id   int
	Name string
}

// User service
type User interface {
	Create(ctx context.Context, staff domain.User) (int, error)
	SignIn(ctx context.Context, email string, password string) (UserSignInInfo, error)
}

//Refresh tokens sessions
type RefreshSessions interface {
	Create(ctx context.Context, userId int) (Tokens, error)
	Update(ctx context.Context, refreshToken string) (Tokens, error)
}
