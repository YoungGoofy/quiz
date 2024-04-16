package app

import (
	"github.com/YoungGoofy/quiz/internal/db"
	"github.com/YoungGoofy/quiz/internal/transport"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("internal/web/views/*")
	d, err := db.NewDatabase()
	if err != nil {
		log.Panicln(err)
	}
	transport.MainHandler(d, router)
	log.Println("Listening server on: http://localhost:3000")
	if err = router.Run(":3000"); err != nil {
		log.Println(err)
	}
}
