package service

import (
	"golang.org/x/xerrors"

	"github.com/mackstann/payment_user_svc_exercise/internal/models"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
	"github.com/mackstann/payment_user_svc_exercise/internal/stripe_gateway"
)

type Service struct {
	store  store.Store
	stripe stripe_gateway.Gateway
}

func NewService(store store.Store, stripe stripe_gateway.Gateway) Service {
	return Service{
		store:  store,
		stripe: stripe,
	}
}

func (svc Service) CreateUser(user models.User) (id string, err error) {
	id, err = svc.store.CreateUser(user)
	if err != nil {
		return "", xerrors.Errorf("Couldn't write user to store: %w", err)
	}
	stripeID, err := svc.stripe.CreateUser(user)
	if err != nil {
		return "", xerrors.Errorf("Stripe gateway failed to create user: %w", err)
	}
	svc.store.LinkUserToStripeCustomer(id, stripeID)
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
