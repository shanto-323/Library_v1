package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

type BorrowedBookHandler struct {
	Service *services.BorrowedBookService
}

func NewBorrowedBookHandler(service *services.BorrowedBookService) *BorrowedBookHandler {
	return &BorrowedBookHandler{Service: service}
}

// CREATE
func (h *BorrowedBookHandler) CreateBorrowedBookHandler(w http.ResponseWriter, r *http.Request) error {
	var book models.BorrowedBook
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}

	if err := h.Service.BorrowBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusCreated, book)
	return nil
}

// GETSTUDENTBYID
func (h *BorrowedBookHandler) GetBorrowedBookByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	book, err := h.Service.GetBorrowedBookByID(uint(id))
	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, book)
	return nil
}

// GETALL
func (h *BorrowedBookHandler) GetAllBorrowedBooksHandler(w http.ResponseWriter, r *http.Request) error {
	books, err := h.Service.GetAllBorrowedBooks()
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusOK, books)
	return nil
}

// UPDATE
func (h *BorrowedBookHandler) UpdateBorrowedBookHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	var book models.BorrowedBook
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}

	book.ID = uint(id)
	if err := h.Service.UpdateBorrowedBookByID(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusOK, book)
	return nil
}

// DELETE
func (h *BorrowedBookHandler) DeleteBorrowedBookHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return nil
	}

	if err := h.Service.ReturnBook(uint(id)); err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, "user deleted")
	return nil
}
