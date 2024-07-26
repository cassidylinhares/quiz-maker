package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"quizArena/models"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
)

// Purpose: Make the quiz and load the questions

func NewQuiz() *models.Quiz {
	return &models.Quiz{
		Questions: make([]models.Question, 0),
	}
}

func LoadQuestions(quiz *models.Quiz) (models.Quiz, error) {
	file, err := os.Open("my_quiz.csv")
	if err != nil {
		log.Fatal("Failed to open file")
		return *quiz, err
	}

	r := csv.NewReader(file)

	for {
		question, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
			return *quiz, err
		}

		q := models.Question{
			Question: question[0],
			Answer:   question[1],
			Options:  question[2:],
		}
		quiz.Questions = append(quiz.Questions, q)
	}

	file.Close()

	return *quiz, nil
}

func quizRunner(conn *ws.Conn) {
	// get questions

	questions, err := LoadQuestions()

	if err != nil {
		log.Fatal(err)
	}

	for _, q := range questions {
		err := conn.WriteMessage(2, []byte(q.Question))
		if err != nil {
			log.Fatal(err)
			return
		}

		// start the timer

		// wait for answers

		// if no answer then move to next question
		// if answer then stop timer and move to next question
	}

	for {
		msgType, p, err := conn.ReadMessage()

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(msgType, p)
	}
}

func createQuiz(c *gin.Context) {
	var form formT

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	file, err := os.Create("my_quiz.csv")
	if err != nil {
		log.Fatal("Failed to create file")
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, question := range form.Questions {
		line := []string{question.Question, question.Answer}
		line = append(line, question.Options...)

		if err := w.Write(line); err != nil {
			log.Fatal("Failed to write to file")
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	c.IndentedJSON(http.StatusOK, file.Name())
}
