// internal/controller/http/v1/auth.go
package v1

import (
	"net/http"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/gin-gonic/gin"
)

type authRoutes struct {
	u usecase.Auth
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, l logger.Interface, a usecase.Auth) {
	r := authRoutes{a, l}

	handler.POST("/admin/login", r.doLogin)
}

type doLoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary     Admin login
// @Description Login admin
// @ID          admin-login
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body doLoginRequest true "Login and password"
// @Success     200 {object} entity.AuthResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/login [post]
func (a *authRoutes) doLogin(ctx *gin.Context) {
	var request doLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		a.l.Error(err, "http - v1 - doLogin")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := a.u.Login(ctx.Request.Context(), request.Login, request.Password)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	ctx.JSON(http.StatusOK, entity.AuthResponse{Token: token})
}
