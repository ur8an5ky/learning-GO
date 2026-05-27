package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ur8an5ky/learning-GO/1-basics/to-do-list/tasks"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to to-do! Type 'help' to see the available commands.")
	tl := tasks.NewTaskList()

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

		command := strings.SplitN(line, " ", 2)
		command[0] = strings.ToLower(command[0])

		switch command[0] {
		case "quit", "exit":
			fmt.Println("See you later!")
			return
		case "help":
			tasks.Help()
		case "add":
			if len(command) < 2 {
				fmt.Println("Usage: add <text>")
				continue
			}
			task := tl.Add(command[1])
			fmt.Printf("Added task #%d: %s\n", task.ID, task.Text)
		case "list":
			tl.List()
		case "done":
			if len(command) != 2 {
				fmt.Println("Usage: done <id>.")
				continue
			}
			id, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("ID must be a number.")
				continue
			}
			if err := tl.Done(id); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Task #%d marked as done\n", id)
		case "delete":
			if len(command) != 2 {
				fmt.Println("Usage: delete <id>.")
				continue
			}
			id, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("ID must be a number!")
				continue
			}
			if err := tl.Delete(id); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Task #%d deleted\n", id)
		default:
			fmt.Printf("Unknown command: %q. Type 'help' for available commands.\n", command[0])
		}
	}
}
