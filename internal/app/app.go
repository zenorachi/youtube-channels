package app

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/zenorachi/youtube-task/api/rest"
	v1handlers "github.com/zenorachi/youtube-task/api/rest/v1/handlers"
	"github.com/zenorachi/youtube-task/internal/config"
	"github.com/zenorachi/youtube-task/internal/repository"
	"github.com/zenorachi/youtube-task/internal/server"
	"github.com/zenorachi/youtube-task/internal/services"
	"github.com/zenorachi/youtube-task/pkg/database/postgres"
	"os"
	"os/signal"
	"syscall"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	channelsService, err := services.NewChannelsService(context.Background(), cfg.APIKey)
	if err != nil {
		return err
	}

	if cfg.DBIntegration {
		db, err := postgres.NewDB(&cfg.DB)
		if err != nil {
			return err
		}

		channelsRepo := repository.NewChannelRepository(db)
		channelsService.SetRepo(channelsRepo)
	}

	srv := server.New(
		cfg,
		rest.InitRouter(v1handlers.NewChannelsHandler(channelsService)),
	)
	srv.Run()
	log.Infof("started listening on port: %d", cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-quit:
		log.Infof("server stopped with signal %s", s.String())
	case err = <-srv.Notify():
		return err
	}

	if srv.Shutdown(ctx) != nil {
		return err
	}

	return nil
}
