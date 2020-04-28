package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mackstann/payment_user_svc_exercise/internal/routes"
	"github.com/mackstann/payment_user_svc_exercise/internal/service"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
	"github.com/mackstann/payment_user_svc_exercise/internal/stripe_gateway"
)

func main() {
	addr := "127.0.0.1:8000"

	stripeKey := os.Getenv("STRIPE_KEY")
	if stripeKey == "" {
		log.Fatal("Please set the STRIPE_KEY env var")
	}

	store := store.NewStore()
	stripe := stripe_gateway.NewGateway(stripeKey)
	svc := service.NewService(store, stripe)
	r := routes.NewRouter(svc, addr)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	log.Fatal(srv.ListenAndServe())
}
