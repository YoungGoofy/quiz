package db

import (
	"github.com/YoungGoofy/quiz/internal/db/connections"
	"github.com/YoungGoofy/quiz/internal/db/handlers"
	"github.com/YoungGoofy/quiz/internal/models"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDatabase() (*DB, error) {
	db, err := connections.NewDatabaseConnection()
	return &DB{db}, err
}

func (d *DB) GetGrades() ([]models.Grade, error) {
	return handlers.GetGrades(d.db)
}

func (d *DB) GetGradeById(id int) (models.Grade, error) {
	return handlers.GetGradeById(id, d.db)
}

func (d *DB) CreateStudent(gradeId int, age int, name, lastname, mail string) (uint, error) {
	return handlers.CreateStudent(gradeId, age, name, lastname, mail, d.db)
}
