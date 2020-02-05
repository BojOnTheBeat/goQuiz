package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getArgs() (int64, string) {

	var filename string
	var timeLimit int64

	if len(os.Args) == 3 {
		filename = os.Args[1]
		timeLimit, _ = strconv.ParseInt(os.Args[2], 10, 64)
	} else if len(os.Args) == 2 {
		fmt.Printf("Using 5 seconds as default")
		filename = os.Args[1]
		timeLimit = 5
	}

	return timeLimit, filename

}

func getQuestions(filename string) [][]string {
	csvFile, _ := os.Open(filename)
	r := csv.NewReader(csvFile)
	records, _ := r.ReadAll()
	return records
}

func getUserAnswer() string {
	userInputReader := bufio.NewReader(os.Stdin)
	userAnswer, _ := userInputReader.ReadString('\n')
	trimUserAnswer := strings.TrimRight(userAnswer, "\n")
	return trimUserAnswer
}

func runQuiz(timer *time.Timer, questions [][]string) {
	var score int

	for _, record := range questions {
		select {
		case <-timer.C:
			fmt.Printf("You're out of time\n")
			fmt.Printf("Your score is: %d out of %d\n", score, len(questions))
			return
		default:
			question := record[0]
			correctSolution := record[1]

			fmt.Print(question + ":")

			userAnswer := getUserAnswer()

			if strings.Compare(userAnswer, correctSolution) == 0 {
				score++
			}

		}

	}
	fmt.Printf("Your score is: %d out of %d\n", score, len(questions))
}

func main() {

	var timeLimit int64
	var fileName string

	timeLimit, fileName = getArgs()

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	records := getQuestions(fileName)
	runQuiz(timer, records)

}
