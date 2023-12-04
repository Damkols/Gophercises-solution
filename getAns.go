package main

import "fmt"

func GetAns(questionNum int, question string) string {
	var userAns string
	fmt.Printf("Question %d: %s", questionNum, question)
	fmt.Scanln(&userAns)
	return userAns
}