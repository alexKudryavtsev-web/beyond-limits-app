// internal/usecase/news.go
package usecase

import (
	"context"
	"fmt"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
)

type NewsUseCase struct {
	repo NewsRepo
}

var _ News = (*NewsUseCase)(nil)

func NewNewsUseCase(repo NewsRepo) *NewsUseCase {
	return &NewsUseCase{repo: repo}
}

func (uc *NewsUseCase) GetNews(ctx context.Context) ([]entity.News, error) {
	news, err := uc.repo.GetNews(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't get news: %w", err)
	}
	return news, nil
}

func (uc *NewsUseCase) GetNewsByID(ctx context.Context, id uint64) (*entity.News, error) {
	news, err := uc.repo.GetNewsByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't get news by id: %w", err)
	}
	return news, nil
}

func (uc *NewsUseCase) CreateNews(ctx context.Context, req entity.NewsCreateRequest) error {
	if err := uc.repo.CreateNews(ctx, req); err != nil {
		return fmt.Errorf("can't create news: %w", err)
	}
	return nil
}

func (uc *NewsUseCase) UpdateNews(ctx context.Context, id uint64, req entity.NewsUpdateRequest) error {
	if err := uc.repo.UpdateNews(ctx, id, req); err != nil {
		return fmt.Errorf("can't update news: %w", err)
	}
	return nil
}

func (uc *NewsUseCase) DeleteNews(ctx context.Context, id uint64) error {
	if err := uc.repo.DeleteNews(ctx, id); err != nil {
		return fmt.Errorf("can't delete news: %w", err)
	}
	return nil
}
