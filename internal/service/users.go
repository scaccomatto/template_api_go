package service

import (
	"context"
	"errors"
	"net/http"
	"template.com/restapi/internal/apperr"
	"template.com/restapi/internal/model"
	"template.com/restapi/internal/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(r repository.Users) *UserService {
	return &UserService{
		repo: r,
	}
}

func (u *UserService) AddUser(ctx context.Context, newUser model.User) (model.User, error) {
	// validation logic
	if newUser.Name == "" || newUser.Lastname == "" {
		return model.User{}, apperr.New(errors.New("user name or lastname is empty"), http.StatusBadRequest, "invalid data")
	}
	// implement here business logic needed...

	return u.repo.AddUser(ctx, newUser)
}

func (u *UserService) GetUserById(ctx context.Context, id int) (model.User, error) {
	if id == 0 {
		return model.User{}, apperr.New(errors.New("user not found. review input data"), http.StatusNotFound, "user not found")
	}
	usr, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return usr, err
}
