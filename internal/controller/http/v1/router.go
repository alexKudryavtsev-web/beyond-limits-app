package v1

import (
	"github.com/alexKudryavtsev-web/beyond-limits-app/config"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, logger logger.Interface, cfg config.Admin, todoUseCase usecase.Todos, authUseCase usecase.Auth, referencesUseCase usecase.References, picturesUseCase usecase.Pictures, newsUseCase usecase.News) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.Static("/uploads", "./uploads")

	router := handler.Group("/api")

	authMiddleare := middleware.AuthMiddleware(logger, cfg.JWTSecret)

	newCommonRoutes(router)
	newTodosRoutes(router, logger, todoUseCase)
	newAuthRoutes(router, logger, authUseCase)
	newReferencesRoutes(router, logger, referencesUseCase, authMiddleare)
	newPicturesRoutes(router, logger, picturesUseCase, authMiddleare)
	newNewsRoutes(router, logger, newsUseCase, authMiddleare)
}
