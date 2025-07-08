package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"template.com/restapi/internal/apperr"
	"template.com/restapi/internal/model"
)

const UserRepoName = "UserRepository"

type Users interface {
	AddUser(ctx context.Context, newUser model.User) (model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
}

type UserDb struct {
	db *gorm.DB
}

func NewUserDb(db *gorm.DB) *UserDb {
	return &UserDb{db: db}
}

func (u *UserDb) AddUser(ctx context.Context, newUser model.User) (model.User, error) {
	res := u.db.Create(&newUser)
	if res.Error != nil {
		return model.User{}, res.Error
	}
	return newUser, nil
}

func (u *UserDb) GetUserById(ctx context.Context, id int) (model.User, error) {
	target := model.User{Id: id}
	res := u.db.First(&target, nil)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.User{}, apperr.New(fmt.Errorf("%w", res.Error), http.StatusNotFound)
	}
	return target, res.Error // Because we don't use a "StatusError", this will always be mapped to 500
}

func (u *UserDb) GetUser(ctx context.Context, target model.User) (model.User, error) {
	res := u.db.First(&target, nil)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return model.User{}, apperr.New(fmt.Errorf("%w", res.Error), http.StatusNotFound)
	}
	return target, res.Error
}
