package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"template.com/restapi/internal/apperr"
)

// @title Template Api
// @version 1.0
// @description This is a template api for a service with DB

// @host localhost:8081
// @BasePath /api/v1

func customHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		return MappingError(err, c)
	}
}

// MappingError will handle all the errors: if the error is coming as StatusError, it will be used that status
// If you want to hide the error message to the external client, you can set the ExternalClientMsg field
func MappingError(err error, c echo.Context) error {
	if err == nil {
		return err
	}
	httpError := &apperr.StatusError{}
	if errors.As(err, httpError) {
		if httpError.ExternalClientMsg != "" { // we only show external client message
			_ = c.JSON(httpError.Status, map[string]string{"error": httpError.ExternalClientMsg})
			return echo.NewHTTPError(httpError.Status, err)
		}
		_ = c.JSON(httpError.Status, map[string]string{"error": err.Error()})
		return echo.NewHTTPError(httpError.Status, err)
	}
	_ = c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}
