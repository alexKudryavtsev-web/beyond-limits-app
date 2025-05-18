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

type picturesRoutes struct {
	u usecase.Pictures
	l logger.Interface
}

func newPicturesRoutes(handler *gin.RouterGroup, l logger.Interface, p usecase.Pictures, authMiddleware gin.HandlerFunc) {
	r := picturesRoutes{p, l}

	// Public routes
	handler.GET("/pictures", r.doGetPictures)
	handler.GET("/pictures/:id", r.doGetPictureByID)

	// Admin routes
	adminHandler := handler.Group("/admin", authMiddleware)
	{
		adminHandler.POST("/pictures", r.doCreatePicture)
		adminHandler.PATCH("/pictures/:id", r.doUpdatePicture)
		adminHandler.DELETE("/pictures/:id", r.doDeletePicture)

		// Фото
		adminHandler.POST("/pictures/:id/photo", r.doUploadMainPhoto)
		adminHandler.POST("/pictures/:id/gallery", r.doUploadGalleryPhoto)
		adminHandler.DELETE("/pictures/:id/gallery/:photo_id", r.doDeleteGalleryPhoto)
	}
}

// @Summary     Get pictures
// @Description Get all pictures
// @ID          get-pictures
// @Tags        pictures
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.Picture
// @Failure     500 {object} response
// @Router      /pictures [get]
func (p *picturesRoutes) doGetPictures(ctx *gin.Context) {
	pictures, err := p.u.GetPictures(ctx.Request.Context())
	if err != nil {
		p.l.Error(err, "http - v1 - doGetPictures")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, pictures)
}

// @Summary     Get picture by ID
// @Description Get picture by ID
// @ID          get-picture-by-id
// @Tags        pictures
// @Accept      json
// @Produce     json
// @Param       id path int true "Picture ID"
// @Success     200 {object} entity.Picture
// @Failure     400 {object} response
// @Failure     404 {object} response
// @Failure     500 {object} response
// @Router      /pictures/{id} [get]
func (p *picturesRoutes) doGetPictureByID(ctx *gin.Context) {
	id := ctx.Param("id")
	pictureID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	picture, err := p.u.GetPictureByID(ctx.Request.Context(), pictureID)
	if err != nil {
		if errors.Is(err, entity.ErrPictureNotFound) {
			errorResponse(ctx, http.StatusNotFound, "picture not found")
			return
		}
		p.l.Error(err, "http - v1 - doGetPictureByID")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, picture)
}

// @Summary     Create picture
// @Description Create new picture (without photos)
// @ID          create-picture
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body entity.PictureCreateRequest true "Picture data"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/pictures [post]
// @Security    BearerAuth
func (p *picturesRoutes) doCreatePicture(ctx *gin.Context) {
	var req entity.PictureCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		p.l.Error(err, "http - v1 - doCreatePicture")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := p.u.CreatePicture(ctx.Request.Context(), req); err != nil {
		p.l.Error(err, "http - v1 - doCreatePicture")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Update picture
// @Description Update picture data
// @ID          update-picture
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Picture ID"
// @Param       request body entity.PictureUpdateRequest true "Picture data"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/pictures/{id} [patch]
// @Security    BearerAuth
func (p *picturesRoutes) doUpdatePicture(ctx *gin.Context) {
	id := ctx.Param("id")
	pictureID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	var req entity.PictureUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		p.l.Error(err, "http - v1 - doUpdatePicture")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := p.u.UpdatePicture(ctx.Request.Context(), pictureID, req); err != nil {
		p.l.Error(err, "http - v1 - doUpdatePicture")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Delete picture
// @Description Delete picture by ID
// @ID          delete-picture
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Picture ID"
// @Success     200
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/pictures/{id} [delete]
// @Security    BearerAuth
func (p *picturesRoutes) doDeletePicture(ctx *gin.Context) {
	id := ctx.Param("id")
	pictureID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := p.u.DeletePicture(ctx.Request.Context(), pictureID); err != nil {
		p.l.Error(err, "http - v1 - doDeletePicture")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// @Summary     Upload main photo
// @Description Upload main photo for picture
// @ID          upload-main-photo
// @Tags        admin
// @Accept      multipart/form-data
// @Produce     json
// @Param       id path int true "Picture ID"
// @Param       file formData file true "Image file"
// @Success     200 {object} entity.PhotoUploadResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/pictures/{id}/photo [post]
// @Security    BearerAuth
func (p *picturesRoutes) doUploadMainPhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	pictureID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid picture ID")
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		p.l.Error(err, "http - v1 - doUploadMainPhoto")
		errorResponse(ctx, http.StatusBadRequest, "file is required")
		return
	}

	response, err := p.u.UploadPhoto(ctx.Request.Context(), file, entity.PhotoUploadRequest{
		PictureID: pictureID,
		IsMain:    true,
	})
	if err != nil {
		p.l.Error(err, "http - v1 - doUploadMainPhoto")
		errorResponse(ctx, http.StatusInternalServerError, "can't upload photo")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary     Upload gallery photo
// @Description Upload photo to picture gallery
// @ID          upload-gallery-photo
// @Tags        admin
// @Accept      multipart/form-data
// @Produce     json
// @Param       id path int true "Picture ID"
// @Param       file formData file true "Image file"
// @Success     200 {object} entity.PhotoUploadResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/pictures/{id}/gallery [post]
// @Security    BearerAuth
func (p *picturesRoutes) doUploadGalleryPhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	pictureID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid picture ID")
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		p.l.Error(err, "http - v1 - doUploadGalleryPhoto")
		errorResponse(ctx, http.StatusBadRequest, "file is required")
		return
	}

	response, err := p.u.UploadPhoto(ctx.Request.Context(), file, entity.PhotoUploadRequest{
		PictureID: pictureID,
		IsMain:    false,
	})
	if err != nil {
		p.l.Error(err, "http - v1 - doUploadGalleryPhoto")
		errorResponse(ctx, http.StatusInternalServerError, "can't upload photo")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary     Delete gallery photo
// @Description Delete photo from picture gallery
// @ID          delete-gallery-photo
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       id path int true "Picture ID"
// @Param       photo_id path int true "Photo ID"
// @Success     200 {object} entity.PhotoDeleteResponse
// @Failure     400 {object} response
// @Failure     401 {object} response
// @Failure     500 {object} response
// @Router      /admin/pictures/{id}/gallery/{photo_id} [delete]
// @Security    BearerAuth
func (p *picturesRoutes) doDeleteGalleryPhoto(ctx *gin.Context) {
	pictureID := ctx.Param("id")
	photoID := ctx.Param("photo_id")

	pID, err := strconv.ParseUint(pictureID, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid picture ID")
		return
	}

	phID, err := strconv.ParseUint(photoID, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid photo ID")
		return
	}

	response, err := p.u.DeletePhoto(ctx.Request.Context(), pID, phID)
	if err != nil {
		p.l.Error(err, "http - v1 - doDeleteGalleryPhoto")
		errorResponse(ctx, http.StatusInternalServerError, "can't delete photo")
		return
	}

	ctx.JSON(http.StatusOK, response)
}
