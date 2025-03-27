package games

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func LcmGame() bool {
	num1 := rand.Intn(20) + 1
	num2 := rand.Intn(20) + 1
	num3 := rand.Intn(20) + 1

	correctAnswer := lcm(lcm(num1, num2), num3)

	fmt.Printf("Question: %d %d %d\n", num1, num2, num3)
	fmt.Print("Your answer: ")

	var userInput string
	fmt.Scanln(&userInput)

	userAnswer, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Invalid input.")
		os.Exit(1)
	}

	if userAnswer == correctAnswer {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n", userAnswer, correctAnswer)
		return false
	}
}
