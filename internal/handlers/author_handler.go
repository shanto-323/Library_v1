package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

type AuthorHandler struct {
	Service *services.AuthorService
}

func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{Service: service}
}

// CREATE
func (h *AuthorHandler) CreateAuthorHandler(w http.ResponseWriter, r *http.Request) error {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		return err
	}

	if err := h.Service.CreateAuthor(&author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusCreated, author)
	return nil
}

// GETSTUDENTBYID
func (h *AuthorHandler) GetAuthorByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	author, err := h.Service.GetAuthorByID(uint(id))
	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, author)
	return nil
}

// GETALL
func (h *AuthorHandler) GetAllAuthorsHandler(w http.ResponseWriter, r *http.Request) error {
	authors, err := h.Service.GetAllAuthors()
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusOK, authors)
	return nil
}

// UPDATE
func (h *AuthorHandler) UpdateAuthorHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		return err
	}

	author.ID = uint(id)
	if err := h.Service.UpdateAuthor(&author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusOK, author)
	return nil
}

// DELETE
func (h *AuthorHandler) DeleteAuthorHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return nil
	}

	if err := h.Service.DeleteAuthor(uint(id)); err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, "user deleted")
	return nil
}
