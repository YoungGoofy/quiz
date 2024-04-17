package user

import (
	"fmt"
	"github.com/YoungGoofy/quiz/internal/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	*db.DB
}

func NewUser(d *db.DB) *User {
	return &User{DB: d}
}

func (u *User) OpenAuthPage(c *gin.Context) {
	grades, err := u.GetGrades()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "auth.html", gin.H{
		"title":  "Auth",
		"grades": grades,
	})
}

func (u *User) CreateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var student struct {
		GradeId int    `form:"grade"`
		Mail    string `form:"email"`
		Age     int    `form:"age"`
	}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := u.CreateStudent(student.GradeId, student.Age, student.Mail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	log.Println("Add new user: ", id)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/quiz/%v/0", student.GradeId))
}
