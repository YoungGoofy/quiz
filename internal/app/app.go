package app

import (
	"github.com/YoungGoofy/quiz/internal/db"
	"github.com/YoungGoofy/quiz/internal/transport"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("web/views/*")
	router.Static("/static", "./web/static")
	d, err := db.NewDatabase()
	if err != nil {
		log.Panicln(err)
	}
	transport.MainHandler(d, router)
	log.Println("Listening server on: http://localhost:5000")
	if err = router.Run(":5000"); err != nil {
		log.Println(err)
	}
}
