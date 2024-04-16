package transport

import (
	"github.com/YoungGoofy/quiz/internal/db"
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
	g.GET("/api/create-user", u.OpenAuthPage)
	g.POST("/api/create-user", u.CreateUser)
}
