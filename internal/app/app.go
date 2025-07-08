package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	_ "template.com/restapi/api/swagger"
	"template.com/restapi/internal/appmiddleware"
	"template.com/restapi/internal/conf"
	"template.com/restapi/internal/handler"
	"template.com/restapi/internal/logger"
	"template.com/restapi/internal/repository"
	"template.com/restapi/internal/service"

	"time"
)

func Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// configuration
	appConfig, err := conf.LoadConfig("")
	if err != nil {
		logger.L.Error("no configuration found")
	}

	// server
	e := echo.New()
	base := fmt.Sprintf("%s/%s", appConfig.BasePath, appConfig.Version)
	serverGroup := e.Group(base)
	// adding swagger definition
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	appmiddleware.Add(e)

	// DI: db -> repo -> services -> handler
	//
	//database pool to manage all the connections to the DB
	dbPool, err := repository.NewConnection("", appConfig)
	if err != nil {
		logger.L.Error("connection issue", "error:", err)
		os.Exit(1)
	}
	// repository for direct injection
	repoUser := repository.NewUserDb(dbPool.Db)

	// Service for direct injection
	us := service.NewUserService(repoUser)

	// Handler for pair requests and services
	_ = handler.NewUserHandle(us, serverGroup)

	// Start and graceful shutdown server
	go func() {
		err := e.Start(fmt.Sprintf(":%d", appConfig.Port))
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.L.Error("shutting down the server", "error:", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.L.Error("job done. shutting down the server", "error:", err)
		os.Exit(1)
	} else {
		logger.L.Info("job done. shutting down the server")
	}
}
