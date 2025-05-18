package usecase

import (
	"context"

	"github.com/alexKudryavtsev-web/beyond-limits-app/internal/entity"
)

type (
	Todos interface {
		Todos(ctx context.Context) ([]entity.Todo, error)
		TodoByID(ctx context.Context, id uint64) (*entity.Todo, error)
		SaveTodo(ctx context.Context, task string) error
	}

	TodosRepo interface {
		GetAllTodos(ctx context.Context) ([]entity.Todo, error)
		GetTodoByID(ctx context.Context, id uint64) (*entity.Todo, error)
		SaveTodo(ctx context.Context, task string) error
	}

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
)
