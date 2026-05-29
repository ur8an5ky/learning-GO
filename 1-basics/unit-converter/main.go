package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ur8an5ky/learning-GO/1-basics/unit-converter/help"
	"github.com/ur8an5ky/learning-GO/1-basics/unit-converter/units"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Unit Converter! Type 'help' to see the available commands.")
	u := units.NewUnits()

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			fmt.Println("\nSee you later!")
			break
		}

		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if line == "" {
			continue
		}

		command := strings.SplitN(line, " ", 4)

		if command[0] == "help" {
			if len(command) == 1 {
				help.Help()
			} else {
				switch command[1] {
				case "temperature":
					help.HelpTemperature()
				case "length":
					help.HelpLength()
				case "weight":
					help.HelpWeight()
				default:
					fmt.Printf("Unknown help topic: %q. Available: temperature, length, weight.\n", command[1])
				}
			}
			continue
		}
		if command[0] == "exit" || command[0] == "quit" {
			fmt.Println("See you later!")
			return
		}

		if len(command) != 4 {
			fmt.Println("Invalid command. Type 'help' for usage.")
			continue
		}

		if command[2] != "to" {
			fmt.Printf("Expected 'to' as separator, got %q. Type 'help' for usage.\n", command[2])
			continue
		}

		result, err := u.Convert(command)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("%s %s = %g %s\n", command[0], command[1], result, command[3])
	}
}
