package main

import (
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	flag.Parse()

	questions, answers, err := ReadCsv(*fileName)
	if err != nil {
		fmt.Println("Error Opening Csv file", err)
	}

	var correctAns int
	var incorrectAns int

	for i, question := range questions {
		userAns := GetAns(i+1, question)

		if userAns == answers[i] {
			correctAns++
		} else {
			incorrectAns++
		}
	}

	// Output the results
	fmt.Printf("Number of Correct Answers: %d\n", correctAns)
	fmt.Printf("Number of Incorrect Answers: %d\n", incorrectAns)
}
