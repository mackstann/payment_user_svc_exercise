package routes

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mackstann/payment_user_svc_exercise/internal/models"
	"github.com/mackstann/payment_user_svc_exercise/internal/service"
)

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request)

func handleUserPOST(svc service.Service, addr string) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validation

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			handleInternalError(err, w, r)
		}

		id, err := svc.CreateUser(user)
		if err != nil {
			handleCreateUserError(err, w, r)
		}

		// fetch a new copy because the service may have applied default values, business logic, etc. to
		// transform the input into a legitimate user entity.
		user, err = svc.GetUser(id)
		if err != nil {
			handleGetUserError(err, w, r)
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			handleInternalError(err, w, r)
		}

		w.Header()["Location"] = []string{buildUserURI(addr, id)}
		w.WriteHeader(http.StatusCreated)
		w.Write(userJSON)
	}
}

func buildUserURI(addr string, userID string) string {
	// this is inelegant.
	return "http://" + addr + pathPrefix + "/user/" + userID
}

func handleUserGET(svc service.Service, addr string) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// TODO: validation

		user, err := svc.GetUser(id)
		if err != nil {
			handleGetUserError(err, w, r)
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			handleInternalError(err, w, r)
		}

		w.Write(userJSON)
	}
}

type HTTPError struct {
	status  int
	message string
}

func handleHTTPError(httpError HTTPError, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(httpError.status)
	w.Write([]byte(httpError.message))
}

func handleGetUserError(err error, w http.ResponseWriter, r *http.Request) {
	var httpError HTTPError
	if xerrors.Is(err, service.UserNotFoundError) {
		httpError.status = http.StatusNotFound
		httpError.message = "User not found"
	} else {
		handleInternalError(err, w, r)
	}
	handleHTTPError(httpError, w, r)
}

func handleCreateUserError(err error, w http.ResponseWriter, r *http.Request) {
	// no known errors to handle yet
	handleInternalError(err, w, r)
}

func handleInternalError(err error, w http.ResponseWriter, r *http.Request) {
	httpError := HTTPError{
		status:  http.StatusInternalServerError,
		message: "Internal Error",
	}
	// would be nice to log more: request ID/details, stack trace, etc.
	log.Printf("Unhandled error: %s", err)
	handleHTTPError(httpError, w, r)
}

const pathPrefix = "/v1"

func NewRouter(svc service.Service, addr string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(pathPrefix+"/user", handleUserPOST(svc, addr))
	r.HandleFunc(pathPrefix+"/user/{id}", handleUserGET(svc, addr))
	return r
}
