package main

import (
	"log"
	"net/http"
	"os"

	"github.com/braintree-go/braintree-go"

	"github.com/mackstann/payment_user_svc_exercise/internal/braintree_gateway"
	"github.com/mackstann/payment_user_svc_exercise/internal/routes"
	"github.com/mackstann/payment_user_svc_exercise/internal/service"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
	"github.com/mackstann/payment_user_svc_exercise/internal/stripe_gateway"
)

const addr = "127.0.0.1:8000"

var braintreeEnv = braintree.Sandbox

func main() {
	stripeKey := os.Getenv("STRIPE_KEY")
	if stripeKey == "" {
		log.Fatal("Please set the STRIPE_KEY env var")
	}
	stripe := stripe_gateway.NewGateway(stripeKey)

	braintreeMerchantID := os.Getenv("BRAINTREE_MERCHANT_ID")
	braintreePublicKey := os.Getenv("BRAINTREE_PUBLIC_KEY")
	braintreePrivateKey := os.Getenv("BRAINTREE_PRIVATE_KEY")
	if braintreeMerchantID == "" {
		log.Fatal("Please set the BRAINTREE_MERCHANT_ID env var")
	}
	if braintreePublicKey == "" {
		log.Fatal("Please set the BRAINTREE_PUBLIC_KEY env var")
	}
	if braintreePrivateKey == "" {
		log.Fatal("Please set the BRAINTREE_PRIVATE_KEY env var")
	}
	braintree := braintree_gateway.NewGateway(
		braintreeEnv, braintreeMerchantID, braintreePublicKey, braintreePrivateKey)

	store := store.NewStore()
	svc := service.NewService(store, stripe, braintree)
	r := routes.NewRouter(svc, addr)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	log.Fatal(srv.ListenAndServe())
}
