package questions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Questions struct {
	Text    string
	Options []string
}

var questions = map[int][]Questions{
	1: {
		{"Вопрос 1", []string{"Ответ 1.1", "Ответ 1.2", "Ответ 1.3"}},
		{"Вопрос 2", []string{"Ответ 2.1", "Ответ 2.2", "Ответ 2.3"}},
		{"Вопрос 3", []string{"Ответ 3.1", "Ответ 3.2", "Ответ 3.3"}},
	},
	2: {
		{"Вопрос 1", []string{"Ответ 1.1", "Ответ 1.2", "Ответ 1.3"}},
		{"Вопрос 2", []string{"Ответ 2.1", "Ответ 2.2", "Ответ 2.3"}},
		{"Вопрос 3", []string{"Ответ 3.1", "Ответ 3.2", "Ответ 3.3"}},
	},
	3: {
		{"Вопрос 1", []string{"Ответ 1.1", "Ответ 1.2", "Ответ 1.3"}},
		{"Вопрос 2", []string{"Ответ 2.1", "Ответ 2.2", "Ответ 2.3"}},
		{"Вопрос 3", []string{"Ответ 3.1", "Ответ 3.2", "Ответ 3.3"}},
	},
	4: {
		{"Вопрос 1", []string{"Ответ 1.1", "Ответ 1.2", "Ответ 1.3"}},
		{"Вопрос 2", []string{"Ответ 2.1", "Ответ 2.2", "Ответ 2.3"}},
		{"Вопрос 3", []string{"Ответ 3.1", "Ответ 3.2", "Ответ 3.3"}},
	},
	5: {
		{"Вопрос 1", []string{"Ответ 1.1", "Ответ 1.2", "Ответ 1.3"}},
		{"Вопрос 2", []string{"Ответ 2.1", "Ответ 2.2", "Ответ 2.3"}},
		{"Вопрос 3", []string{"Ответ 3.1", "Ответ 3.2", "Ответ 3.3"}},
	},
}

func NewQuestion(c *gin.Context) {
	gradeId, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	questionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if questionId < 0 || questionId >= len(questions) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Question not found",
		})
		return
	}
	c.HTML(http.StatusOK, "question.html", gin.H{
		"title":    fmt.Sprintf("Question %d", questionId),
		"question": questions[gradeId][questionId],
	})
}

func NextQuestion(c *gin.Context) {
	gradeId, err := strconv.Atoi(c.Param("grade_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	nextQuestionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Printf("\n%d\t%d\n", nextQuestionID, len(questions[gradeId])-1)
	if nextQuestionID >= len(questions[gradeId])-1 {
		c.Redirect(http.StatusSeeOther, "/result")
		return
	}
	nextQuestionID++
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/quiz/%d/%d", gradeId, nextQuestionID))
}
