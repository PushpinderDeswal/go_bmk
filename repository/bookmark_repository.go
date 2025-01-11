package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PushpinderDeswal/go_bmk/models"
)

type BookmarkRepository interface {
	GetBookmark(ctx context.Context, id string) (*models.Bookmark, error)
	GetAllBookmarks(ctx context.Context) ([]models.Bookmark, error)
	AddBookmark(ctx context.Context, Bookmark *models.Bookmark) error
	DeleteBookmark(ctx context.Context, id string) error
}

type SQLiteBookmarkRepository struct {
	db *sql.DB
}

func MakeSQLiteBookmarkRepository(db *sql.DB) *SQLiteBookmarkRepository {
	return &SQLiteBookmarkRepository{db: db}
}

func (r *SQLiteBookmarkRepository) GetBookmark(ctx context.Context, id string) (*models.Bookmark, error) {
	query := `SELECT id, url, created_at FROM bookmarks WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var bookmark models.Bookmark
	err := row.Scan(&bookmark.ID, &bookmark.Url, &bookmark.CreatedAt)

	if err == nil {
		return &bookmark, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return nil, err
}

func (r *SQLiteBookmarkRepository) GetAllBookmarks(ctx context.Context) ([]models.Bookmark, error) {
	query := `SELECT id, url, created_at FROM bookmarks`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []models.Bookmark
	for rows.Next() {
		var bookmark models.Bookmark
		err := rows.Scan(&bookmark.ID, &bookmark.Url, &bookmark.CreatedAt)

		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func (r *SQLiteBookmarkRepository) AddBookmark(ctx context.Context, bookmark *models.Bookmark) error {
	query := `INSERT INTO bookmarks (id, url, created_at) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, bookmark.ID, bookmark.Url, bookmark.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteBookmarkRepository) DeleteBookmark(ctx context.Context, id string) error {
	query := `DELETE FROM bookmarks WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
