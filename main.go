package main

import (
	// "flag"

	"flag"
	"log"
	"net/http"

	// "net/http"
	// "quizArena/client"
	"quizArena/services"
	// "github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8080", "http server address")
var tables = [...]string{"bird", "cake"}

func main() {
	flag.Parse()

	server := services.NewServer(tables[:])
	go server.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		services.ServeWs(server, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))

}

// type Question struct {
// 	Question string   `json:"question"`
// 	Answer   string   `json:"answer"`
// 	Options  []string `json:"options"`
// }

// type Quiz struct {
// 	Questions []Question `json:"questions"`
// }

// func NewQuiz() *Quiz {
// 	return &Quiz{
// 		Questions: make([]Question, 0),
// 	}
// }

// func (quiz *Quiz) LoadQuestions() (Quiz, error) {
// 	file, err := os.Open("my_quiz.csv")
// 	if err != nil {
// 		log.Fatal("Failed to open file")
// 		return *quiz, err
// 	}

// 	r := csv.NewReader(file)

// 	for {
// 		question, err := r.Read()

// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			log.Fatal(err)
// 			return *quiz, err
// 		}

// 		q := Question{
// 			Question: question[0],
// 			Answer:   question[1],
// 			Options:  question[2:],
// 		}
// 		quiz.Questions = append(quiz.Questions, q)
// 	}

// 	file.Close()

// 	return *quiz, nil
// }

// func countTimer(t *time.Timer, ch *chan int) {

// }

// func verifyAnswer(answer string, questionAnswer string) bool {
// 	return strings.Compare(strings.ToLower(answer), strings.ToLower(questionAnswer)) == 0
// }

// func main() {
// q := NewQuiz()
// quiz, err := q.LoadQuestions()

// if err != nil {
// 	log.Fatal(err)
// }

// Timer
// timer := time.NewTimer(5 * time.Second)
// ch := make(chan int)

// for _, question := range quiz.Questions {
// 	fmt.Println(question.Question)
// 	// go countTimer(timer, &ch)

// 	isCorrect := verifyAnswer("22", question.Answer)
// 	fmt.Println(isCorrect)
// }
// }

// var addr = flag.String("addr", ":8080", "http server address")
// svr := server.NewServer()
// go svr.Run()

// flag.Parse()

// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 	c, err := client.ServeWs(w, r)

// 	if err != nil {
// 		return
// 	}

// 	svr.Register <- c
// })

// log.Fatal(http.ListenAndServe(*addr, nil))
