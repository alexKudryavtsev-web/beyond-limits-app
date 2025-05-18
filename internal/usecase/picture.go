package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
	"github.com/google/uuid"
)

type PicturesUseCase struct {
	repo PicturesRepo
}

var _ Pictures = (*PicturesUseCase)(nil)

func NewPicturesUseCase(repo PicturesRepo) *PicturesUseCase {
	return &PicturesUseCase{repo: repo}
}

func (uc *PicturesUseCase) GetPictures(ctx context.Context) ([]entity.Picture, error) {
	pictures, err := uc.repo.GetPictures(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get pictures: %w", err)
	}
	return pictures, nil
}

func (uc *PicturesUseCase) GetPictureByID(ctx context.Context, id uint64) (*entity.Picture, error) {
	picture, err := uc.repo.GetPictureByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't get picture by id: %w", err)
	}
	return picture, nil
}

func (uc *PicturesUseCase) CreatePicture(ctx context.Context, req entity.PictureCreateRequest) error {
	if err := uc.repo.CreatePicture(ctx, req); err != nil {
		return fmt.Errorf("can't create picture: %w", err)
	}
	return nil
}

func (uc *PicturesUseCase) UpdatePicture(ctx context.Context, id uint64, req entity.PictureUpdateRequest) error {
	if err := uc.repo.UpdatePicture(ctx, id, req); err != nil {
		return fmt.Errorf("can't update picture: %w", err)
	}
	return nil
}

func (uc *PicturesUseCase) DeletePicture(ctx context.Context, id uint64) error {
	if err := uc.repo.DeletePicture(ctx, id); err != nil {
		return fmt.Errorf("can't delete picture: %w", err)
	}
	return nil
}

func (uc *PicturesUseCase) UploadPhoto(
	ctx context.Context,
	fileHeader *multipart.FileHeader,
	req entity.PhotoUploadRequest,
) (*entity.PhotoUploadResponse, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("can't open uploaded file: %w", err)
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String(), ext)
	filePath := filepath.Join("uploads", filename)

	if err := os.MkdirAll("uploads", 0755); err != nil {
		return nil, fmt.Errorf("can't create uploads directory: %w", err)
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("can't create file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return nil, fmt.Errorf("can't save file: %w", err)
	}

	mime := fileHeader.Header.Get("Content-Type")
	photoID, err := uc.repo.SavePhoto(ctx, req.PictureID, "/"+filePath, mime, req.IsMain)
	if err != nil {
		os.Remove(filePath)
		return nil, fmt.Errorf("can't save photo info: %w", err)
	}

	return &entity.PhotoUploadResponse{
		ID:  photoID,
		URL: "/" + filePath,
	}, nil
}

func (uc *PicturesUseCase) DeletePhoto(
	ctx context.Context,
	pictureID, photoID uint64,
) (*entity.PhotoDeleteResponse, error) {
	photo, err := uc.repo.GetPhoto(ctx, photoID)
	if err != nil {
		return nil, fmt.Errorf("can't get photo info: %w", err)
	}

	if err := uc.repo.DeletePhoto(ctx, photoID); err != nil {
		return nil, fmt.Errorf("can't delete photo from db: %w", err)
	}

	if err := os.Remove(strings.TrimPrefix(photo.URL, "/")); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("can't delete photo file: %w", err)
	}

	return &entity.PhotoDeleteResponse{Success: true}, nil
}
