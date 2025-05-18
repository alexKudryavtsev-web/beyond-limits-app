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
	todoUseCase usecase.Todos,
	authUseCase usecase.Auth,
	referencesUseCase usecase.References,
	picturesUseCase usecase.Pictures,
	newsUseCase usecase.News,
) {
	// Мидлвары
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// API Routes
	apiRouter := handler.Group("/api")
	{
		newCommonRoutes(apiRouter)
		newTodosRoutes(apiRouter, logger, todoUseCase)
		newAuthRoutes(apiRouter, logger, authUseCase)

		authMiddleware := middleware.AuthMiddleware(logger, cfg.JWTSecret)

		newReferencesRoutes(apiRouter, logger, referencesUseCase, authMiddleware)
		newPicturesRoutes(apiRouter, logger, picturesUseCase, authMiddleware)
		newNewsRoutes(apiRouter, logger, newsUseCase, authMiddleware)
	}

	// Frontend Routes
	NewFrontendRouter(
		handler,
		logger,
		picturesUseCase,
		referencesUseCase,
	)
}
