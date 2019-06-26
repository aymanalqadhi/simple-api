package services

import (
	"errors"
	"github.com/xSHAD0Wx/simple-api/models"
)

// MySQLClientsRepo is a test repositroy impelementing the IClientsRepo interface
type MySQLClientsRepo struct {
}

// GetByID gets a client by id
func (repo *MySQLClientsRepo) GetByID(id uint) (models.Client, error) {
	ret := models.Client{}

	if Services.DB.Find(&ret, id).Error != nil {
		return ret, errors.New("could not fetch client")
	}

	return ret, nil
}

// GetAll gets all items in the repository
func (repo *MySQLClientsRepo) GetAll() ([]models.Client, error) {
	ret := []models.Client{}

	if Services.DB.Find(&ret).Error != nil {
		return ret, errors.New("could not fetch client")
	}

	return ret, nil
}

// Add adds an item to the repository
func (repo *MySQLClientsRepo) Add(client models.Client) (uint, error) {
	if Services.DB.Create(&client).Error != nil {
		return 0, errors.New("could not add client")
	}

	return client.ID, nil
}

// AddAll adds a group of items to the repository
func (repo *MySQLClientsRepo) AddAll(clients []models.Client) (int, error) {
	succeded := 0
	var err error

	for _, item := range clients {
		_, err = repo.Add(item)
		if err != nil {
			succeded++
		}
	}

	return succeded, nil
}

// Remove : Removes an item from the repository
func (repo *MySQLClientsRepo) Remove(id uint) error {
	if Services.DB.Delete(models.Client{}, id).Error != nil {
		return errors.New("could not delete client")
	}
	return nil
}

// RemoveAll removes a group of items
func (repo *MySQLClientsRepo) RemoveAll(ids []uint) (int, error) {
	deleted := 0
	var err error

	for _, id := range ids {
		err = repo.Remove(id)
		if err != nil {
			deleted++
		}
	}

	return deleted, nil
}
