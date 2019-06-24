package services

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/xSHAD0Wx/simple-api/models"
)

// MySQLUsersRepo is a struct that implements IUsersRepo interface
type MySQLUsersRepo struct {
}

// GetUser gets a user by username and password
func (repo *MySQLUsersRepo) GetUser(username string, password string) (models.User, error) {
	ret := models.User{}
	if err := Services.DB.Find(&ret, "username = ? AND password = ?", username, password).Error; err != nil {
		return ret, err
	}

	return ret, nil
}

// AddUser registers user with username, password, and auth level
func (repo *MySQLUsersRepo) AddUser(username string, email string, password string, authLevel uint) error {
	hasher := sha256.New()
	hasher.Write([]byte(password))

	user := models.User{
		Username:  username,
		Email:     email,
		Password:  hex.EncodeToString(hasher.Sum(nil)),
		AuthLevel: authLevel,
	}

	return Services.DB.Create(&user).Error
}
