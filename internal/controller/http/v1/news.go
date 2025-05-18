package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/gin-gonic/gin"
)

type newsRoutes struct {
	u usecase.News
	l logger.Interface
}

func newNewsRoutes(handler *gin.RouterGroup, l logger.Interface, n usecase.News, authMiddleware gin.HandlerFunc) {
	r := newsRoutes{n, l}

	handler.GET("/news", r.doGetNews)
	handler.GET("/news/:id", r.doGetNewsByID)

	adminHandler := handler.Group("/admin", authMiddleware)
	{
		adminHandler.POST("/news", r.doCreateNews)
		adminHandler.PATCH("/news/:id", r.doUpdateNews)
		adminHandler.DELETE("/news/:id", r.doDeleteNews)
	}
}

// @Summary     Get news
// @Description Get all news
// @ID          get-news
// @Tags        news
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.News
// @Failure     500 {object} response
// @Router      /news [get]
func (n *newsRoutes) doGetNews(ctx *gin.Context) {
	news, err := n.u.GetNews(ctx.Request.Context())
	if err != nil {
		n.l.Error(err, "http - v1 - doGetNews")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, news)
}

// @Summary     Get news by ID
// @Description Get news by ID
// @ID          get-news-by-id
// @Tags        news
// @Accept      json
// @Produce     json
// @Param       id path int true "News ID"
// @Success     200 {object} entity.News
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /news/{id} [get]
func (n *newsRoutes) doGetNewsByID(ctx *gin.Context) {
	id := ctx.Param("id")
	newsID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	news, err := n.u.GetNewsByID(ctx.Request.Context(), newsID)
	if err != nil {
		if errors.Is(err, entity.ErrNewsNotFound) {
			errorResponse(ctx, http.StatusNotFound, "news not found")
			return
		}
		n.l.Error(err, "http - v1 - doGetNewsByID")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, news)
}

// @Summary     Create news
// @Description Create news
// @ID          create-news
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body entity.NewsCreateRequest true "News data"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/news [post]
// @Security    BearerAuth
func (n *newsRoutes) doCreateNews(ctx *gin.Context) {
	var req entity.NewsCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		n.l.Error(err, "http - v1 - doCreateNews")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := n.u.CreateNews(ctx.Request.Context(), req); err != nil {
		n.l.Error(err, "http - v1 - doCreateNews")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Update news
// @Description Update news
// @ID          update-news
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "News ID"
// @Param       request body entity.NewsUpdateRequest true "News data"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/news/{id} [patch]
// @Security    BearerAuth
func (n *newsRoutes) doUpdateNews(ctx *gin.Context) {
	id := ctx.Param("id")
	newsID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	var req entity.NewsUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		n.l.Error(err, "http - v1 - doUpdateNews")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := n.u.UpdateNews(ctx.Request.Context(), newsID, req); err != nil {
		n.l.Error(err, "http - v1 - doUpdateNews")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Delete news
// @Description Delete news
// @ID          delete-news
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "News ID"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/news/{id} [delete]
// @Security    BearerAuth
func (n *newsRoutes) doDeleteNews(ctx *gin.Context) {
	id := ctx.Param("id")
	newsID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := n.u.DeleteNews(ctx.Request.Context(), newsID); err != nil {
		n.l.Error(err, "http - v1 - doDeleteNews")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
