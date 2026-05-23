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

func NewTask(id int, text string) Task {
	return Task{
		ID:   id,
		Text: text,
		Done: false,
	}
}

type TaskList struct {
	Tasks  []Task
	nextID int
}

func (tl *TaskList) Add(task string) {
	t := NewTask(tl.nextID, task)
	tl.Tasks = append(tl.Tasks, t)

	tl.nextID++

	fmt.Printf("Added task #%d: %s\n", t.ID, task)
}

func (tl *TaskList) List() {
	fmt.Println("List of tasks:")
	for _, t := range tl.Tasks {
		fmt.Printf("\t<%d>: %s\n", t.ID, t.Text)
	}
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("\tadd <text>   - add a new task")
	fmt.Println("\tlist         - show all tasks")
	fmt.Println("\tdone <id>    - mark task as done")
	fmt.Println("\tdelete <id>  - delete a task")
	fmt.Println("\thelp         - show this help")
	fmt.Println("\tquit, exit   - exit the program")
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to to-do! Type 'help' to see the available commands.")
	tl := TaskList{nextID: 1}

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
		command[0] = strings.ToLower(command[0])

		//if command[0] == "add" {
		//	if len(command) < 2 {
		//		fmt.Println("Command:", command, "takes at least two arguments! Try again!")
		//		continue
		//	}
		//} else {
		//	if len(command) >= 2 {
		//		fmt.Println("Command:", command, "takes only one argument! Try again!")
		//		continue
		//	}
		//}

		switch command[0] {
		case "quit", "exit":
			fmt.Println("See you later!")
			return
		case "help":
			help()
		case "add":
			if len(command) < 2 {
				fmt.Println("Command:", command, "takes at least two arguments! Try again!")
				continue
			}
			tl.Add(command[1])
		case "list":
			tl.List()
		default:
			fmt.Printf("Unknown command: %q. Type 'help' for available commands.\n", command[0])
		}
	}
}
