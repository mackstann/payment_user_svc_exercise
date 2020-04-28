package stripe_gateway

import (
	"golang.org/x/xerrors"
	"log"

	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/customer"

	"github.com/mackstann/payment_user_svc_exercise/internal/models"
)

type Gateway struct {
}

func NewGateway(key string) Gateway {
	// normally i'd prefer to use a client object and not rely on mutable global state.
	stripe.Key = key

	return Gateway{}
}

func (g Gateway) CreateUser(user models.User) (id string, err error) {
	params := &stripe.CustomerParams{
		Name: stripe.String(user.FirstName + " " + user.MiddleName + " " + user.LastName),
		Address: &stripe.AddressParams{
			Line1:      stripe.String(user.Address.Line1),
			Line2:      stripe.String(user.Address.Line2),
			City:       stripe.String(user.Address.City),
			State:      stripe.String(user.Address.Subdivision),
			PostalCode: stripe.String(user.Address.PostalCode),
			Country:    stripe.String("US"), // TODO: Internationalization :-)
		},
	}
	c, err := customer.New(params)
	if err != nil {
		return "", xerrors.Errorf("Couldn't create stripe user: %w", err)
	}

	log.Printf("Created Stripe user %s", c.ID)
	return c.ID, nil
}

func (g Gateway) GetUser(id string) (stripe.Customer, error) {
	return stripe.Customer{}, nil
}
