package repos

import (
	"github.com/xSHAD0Wx/simple-api/models"
)

// IClientsRepo : An interface with the signatures of common repository methods
type IClientsRepo interface {
	GetByID(uint) (models.Client, error)
	GetAll() ([]models.Client, error)
	Add(models.Client) (uint, error)
	AddAll([]models.Client) (int, error)
	Remove(uint) error
	RemoveAll([]uint) (int, error)
}
