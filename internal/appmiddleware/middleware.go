package appmiddleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"template.com/restapi/internal/logger"
)

const (
	HeaderCorrelationID = "correlation-id"
)

func CorrelationID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			cid := req.Header.Get(HeaderCorrelationID)

			if cid == "" {
				cid = uuid.New().String()
			}

			res.Header().Set(HeaderCorrelationID, cid)
			ctx := context.WithValue(req.Context(), HeaderCorrelationID, cid)
			request := c.Request().WithContext(ctx)
			c.SetRequest(request)

			return next(c)
		}
	}
}

func Add(server *echo.Echo) {
	// Middleware
	server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: false, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			corrId := fmt.Sprintf("%s", c.Request().Context().Value(HeaderCorrelationID))
			if v.Error == nil {
				logger.L.Error("REQUEST",
					slog.String("method", c.Request().Method),
					slog.String("corr-id", corrId),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.L.ErrorContext(context.Background(), "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.String("method", c.Request().Method),
					slog.String("corr-id", corrId),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	server.Use(CorrelationID())
	server.Use(middleware.Recover()) // Recovers from panics
	server.Use(middleware.CORS())    // Enables CORS

}
