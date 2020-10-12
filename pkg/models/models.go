package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Gist struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (g Gist) FormatTime(t time.Time) string {
	return t.Format(time.RFC822)
}
