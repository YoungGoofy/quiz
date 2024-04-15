package app

import (
	"github.com/YoungGoofy/quiz/internal/db"
	"log"
)

func Run() {
	d, err := db.NewDatabase()
	if err != nil {
		log.Panicln(err)
	}
	studentId, err := d.CreateStudent(3, 20, "Frosty", "Frosty", "test@gmail.com")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(studentId)
}
