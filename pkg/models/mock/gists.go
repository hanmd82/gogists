package mock

import (
	"time"

	"github.com/hanmd82/gogists/pkg/models"
)

var mockGist = &models.Gist{
	ID:        1,
	Title:     "An old silent pond",
	Content:   "An old silent pond...",
	CreatedAt: time.Now(),
	ExpiresAt: time.Now(),
}

type GistModel struct{}

func (m *GistModel) Insert(title, content, expires_at string) (int, error) {
	return 2, nil
}

func (m *GistModel) Get(id int) (*models.Gist, error) {
	switch id {
	case 1:
		return mockGist, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *GistModel) Latest() ([]*models.Gist, error) {
	return []*models.Gist{mockGist}, nil
}
