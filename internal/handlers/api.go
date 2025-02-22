package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/shanto-323/Library_v1.git/internal/repositories"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

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
	studentSubRouter := router.PathPrefix("/user").Subrouter()
	studentRouter(student_handler, studentSubRouter)

	bookRepo := repository.NewBookRepo(a.db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := NewBookHandler(bookService)
	bookSubRouter := router.PathPrefix("/book").Subrouter()
	bookRouter(bookHandler, bookSubRouter)

	authorRepo := repository.NewAuthorRepo(a.db)
	authorService := services.NewAuthorService(authorRepo)
	authorHandler := NewAuthorHandler(authorService)
	authorSubRouter := router.PathPrefix("/author").Subrouter()
	authorRouter(authorHandler, authorSubRouter)

	genreRepo := repository.NewGenreRepo(a.db)
	genreService := services.NewGenreService(genreRepo)
	genreHandler := NewGenreHandler(genreService)
	genreSubRouter := router.PathPrefix("/genre").Subrouter()
	genreRouter(genreHandler, genreSubRouter)

	borrowedBooksRepo := repository.NewBorrowedBookRepository(a.db)
	borrowedBooksService := services.NewBorrowedBookService(borrowedBooksRepo)
	borrowedBooksHandler := NewBorrowedBookHandler(borrowedBooksService)
	borrowedBooksSubRouter := router.PathPrefix("/borrowed-books").Subrouter()
	borrowedBooksRouter(borrowedBooksHandler, borrowedBooksSubRouter)

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

func studentRouter(student_handler *StudentHandler, subRouter *mux.Router) {
	subRouter.HandleFunc("/all", makeHttpHandleFunc(student_handler.GetAllStudentsHandler)).Methods("GET")
	subRouter.HandleFunc("/new", makeHttpHandleFunc(student_handler.CreateStudentHandler)).Methods("POST")
	subRouter.HandleFunc("/get/{id}", makeHttpHandleFunc(student_handler.GetStudentByIDHandler)).Methods("GET")
	subRouter.HandleFunc("/delete/{id}", makeHttpHandleFunc(student_handler.DeleteStudentHandler)).Methods("DELETE")
	subRouter.HandleFunc("/update/{id}", makeHttpHandleFunc(student_handler.UpdateStudentHandler)).Methods("PATCH")
}

func bookRouter(book_handler *BookHandler, subRouter *mux.Router) {
	subRouter.HandleFunc("/all", makeHttpHandleFunc(book_handler.GetAllBooksHandler)).Methods("GET")
	subRouter.HandleFunc("/new", makeHttpHandleFunc(book_handler.CreateBookHandler)).Methods("POST")
	subRouter.HandleFunc("/get/{id}", makeHttpHandleFunc(book_handler.GetBookByIDHandler)).Methods("GET")
	subRouter.HandleFunc("/delete/{id}", makeHttpHandleFunc(book_handler.DeleteBookHandler)).Methods("DELETE")
	subRouter.HandleFunc("/update/{id}", makeHttpHandleFunc(book_handler.UpdateBookHandler)).Methods("PATCH")
}

func authorRouter(author_handler *AuthorHandler, subRouter *mux.Router) {
	subRouter.HandleFunc("/all", makeHttpHandleFunc(author_handler.GetAllAuthorsHandler)).Methods("GET")
	subRouter.HandleFunc("/new", makeHttpHandleFunc(author_handler.CreateAuthorHandler)).Methods("POST")
	subRouter.HandleFunc("/get/{id}", makeHttpHandleFunc(author_handler.GetAuthorByIDHandler)).Methods("GET")
	subRouter.HandleFunc("/delete/{id}", makeHttpHandleFunc(author_handler.DeleteAuthorHandler)).Methods("DELETE")
	subRouter.HandleFunc("/update/{id}", makeHttpHandleFunc(author_handler.UpdateAuthorHandler)).Methods("PATCH")
}

func genreRouter(genre_handler *GenreHandler, subRouter *mux.Router) {
	subRouter.HandleFunc("/all", makeHttpHandleFunc(genre_handler.GetAllGenresHandler)).Methods("GET")
	subRouter.HandleFunc("/new", makeHttpHandleFunc(genre_handler.CreateGenreHandler)).Methods("POST")
	subRouter.HandleFunc("/delete/{id}", makeHttpHandleFunc(genre_handler.DeleteGenreHandler)).Methods("DELETE")
	subRouter.HandleFunc("/update/{id}", makeHttpHandleFunc(genre_handler.UpdateGenreHandler)).Methods("PATCH")
}

func borrowedBooksRouter(borrowedBooks_handler *BorrowedBookHandler, subRouter *mux.Router) {
	subRouter.HandleFunc("/all", makeHttpHandleFunc(borrowedBooks_handler.GetAllBorrowedBooksHandler)).Methods("GET")
	subRouter.HandleFunc("/new", makeHttpHandleFunc(borrowedBooks_handler.CreateBorrowedBookHandler)).Methods("POST")
	subRouter.HandleFunc("/get/{id}", makeHttpHandleFunc(borrowedBooks_handler.GetBorrowedBookByIDHandler)).Methods("GET")
	subRouter.HandleFunc("/delete/{id}", makeHttpHandleFunc(borrowedBooks_handler.DeleteBorrowedBookHandler)).Methods("DELETE")
	subRouter.HandleFunc("/update/{id}", makeHttpHandleFunc(borrowedBooks_handler.UpdateBorrowedBookHandler)).Methods("PATCH")
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
