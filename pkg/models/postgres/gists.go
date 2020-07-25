package postgres

import (
	"database/sql"
	"errors"

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
	sqlStatement := `
		SELECT id, title, content, created_at, expires_at FROM gists
		WHERE expires_at > now() at time zone 'utc' AND id = $1`
	row := m.DB.QueryRow(sqlStatement, id)

	gist := &models.Gist{}
	// Copy the values from each field in sql.Row to the corresponding field in the Gist struct
	err := row.Scan(&gist.ID, &gist.Title, &gist.Content, &gist.CreatedAt, &gist.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return gist, nil
}

func (m* GistModel) Latest() ([]*models.Gist, error) {
	sqlStatement := `
		SELECT id, title, content, created_at, expires_at FROM gists
		WHERE expires_at > now() at time zone 'utc' ORDER BY created_at	DESC LIMIT 10`

	rows, err := m.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gists := []*models.Gist{}
	for rows.Next() {
		g := &models.Gist{}
		err = rows.Scan(&g.ID, &g.Title, &g.Content, &g.CreatedAt, &g.ExpiresAt)
		if err != nil {
			return nil, err
		}
		gists = append(gists, g)
	}

	// check for any error encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gists, nil
}
