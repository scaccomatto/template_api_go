package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"template.com/restapi/internal/apperr"
	"template.com/restapi/internal/model"
	"template.com/restapi/internal/service"
)

type UserHandle struct {
	userService *service.UserService
}

// NewUserHandle will create an instance of user handler and also register handler routes
func NewUserHandle(us *service.UserService, group *echo.Group) *UserHandle {

	uh := &UserHandle{userService: us}

	group.GET("/users/:id", customHandler(uh.GetUserById))
	group.POST("/users", customHandler(uh.AddUser))

	return uh
}

// AddUser @Summary it creates a user
// @Tags 		AddUser
// @Accept 		json
// @Produce      json
// @Param 		user body model.User true "User payload"
// @Success 	201  {object}  model.User
// @Failure     400  {object}  apperr.StatusError
// @Failure     500  {object}  error
// @Router 		/users [post]
func (h *UserHandle) AddUser(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	body, err := h.userService.AddUser(c.Request().Context(), newUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, body)
}

// GetUserById @Summary it fetches the user looking for the ID
// @Tags 		UserGetId
// @Accept 		json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success 	200  {object}  model.User
// @Failure     400  {object}  apperr.StatusError
// @Failure     500  {object}  error
// @Router 		/users/{id} [get]
func (h *UserHandle) GetUserById(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if idString == "" || idString == "0" || err != nil {
		return apperr.New(fmt.Errorf("invalid user id"), http.StatusBadRequest)
	}

	body, err := h.userService.GetUserById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
}
