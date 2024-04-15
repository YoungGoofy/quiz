package handlers

import (
	"github.com/YoungGoofy/quiz/internal/models"
	"gorm.io/gorm"
)

func GetGrades(db *gorm.DB) ([]models.Grade, error) {
	var grades []models.Grade
	result := db.Table("grades").Find(&grades)
	return grades, result.Error
}

func GetGradeById(id int, db *gorm.DB) (models.Grade, error) {
	var grade models.Grade
	result := db.Table("grades").First(&grade, id)
	return grade, result.Error
}

func CreateStudent(gradeId int, age int, name, lastname, mail string, db *gorm.DB) (uint, error) {
	student := models.Student{GradeId: gradeId, Name: name, Lastname: lastname, Mail: mail, Age: age}
	result := db.Select("GradeId", "Name", "Lastname", "Mail", "Age").Create(&student)
	return student.StudentId, result.Error
}