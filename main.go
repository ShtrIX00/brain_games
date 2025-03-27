package main

import (
	"brain_games/games"
	"fmt"
)

func main() {
	fmt.Println("Choose a game:")
	fmt.Println("1. Least Common Multiple (LCM)")
	fmt.Println("2. Geometric Progression")

	var choice int
	fmt.Print("Enter 1 or 2: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		runGame(games.LcmGame, "Find the smallest common multiple of given numbers.")
	case 2:
		runGame(games.ProgressionGame, "What number is missing in the progression?")
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}
