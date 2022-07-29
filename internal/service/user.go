package service

import (
	"context"
	"github.com/nekruzrabiev/simple-app/internal/domain"
	"github.com/nekruzrabiev/simple-app/internal/repository"
	"github.com/nekruzrabiev/simple-app/pkg/rnd"
	"github.com/nekruzrabiev/simple-app/pkg/utils"
	"regexp"
)

type userService struct {
	repo              repository.User
	rndGen            rnd.Generator
	refSessionService RefreshSession
	reValidPassword   *regexp.Regexp
	transactor        repository.Transactor
}

func newUserService(repo repository.User, rndGen rnd.Generator, refSessionService RefreshSession, reValidPassword *regexp.Regexp, transactor repository.Transactor) *userService {
	return &userService{
		repo:              repo,
		rndGen:            rndGen,
		refSessionService: refSessionService,
		reValidPassword:   reValidPassword,
		transactor:        transactor,
	}
}

func (s *userService) Create(ctx context.Context, user domain.User) (int, error) {
	if !s.reValidPassword.MatchString(user.Password) {
		return 0, ErrInvalidPassword
	}

	var id int

	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		exists, err := s.repo.Contains(ctx, user.Email)
		if err != nil {
			return err
		}
		if exists {
			return ErrUserExists
		}

		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			return err
		}

		id, err = s.repo.Create(ctx, user)
		return err
	})

	return id, err
}

func (s *userService) SignIn(ctx context.Context, input UserSignInInput) (UserSignInInfo, error) {
	var tokens Tokens

	err := s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		exists, err := s.repo.Contains(ctx, input.Email)
		if err != nil {
			return err
		}
		if !exists {
			return ErrEmailOrPassword
		}

		user, err := s.repo.GetByEmail(ctx, input.Email)
		if err != nil {
			return err
		}

		err = utils.CheckPassword(input.Password, user.Password)
		if err != nil {
			return ErrEmailOrPassword
		}

		tokens, err = s.refSessionService.Create(ctx, user.Id)
		return err
	})
	if err != nil {
		return UserSignInInfo{}, err
	}

	return UserSignInInfo{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (s *userService) UpdateName(ctx context.Context, input UserUpdateInput) error {
	return s.repo.UpdateName(ctx, input.Id, input.Name)
}

func (s *userService) Get(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.Get(ctx, id)
}
