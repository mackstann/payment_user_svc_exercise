package braintree_gateway

import (
	"context"
	"golang.org/x/xerrors"
	"log"

	"github.com/braintree-go/braintree-go"

	"github.com/mackstann/payment_user_svc_exercise/internal/models"
)

type Gateway struct {
	bt *braintree.Braintree
}

func NewGateway(env braintree.Environment, merchantID string, publicKey string, privateKey string) Gateway {
	bt := braintree.New(env, merchantID, publicKey, privateKey)
	return Gateway{bt: bt}
}

func (g Gateway) CreateUser(user models.User) (id string, err error) {
	ctx := context.Background()
	customer, err := g.bt.Customer().Create(ctx, &braintree.CustomerRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		return "", xerrors.Errorf("Couldn't create braintree user: %w", err)
	}

	_, err = g.bt.Address().Create(ctx, customer.Id, &braintree.AddressRequest{
		StreetAddress:     user.Address.Line1,
		ExtendedAddress:   user.Address.Line2,
		Locality:          user.Address.City,
		Region:            user.Address.Subdivision,
		PostalCode:        user.Address.PostalCode,
		CountryCodeAlpha2: "US", // TODO: Internationalization :-)

	})
	if err != nil {
		return "", xerrors.Errorf("Couldn't create braintree address: %w", err)
	}

	log.Printf("Created Braintree user %s", customer.Id)
	return customer.Id, nil
}
