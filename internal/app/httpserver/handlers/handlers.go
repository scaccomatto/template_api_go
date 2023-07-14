package handlers

import (
	"github.com/labstack/echo/v4"
	"template.com/restapi/internal/app/httpserver"
	"template.com/restapi/internal/app/httpserver/handlers/datahandler"
)

func Init(s *httpserver.Server) {
	s.Router.Routes = []*echo.Route{
		//common.GetRootSwaggerRoute(s),
		//common.GetSwaggerRoute(s),

		datahandler.GetDataById(s),
		datahandler.CreateData(s),
	}
}
