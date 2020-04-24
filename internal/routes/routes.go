package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mackstann/payment_user_svc_exercise/internal/models"
	"github.com/mackstann/payment_user_svc_exercise/internal/service"
)

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request)

func PostUserHandler(svc service.Service) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validation

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			// TODO
		}

		id, err := svc.CreateUser(user)
		if err != nil {
			// TODO
		}

		// fetch a new copy because the service may have applied default values, business logic, etc. to
		// transform the input into a legitimate user entity.
		user, err = svc.GetUser(id)
		if err != nil {
			// TODO
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			// TODO
		}

		w.Write(userJSON)
	}
}

func GetUserHandler(svc service.Service) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// TODO: validation

		user, err := svc.GetUser(id)
		if err != nil {
			// TODO
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			// TODO
		}

		w.Write(userJSON)
	}
}

const pathPrefix = "/v1"

func NewRouter(svc service.Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(pathPrefix+"/user", PostUserHandler(svc))
	r.HandleFunc(pathPrefix+"/user/{id}", GetUserHandler(svc))
	return r
}
