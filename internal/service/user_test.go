package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"template.com/restapi/internal/apperr"
	"template.com/restapi/internal/model"
	"template.com/restapi/internal/repository"
	"testing"
)

// for this test, instead of using TestContainer I'll use a mock.

func setupTest(t *testing.T) (*UserService, *repository.MockUsers) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := repository.NewMockUsers(ctrl)
	return NewUserService(m), m
}

func TestGivenInvalidUserWhenCallingAddUserThenError400(t *testing.T) {
	ust, mockRepoUser := setupTest(t)
	mockRepoUser.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("an error"))

	res, err := ust.AddUser(context.Background(), model.User{})
	assert.ErrorAs(t, err, &apperr.StatusError{})
	assert.Contains(t, err.Error(), "invalid data")
	assert.Equal(t, model.User{}, res)
}

func TestGivenUserWhenCallingAddUserThenSuccess(t *testing.T) {
	ust, mockRepoUser := setupTest(t)
	mockRepoUser.EXPECT().AddUser(gomock.Any(), model.User{Name: "pinco", Lastname: "pallino"}).Return(model.User{Name: "pinco", Lastname: "pallino", Id: 2}, nil)

	res, err := ust.AddUser(context.Background(), model.User{Name: "pinco", Lastname: "pallino"})
	assert.NoError(t, err)
	assert.Equal(t, "pinco", res.Name)
	assert.Equal(t, 2, res.Id)

}
