package mock

import (
	"time"

	"github.com/hanmd82/gogists/pkg/models"
)

var mockUser = &models.User{
	ID:        1,
	Name:      "Alice",
	Email:     "alice@example.com",
	CreatedAt: time.Now(),
	Active:    true,
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) (int, error) {
	switch email {
	case "dupe@example.com":
		return 0, models.ErrDuplicateEmail
	default:
		return 2, nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "alice@example.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
