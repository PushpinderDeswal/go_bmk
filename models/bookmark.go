package models

import (
	"time"
)

// TODO: add category, priority, importance etc.
type Bookmark struct {
	ID        string     `json:"id" db:"id"`
	Url       string     `json:"url" db:"url"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}
