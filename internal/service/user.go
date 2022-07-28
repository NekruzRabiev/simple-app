package service

import (
	"context"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/repository"
	"github.com/nekruzrabiev/simple-app/pkg/rnd"
)

type userService struct {
	repo              repository.User
	rndGen            rnd.Generator
	refSessionService RefreshSession
}

func newUserService(repo repository.User, rndGen rnd.Generator, refSessionService RefreshSession) *userService {
	return &userService{
		repo:              repo,
		rndGen:            rndGen,
		refSessionService: refSessionService,
	}
}

func (s *userService) Create(ctx context.Context, user domain.User) (int, error) {

}

func (s *userService) SignIn(ctx context.Context, input UserSignInInput) (UserSignInInfo, error) {

}

func (s *userService) UpdateName(ctx context.Context, id int, name string) error {
	return nil
}

func (s *userService) Get(ctx context.Context, id int) (domain.User, error) {

}
