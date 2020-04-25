package service

import (
	"golang.org/x/xerrors"

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

func (svc Service) CreateUser(user models.User) (id string, err error) {
	id, err = svc.store.CreateUser(user)
	if err != nil {
		return "", xerrors.Errorf("Couldn't write user to store: %w", err)
	}
	return id, nil
}

func (svc Service) GetUser(id string) (models.User, error) {
	user, err := svc.store.GetUser(id)
	if err != nil {
		if err == store.UserNotFoundError {
			return models.User{}, UserNotFoundError
		}
		return models.User{}, xerrors.Errorf("Couldn't fetch user from store: %w", err)
	}
	return user, nil
}
