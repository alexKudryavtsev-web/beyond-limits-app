package usecase

import (
	"context"
	"mime/multipart"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
)

type (
	Auth interface {
		Login(ctx context.Context, login, password string) (string, error)
	}

	References interface {
		Genres
		Authors
		Dimensions
		WorkTechniques
	}

	Genres interface {
		GetGenres(ctx context.Context) ([]entity.Genre, error)
		CreateGenre(ctx context.Context, name string) error
		DeleteGenre(ctx context.Context, id uint64) error
	}

	Authors interface {
		GetAuthors(ctx context.Context) ([]entity.Author, error)
		CreateAuthor(ctx context.Context, fullName string) error
		DeleteAuthor(ctx context.Context, id uint64) error
	}

	Dimensions interface {
		GetDimensions(ctx context.Context) ([]entity.Dimension, error)
		CreateDimension(ctx context.Context, width, height int) error
		DeleteDimension(ctx context.Context, id uint64) error
	}

	WorkTechniques interface {
		GetWorkTechniques(ctx context.Context) ([]entity.WorkTechnique, error)
		CreateWorkTechnique(ctx context.Context, name string) error
		DeleteWorkTechnique(ctx context.Context, id uint64) error
	}

	ReferencesRepo interface {
		GetGenres(ctx context.Context) ([]entity.Genre, error)
		CreateGenre(ctx context.Context, name string) error
		DeleteGenre(ctx context.Context, id uint64) error

		GetAuthors(ctx context.Context) ([]entity.Author, error)
		CreateAuthor(ctx context.Context, fullName string) error
		DeleteAuthor(ctx context.Context, id uint64) error

		GetDimensions(ctx context.Context) ([]entity.Dimension, error)
		CreateDimension(ctx context.Context, width, height int) error
		DeleteDimension(ctx context.Context, id uint64) error

		GetWorkTechniques(ctx context.Context) ([]entity.WorkTechnique, error)
		CreateWorkTechnique(ctx context.Context, name string) error
		DeleteWorkTechnique(ctx context.Context, id uint64) error
	}

	Pictures interface {
		GetPictures(ctx context.Context) ([]entity.Picture, error)
		GetPictureByID(ctx context.Context, id uint64) (*entity.Picture, error)
		CreatePicture(ctx context.Context, req entity.PictureCreateRequest) error
		UpdatePicture(ctx context.Context, id uint64, req entity.PictureUpdateRequest) error
		DeletePicture(ctx context.Context, id uint64) error
		UploadPhoto(ctx context.Context, fileHeader *multipart.FileHeader, req entity.PhotoUploadRequest) (*entity.PhotoUploadResponse, error)
		DeletePhoto(ctx context.Context, pictureID, photoID uint64) (*entity.PhotoDeleteResponse, error)
	}

	PicturesRepo interface {
		GetPictures(ctx context.Context) ([]entity.Picture, error)
		GetPictureByID(ctx context.Context, id uint64) (*entity.Picture, error)
		CreatePicture(ctx context.Context, req entity.PictureCreateRequest) error
		UpdatePicture(ctx context.Context, id uint64, req entity.PictureUpdateRequest) error
		DeletePicture(ctx context.Context, id uint64) error
		SavePhoto(ctx context.Context, pictureID uint64, url, mime string, isMain bool) (uint64, error)
		DeletePhoto(ctx context.Context, photoID uint64) error
		GetPhoto(ctx context.Context, photoID uint64) (*entity.Photo, error)
	}

	News interface {
		GetNews(ctx context.Context) ([]entity.News, error)
		GetNewsByID(ctx context.Context, id uint64) (*entity.News, error)
		CreateNews(ctx context.Context, req entity.NewsCreateRequest) error
		UpdateNews(ctx context.Context, id uint64, req entity.NewsUpdateRequest) error
		DeleteNews(ctx context.Context, id uint64) error
	}

	NewsRepo interface {
		GetNews(ctx context.Context) ([]entity.News, error)
		GetNewsByID(ctx context.Context, id uint64) (*entity.News, error)
		CreateNews(ctx context.Context, req entity.NewsCreateRequest) error
		UpdateNews(ctx context.Context, id uint64, req entity.NewsUpdateRequest) error
		DeleteNews(ctx context.Context, id uint64) error
	}
)
