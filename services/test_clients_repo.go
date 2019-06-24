package services

import (
	"errors"
	"log"

	"github.com/ahmetb/go-linq"
	"github.com/xSHAD0Wx/simple-api/models"
)

// TestClientsRepo is a test repositroy impelementing the IClientsRepo interface
type TestClientsRepo struct {
	clients []models.Client
}

// Initialize initalizes the repositroy
func (repo *TestClientsRepo) Initialize() error {
	return nil
}

// Dispose frees used resources
func (repo *TestClientsRepo) Dispose() {}

// GetByID gets a client by id
func (repo *TestClientsRepo) GetByID(id uint) (models.Client, error) {

	for _, item := range repo.clients {
		if item.ID == id {
			return item, nil
		}
	}

	var ret models.Client
	return ret, errors.New("coulnd not find item")
}

// GetAll gets all items in the repository
func (repo *TestClientsRepo) GetAll() ([]models.Client, error) {
	return repo.clients, nil
}

// Add adds an item to the repository
func (repo *TestClientsRepo) Add(client models.Client) (uint, error) {
	repo.clients = append(repo.clients, client)

	log.Printf("Added a client, client id: %d\n", len(repo.clients))
	return uint(len(repo.clients)), nil
}

// AddAll adds a group of items to the repository
func (repo *TestClientsRepo) AddAll(clients []models.Client) (int, error) {
	repo.clients = append(repo.clients, clients...)
	log.Printf("Added %d clients.\n", len(clients))
	return len(clients), nil
}

// Remove : Removes an item from the repository
func (repo *TestClientsRepo) Remove(client models.Client) error {
	for i, c := range repo.clients {
		if client.ID == c.ID {
			repo.clients[len(repo.clients)-1], repo.clients[i] = repo.clients[i], repo.clients[len(repo.clients)-1]
			return nil
		}
	}

	return errors.New("no such item")
}

// RemoveAll removes a group of items
func (repo *TestClientsRepo) RemoveAll(ids []uint) (int, error) {
	removed := 0

	linq.From(ids).ForEach(func(id interface{}) {
		oldLen := len(repo.clients)

		linq.From(repo.clients).Where(func(c interface{}) bool {
			return id != (c.(models.Client)).ID
		}).ToSlice(&repo.clients)

		if oldLen == len(repo.clients)-1 {
			removed++
		}
	})

	return removed, nil
}
