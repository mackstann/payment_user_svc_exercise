package store

import (
	"github.com/mackstann/payment_user_svc_exercise/internal/models"
)

type Store struct {
	users map[string]models.User
}

func NewStore() Store {
	return Store{
		users: make(map[string]models.User),
	}
}

func (store Store) CreateUser(models.User) (id string, err error) {
	return "", nil
}

func (store Store) GetUser(id string) (models.User, error) {
	return models.User{}, nil
}
