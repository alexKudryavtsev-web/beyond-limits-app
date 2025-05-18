package v1

import (
	"github.com/alexKudryavtsev-web/beyond-limits-app/config"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, logger logger.Interface, cfg config.Admin, todoUseCase usecase.Todos, authUseCase usecase.Auth, referencesUseCase usecase.References) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	router := handler.Group("/api")

	newCommonRoutes(router)
	newTodosRoutes(router, logger, todoUseCase)
	newAuthRoutes(router, logger, authUseCase)
	newReferencesRoutes(router, logger, referencesUseCase, cfg)
}
