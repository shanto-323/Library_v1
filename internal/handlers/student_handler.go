package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shanto-323/Library_v1.git/internal/models"
	"github.com/shanto-323/Library_v1.git/internal/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{Service: service}
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) error {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		return err
	}

	if err := h.Service.CreateStudent(&student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusCreated, student)
	return nil
}

func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	student, err := h.Service.GetStudentById(uint(id))
	if err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, student)
	return nil
}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) error {
	students, err := h.Service.GetAllStudent()
	if err != nil {
		return err
	}
	WriteJson(w, http.StatusOK, students)
	return nil
}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		return err
	}

	student.ID = uint(id)
	if err := h.Service.UpdateStudent(&student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	WriteJson(w, http.StatusOK, student)
	return nil
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return nil
	}

	if err := h.Service.DeleteStudent(uint(id)); err != nil {
		return err
	}

	WriteJson(w, http.StatusOK, "user deleted")
	return nil
}
