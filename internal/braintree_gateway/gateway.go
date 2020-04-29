package braintree_gateway

import (
	"golang.org/x/xerrors"
	"log"

	"github.com/braintree-go/braintree-go"

	"github.com/mackstann/payment_user_svc_exercise/internal/models"
)

type Gateway struct {
	bt *braintree.Braintree
}

func NewGateway(accessToken string) (Gateway, error) {
	bt, err := braintree.NewWithAccessToken(accessToken)
	if err != nil {
		return xerrors.Errorf("Couldn't initialize braintree with access token: %w", err)
	}
	return Gateway{
		bt: bt,
	}
}

func (g Gateway) CreateUser(user models.User) (id string, err error) {
	ctx := context.Background()
	customer, err := bt.Customer().Create(ctx, &braintree.CustomerRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		return "", xerrors.Errorf("Couldn't create braintree user: %w", err)
	}

	address, err := bt.Address().Create(ctx, customer.Id, &braintree.AddressRequest{
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

	return customer.Id, nil

}
