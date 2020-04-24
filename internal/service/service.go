package service

import (
	"github.com/mackstann/payment_user_svc_exercise/internal/models"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
)

type Service struct {
	store store.Store
}

func NewService(store store.Store) Service {
	return Service{
		store: store,
	}
}

func (svc Service) CreateUser(models.User) (id string, err error) {
	return "", nil
}

func (svc Service) GetUser(id string) (models.User, error) {
	return models.User{}, nil
}
