package postgres

import (
	"database/sql"

	"github.com/hanmd82/gogists/pkg/models"
)

type GistModel struct {
	DB *sql.DB
}

func (m *GistModel) Insert(title, content, expiresInDays string) (int, error) {
	sqlStatement := `
		INSERT INTO gists (title, content, created_at, expires_at)
		VALUES ($1, $2, now() at time zone 'utc', now() at time zone 'utc' + $3 * INTERVAL '1 day')
		RETURNING id`

	var id int
	err := m.DB.QueryRow(sqlStatement, title, content, expiresInDays).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *GistModel) Get(id int) (*models.Gist, error) {
	return nil, nil
}

func (m* GistModel) Latest() ([]*models.Gist, error) {
	return nil, nil
}
