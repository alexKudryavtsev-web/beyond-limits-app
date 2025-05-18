package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
	"github.com/alexKudryavtsev-web/beyond-limits-app/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const (
	_defaultNewsListCap = 64
)

type NewsRepo struct {
	*postgres.Postgres
}

func NewNewsRepo(pg *postgres.Postgres) *NewsRepo {
	return &NewsRepo{pg}
}

func (r *NewsRepo) GetNews(ctx context.Context) ([]entity.News, error) {
	query, _, err := r.Builder.
		Select("id", "title", "content", "created_at").
		From("news").
		OrderBy("created_at DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	news := make([]entity.News, 0, _defaultNewsListCap)
	for rows.Next() {
		var n entity.News
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt); err != nil {
			return nil, fmt.Errorf("can't scan row: %w", err)
		}
		news = append(news, n)
	}

	return news, nil
}

func (r *NewsRepo) GetNewsByID(ctx context.Context, id uint64) (*entity.News, error) {
	query, args, err := r.Builder.
		Select("id", "title", "content", "created_at").
		From("news").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	var news entity.News
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&news.ID,
		&news.Title,
		&news.Content,
		&news.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrNewsNotFound
		}
		return nil, fmt.Errorf("can't scan row: %w", err)
	}

	return &news, nil
}

func (r *NewsRepo) CreateNews(ctx context.Context, req entity.NewsCreateRequest) error {
	query, args, err := r.Builder.
		Insert("news").
		Columns("title", "content").
		Values(req.Title, req.Content).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert news: %w", err)
	}

	return nil
}

func (r *NewsRepo) UpdateNews(ctx context.Context, id uint64, req entity.NewsUpdateRequest) error {
	builder := r.Builder.Update("news")

	if req.Title != nil {
		builder = builder.Set("title", *req.Title)
	}
	if req.Content != nil {
		builder = builder.Set("content", *req.Content)
	}

	builder = builder.Where(squirrel.Eq{"id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("can't build update query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("can't update news: %w", err)
	}

	return nil
}

func (r *NewsRepo) DeleteNews(ctx context.Context, id uint64) error {
	query, args, err := r.Builder.
		Delete("news").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = r.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't delete news: %w", err)
	}

	return nil
}
