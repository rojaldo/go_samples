package trivia

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sample/apod"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

type OpenTDB struct {
	Category      string
	Type          string
	Difficulty    string
	Question      string
	CorrectAnswer string
	Answers       []string
}

func (*OpenTDB) GetOpenTDB() OpenTDB {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://opentdb.com/api.php?amount=1")
	if err != nil {
		log.Fatal(err)
		return OpenTDB{}
	}
	json := gjson.Get(string(resp.Body()), "results.0")

	var tempAnswers []string
	for _, answer := range json.Get("incorrect_answers").Array() {
		tempAnswers = append(tempAnswers, answer.String())
	}
	tempAnswers = append(tempAnswers, json.Get("correct_answer").String())

	// shuffle answers
	for i := range tempAnswers {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		j := r1.Intn(i + 1)
		tempAnswers[i], tempAnswers[j] = tempAnswers[j], tempAnswers[i]
	}
	return OpenTDB{
		Category:      json.Get("category").String(),
		Type:          json.Get("type").String(),
		Difficulty:    json.Get("difficulty").String(),
		Question:      json.Get("question").String(),
		CorrectAnswer: json.Get("correct_answer").String(),
		Answers:       tempAnswers,
	}
}

func CreateApp() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "date",
			Value: "today",
			Usage: "date to get data for (format: YYYY-MM-DD) (default: today)",
		},
		cli.StringFlag{
			Name:  "trivia",
			Value: "today",
			Usage: "Answer a question about anything",
		},
	}
	app.Action = func(c *cli.Context) error {
		if apod.IsValidDate(c.String("apod")) {
			fmt.Println("Valid date: ", c.String("apod"))
			apod := apod.Apod{}
			apod = apod.GetApod(c.String("date"))
			fmt.Println(apod.Title)
			fmt.Println(apod.Date)
			fmt.Println(apod.Explanation)
		} else if c.String("trivia") != "" { // trivia parameter is a number
			fmt.Println("=========== Today's Trivia ===========")
			trivia := OpenTDB{}
			trivia = trivia.GetOpenTDB()
			fmt.Println(trivia.Question)
			for index, answer := range trivia.Answers {
				fmt.Printf("%d.- "+answer, index+1)
				fmt.Println()
			}
			answerIndex := 0
			fmt.Scanln(&answerIndex)
			if trivia.Answers[answerIndex-1] == trivia.CorrectAnswer {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect!")
			}
		} else {
			fmt.Println("Invalid date: ", c.String("date"))
		}
		return nil
	}

	app.Run(os.Args)
}
