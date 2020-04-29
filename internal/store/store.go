package store

import (
	"github.com/google/uuid"
	"github.com/mackstann/payment_user_svc_exercise/internal/models"
)

type Store struct {
	users        map[string]models.User
	stripeIDs    map[string]string
	braintreeIDs map[string]string
}

func NewStore() Store {
	return Store{
		users:        make(map[string]models.User),
		stripeIDs:    make(map[string]string),
		braintreeIDs: make(map[string]string),
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

func (store Store) LinkUserToStripeCustomer(userID string, stripeCustomerID string) error {
	store.stripeIDs[userID] = stripeCustomerID
	return nil
}

func (store Store) LinkUserToBraintreeCustomer(userID string, braintreeCustomerID string) error {
	store.braintreeIDs[userID] = braintreeCustomerID
	return nil
}

func (store Store) GetUser(id string) (models.User, error) {
	user, ok := store.users[id]
	if !ok {
		return models.User{}, UserNotFoundError
	}
	return user, nil
}

func (store Store) GetStripeCustomerID(userID string) (*string, error) {
	id, ok := store.stripeIDs[userID]
	if !ok {
		return nil, nil
	}
	return &id, nil
}

func (store Store) GetBraintreeCustomerID(userID string) (*string, error) {
	id, ok := store.braintreeIDs[userID]
	if !ok {
		return nil, nil
	}
	return &id, nil
}
