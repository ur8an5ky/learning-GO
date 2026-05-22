package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

type TaskList struct {
	Tasks []Task
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  add <text>   - add a new task")
	fmt.Println("  list         - show all tasks")
	fmt.Println("  done <id>    - mark task as done")
	fmt.Println("  delete <id>  - delete a task")
	fmt.Println("  help         - show this help")
	fmt.Println("  quit, exit   - exit the program")
}

func add(task string) {
	fmt.Println("Adding task:", task)
}

func list() {
	fmt.Println("Listing tasks:")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to to-do! Type 'help' to see the available commands.")

	for {
		fmt.Print(">  ")

		if !scanner.Scan() {
			fmt.Println("\nSee you later!")
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		command := strings.SplitN(line, " ", 2)
		if (len(command) >= 2) && (command[0] != "add") {
			fmt.Println("Command:", command, "takes only one argument! Try again!")
			continue
		}

		switch command[0] {
		case "quit", "exit":
			fmt.Println("See you later!")
			return
		case "help":
			help()
		case "add":
			add(command[1])
		case "list":
			list()
		default:
			fmt.Printf("Unknown command: %q. Type 'help' for available commands.\n", command[0])
		}
	}
}
