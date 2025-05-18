package v1

import (
	"github.com/alexKudryavtsev-web/beyond-limits-app/config"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	handler *gin.Engine,
	logger logger.Interface,
	cfg config.Admin,
	authUseCase usecase.Auth,
	referencesUseCase usecase.References,
	picturesUseCase usecase.Pictures,
	newsUseCase usecase.News,
) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	apiRouter := handler.Group("/api")
	{
		newCommonRoutes(apiRouter)
		newAuthRoutes(apiRouter, logger, authUseCase)

		authMiddleware := middleware.AuthMiddleware(logger, cfg.JWTSecret)

		newReferencesRoutes(apiRouter, logger, referencesUseCase, authMiddleware)
		newPicturesRoutes(apiRouter, logger, picturesUseCase, authMiddleware)
		newNewsRoutes(apiRouter, logger, newsUseCase, authMiddleware)
	}

	NewFrontendRouter(
		handler,
		logger,
		picturesUseCase,
		referencesUseCase,
	)
}
