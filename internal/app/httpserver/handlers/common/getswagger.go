package common

import (
	"github.com/labstack/echo/v4"
	"template.com/restapi/internal/app/httpserver"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func GetRootSwaggerRoute(s *httpserver.Server) *echo.Route {
	return s.Router.Root.GET("*", getSwaggerHandler(s))
}

func GetSwaggerRoute(s *httpserver.Server) *echo.Route {
	return s.Router.Swagger.GET("/*", getSwaggerHandler(s))
}

func getSwaggerHandler(s *httpserver.Server) echo.HandlerFunc {
	return echoSwagger.WrapHandler
}
