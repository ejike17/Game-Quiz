package main

import (
	"fmt"
	"time"
)

// Package level variable to make it accesible to all functions
var fullScore int = 0

func main() {

	fmt.Println("You are welccome to my game")
	fmt.Println("Please enter your age: ")
	var age int
	fmt.Scan(&age)

	// Ensuring underage users cannot play this game
	if age < 18 {
		fmt.Printf("I am sorry you cannot play this game \n")
	} else {
		fmt.Printf("Checking to see if you qualify to play this game... \n")
		time.Sleep(time.Second * 3)
		fmt.Printf("You qualify \n")
	}

	fmt.Println("Now let us continue with the game")
	fmt.Println("Loading questions please wait.....")

	time.Sleep(time.Second * 5)

	// Initializing slice to store the answer choices of users
	answerChoices := []string{}

	// ##### Questions starts here #################
	fmt.Printf("Question 1: Who will make the best president in 2023? \n")
	fmt.Printf("A. Kogi Bello \nB. Peter Obi\nC. Atiku Abubakar\nD. Bola Ahmed Tinubu \n")
	var answerChoice1 string
	correctAnswer1 := "B"
	fmt.Scan(&answerChoice1)
	//Checking if answer is correct in order to increase total point variable by 1
	checkAnswer(answerChoice1, correctAnswer1)
	answerChoices = append(answerChoices, answerChoice1)

	fmt.Printf("Question 2: Who will make the worst vice-president in 2023? \n")
	fmt.Printf("A. El-Rufai\nB. Kogi Bello\nC. Atiku Abubakar\nD. Osibanjo \n")
	var answerChoice2 string
	correctAnswer2 := "A"
	fmt.Scan(&answerChoice2)
	checkAnswer(answerChoice2, correctAnswer2)
	answerChoices = append(answerChoices, answerChoice2)

	fmt.Printf("Question 3: Whose tenure was better in presidency \n")
	fmt.Printf("A. Olusegun Obasanjo \nB. Muhammadu Buhari\nC. Goodluck Jonathan\n")
	var answerChoice3 string
	correctAnswer3 := "C"
	fmt.Scan(&answerChoice3)
	checkAnswer(answerChoice3, correctAnswer3)
	answerChoices = append(answerChoices, answerChoice3)
	// ##### End of questions #################

	// Calculating percentage score of the user
	var result float64 = (float64(fullScore) / 3.0) * 100.0
	fmt.Printf("You scored %v percent in this quiz\n", result)

	// Array of correct answer option to compare with the user's chosen answer
	correctArray := [3]string{"B", "A", "C"}

	//Calling the checkQuestions function to compare user's answer choices
	for i := 0; i < 3; i++ {
		checkQuestion(answerChoices[i], correctArray[i], i)
	}

}

// Increase the final score of the user based on correct answers
func checkAnswer(answer string, correct string) {
	if answer == correct {
		fullScore = fullScore + 1
	}
}

// At the end of the question, give the user correction on thier answers
func checkQuestion(answer string, correct string, questionNumber int) {
	if answer == correct {
		fmt.Printf("Question %v was answered correctly.\n", questionNumber+1)
		fmt.Printf("The correct answer is %v \n", correct)
	} else {
		fmt.Printf("Question %v was answered incorrectly.\n", questionNumber+1)
		fmt.Printf("The correct answer was %v \n", correct)
	}
}
