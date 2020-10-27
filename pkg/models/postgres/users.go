package postgres

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/hanmd82/gogists/pkg/models"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	sqlStatement := `
		INSERT INTO users (name, email, hashed_password, created_at)
		VALUES ($1, $2, $3, now() at time zone 'utc')
		RETURNING id`

	var id int
	err = m.DB.QueryRow(sqlStatement, name, email, hashedPassword).Scan(&id)
	if err != nil {
		var pqError *pq.Error
		if errors.As(err, &pqError) {
			if pqError.Code.Name() == "unique_violation" && strings.Contains(pqError.Message, "users_uc_email") {
				//return 0, errors.New(fmt.Sprintf("%s, %s", pqError.Code, pqError.Constraint))
				return 0, models.ErrDuplicateEmail
			}
		}
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
