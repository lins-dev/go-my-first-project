package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("Try to guess the chosen number.")
	fmt.Println("You only have 10 tries.")
	if x := 10; x > 5 {

	}
	var userChoices []int
	var choiceNumber int = rand.Intn(20)
	var isCorrect bool = false
	for i := 10; i > 0; i-- {
		var userChoice int

		fmt.Printf("You have %d tries \n", i)
		fmt.Println("Enter your number:")
		fmt.Scan(&userChoice)

		userChoices = append(userChoices, userChoice)

		if choiceNumber == userChoice {
			isCorrect = true
			break
		}
	}
	if isCorrect {
		fmt.Println("CONGRATS! You got it right!")
	} else {
		fmt.Println("CONGRATS! You are a LOSER!")
	}

	fmt.Println("Yours choice are:")
	fmt.Println(userChoices)
	fmt.Printf("The Correct choice is: %d \n", choiceNumber)

}
