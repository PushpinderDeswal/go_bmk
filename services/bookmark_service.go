package services

import (
	"context"
	"time"

	"github.com/PushpinderDeswal/go_bmk/models"
	"github.com/PushpinderDeswal/go_bmk/repository"
	"github.com/google/uuid"
)

type BookmarkService struct {
	repo repository.BookmarkRepository
}

func (s *BookmarkService) AddBookmark(ctx *context.Context, url string) error {
	now := time.Now()

	bookmark := &models.Bookmark{
		ID:        uuid.New().String(),
		Url:       url,
		CreatedAt: &now,
	}

	return s.repo.AddBookmark(ctx, bookmark)
}

func (s *BookmarkService) GetBookmark(ctx *context.Context, id string) (*models.Bookmark, error) {
	return s.repo.GetBookmark(ctx, id)
}

func (s *BookmarkService) GetAllBookmarks(ctx *context.Context, id string) ([]models.Bookmark, error) {
	return s.repo.GetAllBookmarks(ctx)
}

func (s *BookmarkService) DeleteBookmark(ctx *context.Context, id string) error {
	return s.repo.DeleteBookmark(ctx, id)
}
