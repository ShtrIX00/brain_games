package games

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func ProgressionGame() bool {
	start := rand.Intn(5) + 1
	ratio := rand.Intn(5) + 2
	length := rand.Intn(6) + 5 // Длина от 5 до 10

	progression := make([]int, length)
	for i := 0; i < length; i++ {
		progression[i] = start * pow(ratio, i)
	}

	hiddenIndex := rand.Intn(length)
	correctAnswer := progression[hiddenIndex]
	progression[hiddenIndex] = -1

	fmt.Print("Question: ")
	for _, num := range progression {
		if num == -1 {
			fmt.Print(".. ")
		} else {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()

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

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
