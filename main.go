package main

import (
	"log"
	"net/http"

	"github.com/mackstann/payment_user_svc_exercise/internal/routes"
	"github.com/mackstann/payment_user_svc_exercise/internal/service"
	"github.com/mackstann/payment_user_svc_exercise/internal/store"
)

func main() {
	addr := "127.0.0.1:8000"

	store := store.NewStore()
	svc := service.NewService(store)
	r := routes.NewRouter(svc, addr)

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	log.Fatal(srv.ListenAndServe())
}
