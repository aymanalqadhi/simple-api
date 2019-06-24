package repos

import (
	"github.com/xSHAD0Wx/simple-api/models"
)

// IUsersRepo an interface to represent a users repository
type IUsersRepo interface {
	GetUser(string, string) (models.User, error)
	AddUser(string, string, string, uint) error
}
