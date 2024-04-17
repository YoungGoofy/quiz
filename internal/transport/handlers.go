package transport

import (
	"github.com/YoungGoofy/quiz/internal/db"
	"github.com/YoungGoofy/quiz/internal/transport/questions"
	"github.com/YoungGoofy/quiz/internal/transport/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainHandler(d *db.DB, g *gin.Engine) {
	u := user.NewUser(d)
	g.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home",
		})
	})
	g.GET("/create-user", u.OpenAuthPage)
	g.POST("/create-user", u.CreateUser)
	g.GET("/quiz/:grade_id/:id", questions.NewQuestion)
	g.POST("/quiz/:grade_id/:id", questions.NextQuestion)
	g.GET("/result", Result)
}

func Result(c *gin.Context) {
	c.HTML(http.StatusOK, "result.html", gin.H{
		"title": "Result",
	})
}
