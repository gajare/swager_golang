package service

import (
	"github.com/gajare/swager_golang/db"
	"github.com/gajare/swager_golang/models"
)

func CreateStudent(student models.Student) (models.Student, error) {
	result := db.DB.Create(&student)
	return student, result.Error
}

func GetStudents() ([]models.Student, error) {
	var students []models.Student
	result := db.DB.Find(&students)
	return students, result.Error
}

func GetStudentByID(id int) (models.Student, error) {
	var student models.Student
	result := db.DB.First(&student, id)
	return student, result.Error
}

func UpdateStudent(id int, updated models.Student) (models.Student, error) {
	var student models.Student
	result := db.DB.First(&student, id)
	if result.Error != nil {
		return student, result.Error
	}
	student.Name = updated.Name
	student.Email = updated.Email
	student.Age = updated.Age
	db.DB.Save(&student)
	return student, nil
}

func DeleteStudent(id int) error {
	result := db.DB.Delete(&models.Student{}, id)
	return result.Error
}
