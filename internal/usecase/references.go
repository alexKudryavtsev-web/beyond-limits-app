package usecase

import (
	"context"
	"fmt"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
)

type ReferencesUseCase struct {
	repo ReferencesRepo
}

func NewReferencesUseCase(repo ReferencesRepo) *ReferencesUseCase {
	return &ReferencesUseCase{repo: repo}
}

var _ References = (*ReferencesUseCase)(nil)

func (r *ReferencesUseCase) GetGenres(ctx context.Context) ([]entity.Genre, error) {
	genres, err := r.repo.GetGenres(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get genres: %w", err)
	}
	return genres, nil
}

func (r *ReferencesUseCase) CreateGenre(ctx context.Context, name string) error {
	if err := r.repo.CreateGenre(ctx, name); err != nil {
		return fmt.Errorf("can't create genre: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) DeleteGenre(ctx context.Context, id uint64) error {
	if err := r.repo.DeleteGenre(ctx, id); err != nil {
		return fmt.Errorf("can't delete genre: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) GetAuthors(ctx context.Context) ([]entity.Author, error) {
	authors, err := r.repo.GetAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get authors: %w", err)
	}
	return authors, nil
}

func (r *ReferencesUseCase) CreateAuthor(ctx context.Context, fullName string) error {
	if err := r.repo.CreateAuthor(ctx, fullName); err != nil {
		return fmt.Errorf("can't create author: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) DeleteAuthor(ctx context.Context, id uint64) error {
	if err := r.repo.DeleteAuthor(ctx, id); err != nil {
		return fmt.Errorf("can't delete author: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) GetDimensions(ctx context.Context) ([]entity.Dimension, error) {
	dimensions, err := r.repo.GetDimensions(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get dimensions: %w", err)
	}
	return dimensions, nil
}

func (r *ReferencesUseCase) CreateDimension(ctx context.Context, width, height int) error {
	if err := r.repo.CreateDimension(ctx, width, height); err != nil {
		return fmt.Errorf("can't create dimension: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) DeleteDimension(ctx context.Context, id uint64) error {
	if err := r.repo.DeleteDimension(ctx, id); err != nil {
		return fmt.Errorf("can't delete dimension: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) GetWorkTechniques(ctx context.Context) ([]entity.WorkTechnique, error) {
	techniques, err := r.repo.GetWorkTechniques(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get work techniques: %w", err)
	}
	return techniques, nil
}

func (r *ReferencesUseCase) CreateWorkTechnique(ctx context.Context, name string) error {
	if err := r.repo.CreateWorkTechnique(ctx, name); err != nil {
		return fmt.Errorf("can't create work technique: %w", err)
	}
	return nil
}

func (r *ReferencesUseCase) DeleteWorkTechnique(ctx context.Context, id uint64) error {
	if err := r.repo.DeleteWorkTechnique(ctx, id); err != nil {
		return fmt.Errorf("can't delete work technique: %w", err)
	}
	return nil
}
