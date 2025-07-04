package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	tcmysql "github.com/testcontainers/testcontainers-go/modules/mysql"
	"os"
	"path/filepath"
	"template.com/restapi/internal/model"
	"testing"
)

var repoUser *UserDb

func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlContainer, err := tcmysql.Run(ctx,
		"mysql:8.0.36",
		tcmysql.WithDatabase("foo"),
		tcmysql.WithUsername("root"),
		tcmysql.WithPassword("password"),
		tcmysql.WithScripts(filepath.Join("schema.sql")),
	)

	if err != nil {
		fmt.Printf("failed to start container: %s", err)
		panic(err)
	}

	connString, err := mysqlContainer.ConnectionString(ctx)
	if err != nil {
		panic(err)
	}
	db, err := NewConnection(connString, nil)
	repoUser = NewUserDb(db.Db)

	os.Exit(m.Run())
}

func TestGivenValidUserWhenCallingAddUserThenUserIsCreated(t *testing.T) {
	resp, err := repoUser.AddUser(context.Background(), model.User{Name: "pippo"})
	assert.NoError(t, err)
	assert.Equal(t, resp.Name, "pippo")
	assert.Greater(t, resp.Id, 0)
}

func TestGivenValidIdWhenCallingGetUserByIdThenUserIsFetched(t *testing.T) {
	resp, err := repoUser.AddUser(context.Background(), model.User{Name: "pinco"})
	assert.NoError(t, err)
	found, err := repoUser.GetUserById(context.Background(), resp.Id)
	assert.NoError(t, err)
	assert.Equal(t, resp.Name, found.Name)
	assert.Equal(t, resp.Id, found.Id)
}

func TestGivenInvalidIdWhenCallingGetUserByIdThenUserNotFound(t *testing.T) {
	_, err := repoUser.AddUser(context.Background(), model.User{Name: "pinco"})
	assert.NoError(t, err)
	_, err = repoUser.GetUserById(context.Background(), -1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}
