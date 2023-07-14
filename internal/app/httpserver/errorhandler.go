package httpserver

import (
	"net/http"
	"template.com/restapi/internal/pkg/apperrors"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var statusCode int
		var errResp apperrors.HttpError

		switch err := err.(type) {
		case apperrors.AppError:
			switch err.Status {
			case apperrors.NotFound:
				statusCode = http.StatusNotFound
			case apperrors.BadRequest:
				statusCode = http.StatusBadRequest
			}
			errResp = apperrors.HttpError{
				StatusCode: statusCode,
				Message:    err.Message,
				Details:    err.Details,
				MetaData:   err.InnerError,
			}
		default:
			statusCode = http.StatusInternalServerError
			errResp = apperrors.HttpError{
				Message: "A general, non-specific exception occurred.",
				Details: err.Error(),
			}
		}

		if !c.Response().Committed {
			var responseSetErr error

			if c.Request().Method == http.MethodHead {
				responseSetErr = c.NoContent(statusCode)
			} else {
				responseSetErr = c.JSON(statusCode, errResp)
			}

			if responseSetErr != nil {
				log.Error().Msg("error setting error response handling error")
			}
		}
	}
}
