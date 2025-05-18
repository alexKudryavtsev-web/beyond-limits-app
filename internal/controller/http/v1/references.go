package v1

import (
	"net/http"
	"strconv"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/usecase"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/logger"
	"github.com/gin-gonic/gin"
)

type referencesRoutes struct {
	u usecase.References
	l logger.Interface
}

func newReferencesRoutes(handler *gin.RouterGroup, l logger.Interface, r usecase.References, authMiddleware gin.HandlerFunc) {
	routes := referencesRoutes{r, l}

	handler.GET("/genres", routes.doGetGenres)
	handler.GET("/authors", routes.doGetAuthors)
	handler.GET("/dimensions", routes.doGetDimensions)
	handler.GET("/work-techniques", routes.doGetWorkTechniques)

	adminHandler := handler.Group("/admin", authMiddleware)
	{
		adminHandler.POST("/genres", routes.doCreateGenre)
		adminHandler.DELETE("/genres/:id", routes.doDeleteGenre)

		adminHandler.POST("/authors", routes.doCreateAuthor)
		adminHandler.DELETE("/authors/:id", routes.doDeleteAuthor)

		adminHandler.POST("/dimensions", routes.doCreateDimension)
		adminHandler.DELETE("/dimensions/:id", routes.doDeleteDimension)

		adminHandler.POST("/work-techniques", routes.doCreateWorkTechnique)
		adminHandler.DELETE("/work-techniques/:id", routes.doDeleteWorkTechnique)
	}
}

// @Summary     Get genres
// @Description Get all genres
// @ID          get-genres
// @Tags        references
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.Genre
// @Failure     500 {object} response
// @Router      /genres [get]
func (r *referencesRoutes) doGetGenres(ctx *gin.Context) {
	genres, err := r.u.GetGenres(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - doGetGenres")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, genres)
}

type doCreateGenreRequest struct {
	Name string `json:"name" binding:"required"`
}

// @Summary     Create genre
// @Description Create new genre
// @ID          create-genre
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body doCreateGenreRequest true "Genre name"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/genres [post]
// @Security    BearerAuth
func (r *referencesRoutes) doCreateGenre(ctx *gin.Context) {
	var request doCreateGenreRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateGenre")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.u.CreateGenre(ctx.Request.Context(), request.Name); err != nil {
		r.l.Error(err, "http - v1 - doCreateGenre")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Delete genre
// @Description Delete genre by ID
// @ID          delete-genre
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Genre ID"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/genres/{id} [delete]
// @Security    BearerAuth
func (r *referencesRoutes) doDeleteGenre(ctx *gin.Context) {
	id := ctx.Param("id")
	genreID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := r.u.DeleteGenre(ctx.Request.Context(), genreID); err != nil {
		r.l.Error(err, "http - v1 - doDeleteGenre")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// Authors handlers
// @Summary     Get authors
// @Description Get all authors
// @ID          get-authors
// @Tags        references
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.Author
// @Failure     500 {object} response
// @Router      /authors [get]
func (r *referencesRoutes) doGetAuthors(ctx *gin.Context) {
	authors, err := r.u.GetAuthors(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - doGetAuthors")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

type doCreateAuthorRequest struct {
	FullName string `json:"full_name" binding:"required"`
}

// @Summary     Create author
// @Description Create new author
// @ID          create-author
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body doCreateAuthorRequest true "Author full name"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/authors [post]
// @Security    BearerAuth
func (r *referencesRoutes) doCreateAuthor(ctx *gin.Context) {
	var request doCreateAuthorRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateAuthor")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.u.CreateAuthor(ctx.Request.Context(), request.FullName); err != nil {
		r.l.Error(err, "http - v1 - doCreateAuthor")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Delete author
// @Description Delete author by ID
// @ID          delete-author
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Author ID"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/authors/{id} [delete]
// @Security    BearerAuth
func (r *referencesRoutes) doDeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	authorID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := r.u.DeleteAuthor(ctx.Request.Context(), authorID); err != nil {
		r.l.Error(err, "http - v1 - doDeleteAuthor")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// Dimensions handlers (аналогично Authors)
// @Summary     Get dimensions
// @Description Get all dimensions
// @ID          get-dimensions
// @Tags        references
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.Dimension
// @Failure     500 {object} response
// @Router      /dimensions [get]
func (r *referencesRoutes) doGetDimensions(ctx *gin.Context) {
	dimensions, err := r.u.GetDimensions(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - doGetDimensions")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, dimensions)
}

type doCreateDimensionRequest struct {
	Width  int `json:"width" binding:"required"`
	Height int `json:"height" binding:"required"`
}

// @Summary     Create dimension
// @Description Create new dimension
// @ID          create-dimension
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body doCreateDimensionRequest true "Dimension width and height"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/dimensions [post]
// @Security    BearerAuth
func (r *referencesRoutes) doCreateDimension(ctx *gin.Context) {
	var request doCreateDimensionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateDimension")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.u.CreateDimension(ctx.Request.Context(), request.Width, request.Height); err != nil {
		r.l.Error(err, "http - v1 - doCreateDimension")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Delete dimension
// @Description Delete dimension by ID
// @ID          delete-dimension
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Dimension ID"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/dimensions/{id} [delete]
// @Security    BearerAuth
func (r *referencesRoutes) doDeleteDimension(ctx *gin.Context) {
	id := ctx.Param("id")
	dimensionID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := r.u.DeleteDimension(ctx.Request.Context(), dimensionID); err != nil {
		r.l.Error(err, "http - v1 - doDeleteDimension")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// WorkTechniques handlers (аналогично Authors)
// @Summary     Get work techniques
// @Description Get all work techniques
// @ID          get-work-techniques
// @Tags        references
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.WorkTechnique
// @Failure     500 {object} response
// @Router      /work-techniques [get]
func (r *referencesRoutes) doGetWorkTechniques(ctx *gin.Context) {
	techniques, err := r.u.GetWorkTechniques(ctx.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - doGetWorkTechniques")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, techniques)
}

type doCreateWorkTechniqueRequest struct {
	Name string `json:"name" binding:"required"`
}

// @Summary     Create work technique
// @Description Create new work technique
// @ID          create-work-technique
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body doCreateWorkTechniqueRequest true "Work technique name"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/work-techniques [post]
// @Security    BearerAuth
func (r *referencesRoutes) doCreateWorkTechnique(ctx *gin.Context) {
	var request doCreateWorkTechniqueRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateWorkTechnique")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.u.CreateWorkTechnique(ctx.Request.Context(), request.Name); err != nil {
		r.l.Error(err, "http - v1 - doCreateWorkTechnique")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Delete work technique
// @Description Delete work technique by ID
// @ID          delete-work-technique
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Work technique ID"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/work-techniques/{id} [delete]
// @Security    BearerAuth
func (r *referencesRoutes) doDeleteWorkTechnique(ctx *gin.Context) {
	id := ctx.Param("id")
	techniqueID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := r.u.DeleteWorkTechnique(ctx.Request.Context(), techniqueID); err != nil {
		r.l.Error(err, "http - v1 - doDeleteWorkTechnique")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
