package postgres

import (
	"database/sql"

	"github.com/hanmd82/gogists/pkg/models"
)

type GistModel struct {
	DB *sql.DB
}

func (m *GistModel) Insert(title, content, expiresAt string) (int, error) {
	return 0, nil
}

func (m *GistModel) Get(id int) (*models.Gist, error) {
	return nil, nil
}

func (m* GistModel) Latest() ([]*models.Gist, error) {
	return nil, nil
}
