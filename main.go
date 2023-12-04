package main

import (
	"flag"
	"fmt"
)

func main() {
	fileNmae := flag.String("Quiz", "Quiz Question", "Quiz Answers")
	flag.Parse()

	questions, answers, err := ReadCsv(*fileNmae)
	if err != nil {
		fmt.Println("Error Openning Csv file", err)
	}

	var correctAns int
	var incorrectAns int

	for i, question := range questions{
		userAns := GetAns(i+1, question)

		if userAns == answers[1]{
			correctAns += 1
		} else{
			incorrectAns += 1
		}
	}
}