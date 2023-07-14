package httpserver

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"template.com/restapi/internal/pkg/apperrors"
	"template.com/restapi/internal/pkg/configuration"
	"template.com/restapi/internal/pkg/services/data"

	"github.com/labstack/echo-contrib/pprof"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *Router
	echo   *echo.Echo

	DataService DataService
}

type (
	Router struct {
		Routes  []*echo.Route
		Root    *echo.Group
		Swagger *echo.Group
	}

	DataService interface {
		CreateData(name string, value int) (data.Data, error)
		GetDataById(id uuid.UUID) (data.Data, error)
	}
)

func New(ds DataService) *Server {
	s := &Server{
		echo:        echo.New(),
		DataService: ds,
	}

	s.echo.HideBanner = true
	s.echo.HTTPErrorHandler = HTTPErrorHandler()

	//docs.SwaggerInfo.BasePath = configuration.Application.Http.PrefixPath
	s.echo.Use(echoMiddleware.Recover())

	pprof.Register(s.echo)

	s.Router = &Router{
		Routes: nil, // will be populated by handlers.AttachAllRoutes(s)

		Root:    s.echo.Group(""),
		Swagger: s.echo.Group(configuration.Application.Http.BasePath + "swagger"),
	}

	//docs.SwaggerInfo.BasePath = configuration.Application.Http.BasePath

	return s
}

func (s *Server) Start(ctx context.Context, shutdownCallback func()) {
	var err error
	defer func() {
		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("http server encountered an error, initiating shutdown")
			shutdownCallback()
		}
	}()

	if !s.Ready() {
		err = apperrors.Builder().Message("server not ready").Build()
	}

	err = s.echo.Start(configuration.Application.Http.HostPort)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Ready() bool {
	return s.echo != nil && s.Router != nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.echo.ServeHTTP(w, r)
}
