package main

import (
	"bufio"
	"fmt"
	"os"
)

const rounds = 3

func runGame(game func() bool, gameName string) {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your name? ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Hello, %s!\n", name)
	fmt.Println(gameName)

	for i := 0; i < rounds; i++ {
		if !game() {
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}
	}
	fmt.Printf("Congratulations, %s!\n", name)
}
