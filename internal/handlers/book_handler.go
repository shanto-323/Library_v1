package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

type BookHandler struct {
	Service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// CREATE
func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) error {
	var genre models.Book
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		return err
	}

	if err := h.Service.CreateBook(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusCreated, genre)
	return nil
}

// GETSTUDENTBYID
func (h *BookHandler) GetBookByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	genre, err := h.Service.GetBookByID(uint(id))
	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, genre)
	return nil
}

// GETALL
func (h *BookHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) error {
	books, err := h.Service.GetAllBooks()
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusOK, books)
	return nil
}

// UPDATE
func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	var genre models.Book
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		return err
	}

	genre.ID = uint(id)
	if err := h.Service.UpdateBook(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusOK, genre)
	return nil
}

// DELETE
func (h *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return nil
	}

	if err := h.Service.DeleteBook(uint(id)); err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, "user deleted")
	return nil
}
