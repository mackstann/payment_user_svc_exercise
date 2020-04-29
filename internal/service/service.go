package service

import (
	"golang.org/x/xerrors"

	"github.com/mackstann/payment_user_svc_exercise/internal/braintree_gateway"
	"github.com/mackstann/payment_user_svc_exercise/internal/models"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
	"github.com/mackstann/payment_user_svc_exercise/internal/stripe_gateway"
)

type Service struct {
	store     store.Store
	stripe    stripe_gateway.Gateway
	braintree braintree_gateway.Gateway
}

func NewService(store store.Store, stripe stripe_gateway.Gateway,
	braintree braintree_gateway.Gateway) Service {

	return Service{
		store:     store,
		stripe:    stripe,
		braintree: braintree,
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
	braintreeID, err := svc.braintree.CreateUser(user)
	if err != nil {
		return "", xerrors.Errorf("Braintree gateway failed to create user: %w", err)
	}
	svc.store.LinkUserToStripeCustomer(id, stripeID)
	svc.store.LinkUserToBraintreeCustomer(id, braintreeID)
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

func (svc Service) GetUserWithGatewayAccounts(id string) (models.UserWithGatewayAccounts, error) {
	user, err := svc.GetUser(id)
	if err != nil {
		return models.UserWithGatewayAccounts{}, err
	}

	fullUser := models.UserWithGatewayAccounts{User: user}
	fullUser.StripeCustomerID, err = svc.store.GetStripeCustomerID(id)
	if err != nil {
		return models.UserWithGatewayAccounts{}, xerrors.Errorf("Couldn't fetch user's Stripe ID: %w", err)
	}

	fullUser.BraintreeCustomerID, err = svc.store.GetBraintreeCustomerID(id)
	if err != nil {
		return models.UserWithGatewayAccounts{}, xerrors.Errorf("Couldn't fetch user's Braintree ID: %w", err)
	}

	return fullUser, nil
}
