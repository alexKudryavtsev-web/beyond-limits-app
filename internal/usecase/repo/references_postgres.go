package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/postgres"
)

type ReferencesRepo struct {
	*postgres.Postgres
}

func NewReferencesRepo(pg *postgres.Postgres) *ReferencesRepo {
	return &ReferencesRepo{pg}
}

// Genres
func (r *ReferencesRepo) GetGenres(ctx context.Context) ([]entity.Genre, error) {
	query, _, err := r.Builder.Select("id", "name").From("genres").ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	genres := make([]entity.Genre, 0, _defaultListCap)
	for rows.Next() {
		var genre entity.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, fmt.Errorf("can't scan row: %w", err)
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *ReferencesRepo) CreateGenre(ctx context.Context, name string) error {
	query, args, err := r.Builder.
		Insert("genres").
		Columns("name").
		Values(name).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert genre: %w", err)
	}

	return nil
}

func (r *ReferencesRepo) DeleteGenre(ctx context.Context, id uint64) error {
	query, args, err := r.Builder.
		Delete("genres").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't delete genre: %w", err)
	}

	return nil
}

// internal/repo/references_postgres.go

// Authors
func (r *ReferencesRepo) GetAuthors(ctx context.Context) ([]entity.Author, error) {
	query, _, err := r.Builder.Select("id", "full_name").From("authors").ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	authors := make([]entity.Author, 0, _defaultListCap)
	for rows.Next() {
		var author entity.Author
		if err := rows.Scan(&author.ID, &author.FullName); err != nil {
			return nil, fmt.Errorf("can't scan row: %w", err)
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *ReferencesRepo) CreateAuthor(ctx context.Context, fullName string) error {
	query, args, err := r.Builder.
		Insert("authors").
		Columns("full_name").
		Values(fullName).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert author: %w", err)
	}

	return nil
}

func (r *ReferencesRepo) DeleteAuthor(ctx context.Context, id uint64) error {
	query, args, err := r.Builder.
		Delete("authors").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't delete author: %w", err)
	}

	return nil
}

// Dimensions
func (r *ReferencesRepo) GetDimensions(ctx context.Context) ([]entity.Dimension, error) {
	query, _, err := r.Builder.Select("id", "width", "height").From("dimensions").ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	dimensions := make([]entity.Dimension, 0, _defaultListCap)
	for rows.Next() {
		var dimension entity.Dimension
		if err := rows.Scan(&dimension.ID, &dimension.Width, &dimension.Height); err != nil {
			return nil, fmt.Errorf("can't scan row: %w", err)
		}
		dimensions = append(dimensions, dimension)
	}

	return dimensions, nil
}

func (r *ReferencesRepo) CreateDimension(ctx context.Context, width, height int) error {
	query, args, err := r.Builder.
		Insert("dimensions").
		Columns("width", "height").
		Values(width, height).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert dimension: %w", err)
	}

	return nil
}

func (r *ReferencesRepo) DeleteDimension(ctx context.Context, id uint64) error {
	query, args, err := r.Builder.
		Delete("dimensions").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't delete dimension: %w", err)
	}

	return nil
}

// WorkTechniques
func (r *ReferencesRepo) GetWorkTechniques(ctx context.Context) ([]entity.WorkTechnique, error) {
	query, _, err := r.Builder.Select("id", "name").From("work_techniques").ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	techniques := make([]entity.WorkTechnique, 0, _defaultListCap)
	for rows.Next() {
		var technique entity.WorkTechnique
		if err := rows.Scan(&technique.ID, &technique.Name); err != nil {
			return nil, fmt.Errorf("can't scan row: %w", err)
		}
		techniques = append(techniques, technique)
	}

	return techniques, nil
}

func (r *ReferencesRepo) CreateWorkTechnique(ctx context.Context, name string) error {
	query, args, err := r.Builder.
		Insert("work_techniques").
		Columns("name").
		Values(name).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert work technique: %w", err)
	}

	return nil
}

func (r *ReferencesRepo) DeleteWorkTechnique(ctx context.Context, id uint64) error {
	query, args, err := r.Builder.
		Delete("work_techniques").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't delete work technique: %w", err)
	}

	return nil
}
