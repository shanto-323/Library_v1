package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/shanto-323/Library_v1.git/internal/repository"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

type LibraryApi struct {
	ipAddr string
	db     *gorm.DB
}

func NewLibraryApi(ipAddr string, db *gorm.DB) *LibraryApi {
	return &LibraryApi{
		ipAddr: ipAddr,
		db:     db,
	}
}

func (a *LibraryApi) Start() {
	router := mux.NewRouter()

	student_repo := repository.NewStudentRepo(a.db)
	student_services := services.NewStudentService(*student_repo)
	student_handler := NewStudentHandler(student_services)

	router.HandleFunc("/user", makeHttpHandleFunc(student_handler.GetAllStudents))
	router.HandleFunc("/user/new", makeHttpHandleFunc(student_handler.CreateStudent))

	fmt.Println(`
      _        _____   ____    _____               _____   __     __                       __ 
     | |      |_   _| |  _ \  |  __ \      /\     |  __ \  \ \   / /                      /_ |
     | |        | |   | |_) | | |__) |    /  \    | |__) |  \ \_/ /     ______    __   __  | |
     | |        | |   |  _ <  |  _  /    / /\ \   |  _  /    \   /     |______|   \ \ / /  | |
     | |____   _| |_  | |_) | | | \ \   / ____ \  | | \ \     | |                  \ V /   | |
     |______| |_____| |____/  |_|  \_\ /_/    \_\ |_|  \_\    |_|                   \_/    |_|
   `)
	if err := http.ListenAndServe(a.ipAddr, router); err != nil {
		log.Fatal(err)
	}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, err.Error())
		}
	}
}

func getId(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id '%v' given", id)
	}
	return id, nil
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
