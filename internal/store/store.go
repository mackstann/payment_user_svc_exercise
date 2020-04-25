package store

import (
	"github.com/google/uuid"
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

func (store Store) CreateUser(user models.User) (id string, err error) {
	uu, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	user.ID = uu.String()
	store.users[user.ID] = user
	return user.ID, nil
}

func (store Store) GetUser(id string) (models.User, error) {
	user, ok := store.users[id]
	if !ok {
		return models.User{}, UserNotFoundError
	}
	return user, nil
}
