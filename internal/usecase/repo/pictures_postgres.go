package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type PicturesRepo struct {
	*postgres.Postgres
}

func NewPicturesRepo(pg *postgres.Postgres) *PicturesRepo {
	return &PicturesRepo{pg}
}

func (r *PicturesRepo) GetPictures(ctx context.Context) ([]entity.Picture, error) {
	sql := `
	SELECT 
		p.id, p.title, p.price, p.created_at,
		a.id, a.full_name,
		d.id, d.width, d.height,
		wt.id, wt.name,
		g.id, g.name,
		pp.id, pp.url, pp.mime
	FROM pictures p
	JOIN authors a ON p.author_id = a.id
	JOIN dimensions d ON p.dimensions_id = d.id
	JOIN work_techniques wt ON p.work_technique_id = wt.id
	JOIN genres g ON p.genre_id = g.id
	JOIN pictures_photos pp ON p.id = pp.picture_id AND pp.is_main = true
	`

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("can't query pictures: %w", err)
	}
	defer rows.Close()

	var pictures []entity.Picture
	for rows.Next() {
		var pic entity.Picture
		err := rows.Scan(
			&pic.ID, &pic.Title, &pic.Price, &pic.CreatedAt,
			&pic.Author.ID, &pic.Author.FullName,
			&pic.Dimensions.ID, &pic.Dimensions.Width, &pic.Dimensions.Height,
			&pic.WorkTechnique.ID, &pic.WorkTechnique.Name,
			&pic.Genre.ID, &pic.Genre.Name,
			&pic.Photo.ID, &pic.Photo.URL, &pic.Photo.Mime,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan picture: %w", err)
		}

		gallery, err := r.getPictureGallery(ctx, pic.ID)
		if err != nil {
			return nil, err
		}
		pic.Gallery = gallery

		pictures = append(pictures, pic)
	}

	return pictures, nil
}

func (r *PicturesRepo) getPictureGallery(ctx context.Context, pictureID uint64) ([]entity.Photo, error) {
	sql := `
	SELECT id, url, mime 
	FROM pictures_photos 
	WHERE picture_id = $1 AND is_main = false
	`

	rows, err := r.Pool.Query(ctx, sql, pictureID)
	if err != nil {
		return nil, fmt.Errorf("can't query gallery: %w", err)
	}
	defer rows.Close()

	var gallery []entity.Photo
	for rows.Next() {
		var photo entity.Photo
		if err := rows.Scan(&photo.ID, &photo.URL, &photo.Mime); err != nil {
			return nil, fmt.Errorf("can't scan gallery photo: %w", err)
		}
		gallery = append(gallery, photo)
	}

	return gallery, nil
}

func (r *PicturesRepo) GetPictureByID(ctx context.Context, id uint64) (*entity.Picture, error) {
	sql := `
	SELECT 
		p.id, p.title, p.price, p.created_at,
		a.id, a.full_name,
		d.id, d.width, d.height,
		wt.id, wt.name,
		g.id, g.name,
		pp.id, pp.url, pp.mime
	FROM pictures p
	JOIN authors a ON p.author_id = a.id
	JOIN dimensions d ON p.dimensions_id = d.id
	JOIN work_techniques wt ON p.work_technique_id = wt.id
	JOIN genres g ON p.genre_id = g.id
	JOIN pictures_photos pp ON p.id = pp.picture_id AND pp.is_main = true
	WHERE p.id = $1
	`

	var pic entity.Picture
	err := r.Pool.QueryRow(ctx, sql, id).Scan(
		&pic.ID, &pic.Title, &pic.Price, &pic.CreatedAt,
		&pic.Author.ID, &pic.Author.FullName,
		&pic.Dimensions.ID, &pic.Dimensions.Width, &pic.Dimensions.Height,
		&pic.WorkTechnique.ID, &pic.WorkTechnique.Name,
		&pic.Genre.ID, &pic.Genre.Name,
		&pic.Photo.ID, &pic.Photo.URL, &pic.Photo.Mime,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrPictureNotFound
		}
		return nil, fmt.Errorf("can't get picture by id: %w", err)
	}

	gallery, err := r.getPictureGallery(ctx, id)
	if err != nil {
		return nil, err
	}
	pic.Gallery = gallery

	return &pic, nil
}

func (r *PicturesRepo) CreatePicture(ctx context.Context, req entity.PictureCreateRequest) error {
	sql := `
	INSERT INTO pictures (title, price, author_id, dimensions_id, work_technique_id, genre_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.Pool.Exec(ctx, sql,
		req.Title,
		req.Price,
		req.AuthorID,
		req.DimensionsID,
		req.WorkTechniqueID,
		req.GenreID,
	)
	if err != nil {
		return fmt.Errorf("can't create picture: %w", err)
	}

	return nil
}

func (r *PicturesRepo) UpdatePicture(ctx context.Context, id uint64, req entity.PictureUpdateRequest) error {
	builder := r.Builder.Update("pictures")

	if req.Title != nil {
		builder = builder.Set("title", *req.Title)
	}
	if req.Price != nil {
		builder = builder.Set("price", *req.Price)
	}
	if req.AuthorID != nil {
		builder = builder.Set("author_id", *req.AuthorID)
	}
	if req.DimensionsID != nil {
		builder = builder.Set("dimensions_id", *req.DimensionsID)
	}
	if req.WorkTechniqueID != nil {
		builder = builder.Set("work_technique_id", *req.WorkTechniqueID)
	}
	if req.GenreID != nil {
		builder = builder.Set("genre_id", *req.GenreID)
	}

	builder = builder.Where(squirrel.Eq{"id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("can't build update query: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("can't update picture: %w", err)
	}

	return nil
}

func (r *PicturesRepo) DeletePicture(ctx context.Context, id uint64) error {
	sql := "DELETE FROM pictures WHERE id = $1"

	_, err := r.Pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("can't delete picture: %w", err)
	}

	return nil
}

func (r *PicturesRepo) SavePhoto(
	ctx context.Context,
	pictureID uint64,
	url, mime string,
	isMain bool,
) (uint64, error) {
	fmt.Println(pictureID, url, mime, isMain)
	sql := `
	INSERT INTO pictures_photos (picture_id, url, mime, is_main)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	var id uint64
	err := r.Pool.QueryRow(ctx, sql, pictureID, url, mime, isMain).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't save photo: %w", err)
	}

	return id, nil
}

func (r *PicturesRepo) DeletePhoto(ctx context.Context, photoID uint64) error {
	sql := "DELETE FROM pictures_photos WHERE id = $1"

	_, err := r.Pool.Exec(ctx, sql, photoID)
	if err != nil {
		return fmt.Errorf("can't delete photo: %w", err)
	}

	return nil
}

func (r *PicturesRepo) GetPhoto(ctx context.Context, photoID uint64) (*entity.Photo, error) {
	sql := "SELECT id, url, mime FROM pictures_photos WHERE id = $1"

	var photo entity.Photo
	err := r.Pool.QueryRow(ctx, sql, photoID).Scan(&photo.ID, &photo.URL, &photo.Mime)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrPhotoNotFound
		}
		return nil, fmt.Errorf("can't get photo: %w", err)
	}

	return &photo, nil
}
