package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"template.com/restapi/internal/app/httpserver"
	"template.com/restapi/internal/app/httpserver/handlers"
	"template.com/restapi/internal/pkg/apperrors"
	"template.com/restapi/internal/pkg/services/data"

	"template.com/restapi/internal/pkg/configuration"
)

func Start() error {
	c := context.TODO()
	ctx, stopService := context.WithCancel(c)

	defer stopService()

	appId := uuid.New()
	log.Info().Msgf("starting up app instance with id:%v", appId)
	defer log.Info().Msgf("app instance %v stopped successfully", appId)

	if err := configuration.LoadAppConfig(); err != nil {
		log.Error().Err(err).Msg("no environment config found, using default config")
		return apperrors.Builder().Message("no config found").Build()
	}

	dataServices := data.NewService()
	httpServer := httpserver.New(dataServices)

	handlers.Init(httpServer)

	go httpServer.Start(ctx, stopService)
	defer httpServer.Shutdown(ctx)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	var ctxError error

terminate:
	for {
		select {
		case <-termChan:
			break terminate
		case <-ctx.Done():
			ctxError = ctx.Err()
			break terminate
		}
	}

	return ctxError
}
