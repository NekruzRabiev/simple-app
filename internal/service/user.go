package service

import "github.com/nekruzrabiev/simple-app/internal/repository"

type userService struct {
	repo repository.Users
}