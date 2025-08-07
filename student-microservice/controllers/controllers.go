package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gajare/swager_golang/models"
	service "github.com/gajare/swager_golang/services"

	"github.com/gorilla/mux"
)

// CreateStudent godoc
// @Summary      Create Student
// @Description  Create new student
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student body models.Student true "Student Data"
// @Success      200  {object}  models.Student
// @Router       /students [post]
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)
	created, err := service.CreateStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

// GetStudents godoc
// @Summary      Get Students
// @Description  Get all students
// @Tags         students
// @Produce      json
// @Success      200  {array}  models.Student
// @Router       /students [get]
func GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := service.GetStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// GetStudent godoc
// @Summary      Get Student by ID
// @Description  Get student details by ID
// @Tags         students
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      200  {object}  models.Student
// @Router       /students/{id} [get]
func GetStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	student, err := service.GetStudentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

// UpdateStudent godoc
// @Summary      Update Student
// @Description  Update student by ID
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID"
// @Param        student body models.Student true "Updated Data"
// @Success      200  {object}  models.Student
// @Router       /students/{id} [put]
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updated models.Student
	json.NewDecoder(r.Body).Decode(&updated)
	student, err := service.UpdateStudent(id, updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(student)
}

// DeleteStudent godoc
// @Summary      Delete Student
// @Description  Delete student by ID
// @Tags         students
// @Produce      json
// @Param        id path int true "Student ID"
// @Success      204  "No Content"
// @Router       /students/{id} [delete]
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := service.DeleteStudent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
