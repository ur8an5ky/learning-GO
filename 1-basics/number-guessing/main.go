package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	guessingGoal := rand.Intn(100) + 1
	iter := 0

	fmt.Println("I have a randomly generated number in my memory (between 1 and 100).")
	fmt.Println("Let's play a game – guess what the number is – I'll tell you whether it's 'higher' or 'lower'.")
	fmt.Println("Good luck! [Type 'quit' if you're giving up]")

	for {
		if iter > 0 {
			fmt.Println("Try again!")
		}
		fmt.Print(">  ")

		// czytanie kolejnej linii
		if !scanner.Scan() {
			fmt.Println("\nThanks for playing!")
			break
		}

		// obcięcie znaków białych
		line := strings.TrimSpace(scanner.Text())

		// jeśli podaliśmy quit to koniec
		if line == "quit" {
			fmt.Println("Better luck next time!")
			return
		}

		// konwersja inputu do inta
		guess, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("It's not a number! Try again!")
			continue
		}

		// nakierowywanie na kolejny strzał
		if guess > 100 || guess < 1 {
			fmt.Println("Your guess is out of the range. Please enter a number between 1 and 100.")
			continue
		}

		iter++

		if guess == guessingGoal {
			fmt.Printf("You guessed it in %d tries!\n", iter)
			return
		} else if guess > guessingGoal {
			fmt.Println("Lower!")
		} else if guess < guessingGoal {
			fmt.Println("Higher!")
		}
	}

}
