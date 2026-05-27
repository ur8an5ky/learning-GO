package tasks

import "fmt"

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
	tasks  []Task
	nextID int
}

func NewTaskList() TaskList {
	return TaskList{nextID: 1}
}

func (tl *TaskList) Add(task string) Task {
	t := NewTask(tl.nextID, task)
	tl.tasks = append(tl.tasks, t)

	tl.nextID++
	return t
}

func (tl *TaskList) List() {
	if len(tl.tasks) == 0 {
		fmt.Println("(no tasks)")
		return
	}
	fmt.Println("List of tasks:")

	for _, t := range tl.tasks {
		isDone := " "
		if t.Done {
			isDone = "X"
		}
		fmt.Printf("\t[%s]  #%d: %s\n", isDone, t.ID, t.Text)
	}
}

func (tl *TaskList) Done(id int) error {
	for i := range tl.tasks {
		if tl.tasks[i].ID == id {
			tl.tasks[i].Done = true
			return nil
		}
	}

	return fmt.Errorf("task #%d not found", id)
}

func (tl *TaskList) Delete(id int) error {
	it := -1
	for i := range tl.tasks {
		if tl.tasks[i].ID == id {
			it = i
			break
		}
	}
	if it == -1 {
		return fmt.Errorf("task #%d not found", id)
	}

	tl.tasks = append(tl.tasks[:it], tl.tasks[it+1:]...)
	return nil
}

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("\tadd <text>   - add a new task")
	fmt.Println("\tlist         - show all tasks")
	fmt.Println("\tdone <id>    - mark task as done")
	fmt.Println("\tdelete <id>  - delete a task")
	fmt.Println("\thelp         - show this help")
	fmt.Println("\tquit, exit   - exit the program")
	fmt.Println()
}
