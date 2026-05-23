package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (tl *TaskList) Add(task string) Task {
	t := NewTask(tl.nextID, task)
	tl.Tasks = append(tl.Tasks, t)

	tl.nextID++
	return t
}

func (tl *TaskList) List() {
	if len(tl.Tasks) == 0 {
		fmt.Println("(no tasks)")
		return
	}
	fmt.Println("List of tasks:")

	for _, t := range tl.Tasks {
		isDone := " "
		if t.Done {
			isDone = "X"
		}
		fmt.Printf("\t[%s]  #%d: %s\n", isDone, t.ID, t.Text)
	}
}

func (tl *TaskList) Done(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Done = true
			return nil
		}
	}

	return fmt.Errorf("task #%d not found", id)
}

func (tl *TaskList) Delete(id int) error {
	it := -1
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			it = i
			break
		}
	}
	if it == -1 {
		return fmt.Errorf("task #%d not found", id)
	}

	tl.Tasks = append(tl.Tasks[:it], tl.Tasks[it+1:]...)
	return nil
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
			help()
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
