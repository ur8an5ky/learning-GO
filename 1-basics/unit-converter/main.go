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

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		command := strings.SplitN(line, " ", 4)
		for i := range command {
			command[i] = strings.ToLower(command[i])
		}

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
				}
			}
			continue
		}
		if command[0] == "exit" || command[0] == "quit" {
			fmt.Println("See you later!")
			return
		}

		category := units.TypeOfCategory(command[1], command[3])

		switch category {
		case units.CategoryTemperature:
			result, err := u.ConvertTemperature(command)
			if err != nil {
				fmt.Printf("Error while converting unit %q to unit type: %q\n", command, err)
			}
			fmt.Printf("%s %s = %g %s\n", command[0], command[1], result, command[3])
		case units.CategoryLength:
			result, err := u.ConvertLength(command)
			if err != nil {
				fmt.Printf("Error while converting unit %q to unit type: %q\n", command, err)
			}
			fmt.Printf("%s %s = %g %s\n", command[0], command[1], result, command[3])
		case units.CategoryWeight:
			result, err := u.ConvertWeight(command)
			if err != nil {
				fmt.Printf("Error while converting unit %q to unit type: %q\n", command, err)
			}
			fmt.Printf("%s %s = %g %s\n", command[0], command[1], result, command[3])
		case units.CategoryUnknown:
			fmt.Printf("Unknown help topic: %q. Type 'help' to see available categories.\n", command[1])
		}
	}
}
