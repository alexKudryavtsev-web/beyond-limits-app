package entity

import (
	"errors"
	"time"
)

type News struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type NewsCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type NewsUpdateRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

var (
	ErrNewsNotFound = errors.New("news not found")
)
