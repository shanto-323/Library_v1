package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

type GenreHandler struct {
	Service *services.GenreService
}

func NewGenreHandler(service *services.GenreService) *GenreHandler {
	return &GenreHandler{Service: service}
}

// CREATE
func (h *GenreHandler) CreateGenreHandler(w http.ResponseWriter, r *http.Request) error {
	var genre models.Genre
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		return err
	}

	if err := h.Service.CreateGenre(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusCreated, genre)
	return nil
}

// GETALL
func (h *GenreHandler) GetAllGenresHandler(w http.ResponseWriter, r *http.Request) error {
	books, err := h.Service.GetAllGenres()
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusOK, books)
	return nil
}

// UPDATE
func (h *GenreHandler) UpdateGenreHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	var genre models.Genre
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		return err
	}

	genre.ID = uint(id)
	if err := h.Service.UpdateGenre(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusOK, genre)
	return nil
}

// DELETE
func (h *GenreHandler) DeleteGenreHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return nil
	}

	if err := h.Service.DeleteGenre(uint(id)); err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, "user deleted")
	return nil
}
