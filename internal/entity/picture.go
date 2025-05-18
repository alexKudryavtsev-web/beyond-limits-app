package entity

import (
	"errors"
	"time"
)

type Picture struct {
	ID            uint64        `json:"id"`
	Title         string        `json:"title"`
	Price         int           `json:"price"`
	Author        Author        `json:"author"`
	Dimensions    Dimension     `json:"dimensions"`
	WorkTechnique WorkTechnique `json:"work_technique"`
	Genre         Genre         `json:"genre"`
	Photo         Photo         `json:"photo"`
	Gallery       []Photo       `json:"gallery"`
	CreatedAt     time.Time     `json:"created_at"`
}

type Photo struct {
	ID   uint64 `json:"id"`
	URL  string `json:"url"`
	Mime string `json:"mime"`
}

type PictureCreateRequest struct {
	Title           string `json:"title" binding:"required"`
	Price           int    `json:"price" binding:"required"`
	AuthorID        uint64 `json:"author_id" binding:"required"`
	DimensionsID    uint64 `json:"dimensions_id" binding:"required"`
	WorkTechniqueID uint64 `json:"work_technique_id" binding:"required"`
	GenreID         uint64 `json:"genre_id" binding:"required"`
}

type PictureUpdateRequest struct {
	Title           *string `json:"title"`
	Price           *int    `json:"price"`
	AuthorID        *uint64 `json:"author_id"`
	DimensionsID    *uint64 `json:"dimensions_id"`
	WorkTechniqueID *uint64 `json:"work_technique_id"`
	GenreID         *uint64 `json:"genre_id"`
}

type PhotoUploadRequest struct {
	PictureID uint64 `form:"picture_id" binding:"required"`
	IsMain    bool   `form:"is_main"`
}

type PhotoUploadResponse struct {
	ID  uint64 `json:"id"`
	URL string `json:"url"`
}

type PhotoDeleteResponse struct {
	Success bool `json:"success"`
}

var (
	ErrPictureNotFound = errors.New("picture not found")
	ErrPhotoNotFound   = errors.New("photo not found")
)
