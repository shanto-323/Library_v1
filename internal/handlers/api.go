package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	store "github.com/shanto-323/Library_v1.git/internal/storage"
)

type LibraryApi struct {
	ipAddr string
	user   store.UserStorage
	//Add Books
}

func NewLibraryApi(ipAddr string, user store.UserStorage) *LibraryApi {
	return &LibraryApi{
		ipAddr: ipAddr,
		user:   user,
	}
}

func (a *LibraryApi) Start() {
	router := mux.NewRouter()

	userHandler := NewUserhandler(a.user)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, err.Error())
		}
	}
}
