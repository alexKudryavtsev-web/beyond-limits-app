package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexKudryavtsev-web/beyond-limits-app/config"
	v1 "github.com/alexKudryavtsev-web/beyond-limits-app/internal/controller/http/v1"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase/repo"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/httpserver"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	logger, err := logger.New(cfg.Log.Level, cfg.Log.Destination)
	if err != nil {
		log.Fatalf("can't init logger: %s", err)
	}
	logger.Info("logger init")

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatalf("can't init postgres: %s", err)
	}
	defer pg.Close()

	adminUseCase := usecase.NewAuthUseCase(cfg.Admin)

	referencesRepo := repo.NewReferencesRepo(pg)
	referencesUseCase := usecase.NewReferencesUseCase(referencesRepo)

	picturesRepo := repo.NewPicturesRepo(pg)
	picturesUseCase := usecase.NewPicturesUseCase(picturesRepo)

	newsRepo := repo.NewNewsRepo(pg)
	newsUseCase := usecase.NewNewsUseCase(newsRepo)

	handler := gin.New()
	v1.NewRouter(handler, logger, cfg.Admin, adminUseCase, referencesUseCase, picturesUseCase, newsUseCase)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
