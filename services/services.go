package services

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/xSHAD0Wx/simple-api/models"
	"github.com/xSHAD0Wx/simple-api/repos"
	"github.com/xSHAD0Wx/simple-api/shared"
)

// AppServices is a struct to represent a services container (DI emulation)
type AppServices struct {
	DB          *gorm.DB
	ClientsRepo repos.IClientsRepo
	UsersRepo   repos.IUsersRepo
	AuthService IAuthService
}

// Services is the app services container
var Services *AppServices

// Inits databse connection
func initDatabase() error {
	var err error
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", shared.DBUsername, shared.DBPassword, shared.DBaseHost, shared.DBPort, shared.DBName, shared.DBOptions)
	Services.DB, err = gorm.Open(shared.DBDriver, connString)

	if err != nil {
		return err
	}

	log.Printf("Connected to %s:%d", shared.DBaseHost, shared.ListenPort)
	log.Printf("Migrating DB: %s", shared.DBName)

	return Services.DB.AutoMigrate(&models.Client{}, &models.User{}).Error
}

// Inits services
func init() {
	Services = &AppServices{
		// Add services
		ClientsRepo: &MySQLClientsRepo{},
		UsersRepo:   &MySQLUsersRepo{},
		AuthService: &JwtAuthService{},
	}

	// Init Services
	if err := initDatabase(); err != nil {
		panic(err.Error())
	}
}
