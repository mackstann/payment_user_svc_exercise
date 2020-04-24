package main

import (
	"log"
	"net/http"

	"github.com/mackstann/payment_user_svc_exercise/internal/routes"
	"github.com/mackstann/payment_user_svc_exercise/internal/service"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
)

func main() {
	store := store.NewStore()
	svc := service.NewService(store)
	r := routes.NewRouter(svc)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}
