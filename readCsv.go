package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(fileName string) ([]string, []string, error) {
	file, err := os.Open(fileName)
	if err != nil {
	return nil, nil, err
	}

	read := csv.NewReader(file)
	quizes, err := read.ReadAll()
	
	if err !=  nil {
		exit("Failed to parse the provided CSV file")
	}


	var questions []string
	var answers []string

	for _, record := range quizes {
		questions = append(questions, record[0])
		answers = append(answers, record[1])
	}

	return questions, answers, nil
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}