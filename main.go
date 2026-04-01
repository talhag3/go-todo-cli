package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices" // Helper Methods for a slice
	"strconv"
)

type Todo struct {
	Id   int
	Task string
	Done bool
}

// Todos stores the list of todo items
var Todos = []Todo{}

// PrintList displays all todo items with their index numbers
func PrintList(todos *[]Todo) {
	fmt.Println("\n--------------\nAll Todos\n--------------\n")
	for _, todo := range *todos {
		fmt.Println("ID ", todo.Id)
		fmt.Println("Task ", todo.Task)
		fmt.Println("Done ", todo.Done)
		fmt.Println("------------------")
	}
}

// AddTodo appends a new task to the todo list
func AddTodo(todos *[]Todo) {
	*todos = append(*todos, Todo{3, "Task3", false})
}

// Remove deletes a task from the todo list at the given index
func Remove(todos *[]Todo, id int) {
	*todos = slices.DeleteFunc(*todos, func(t Todo) bool {
		return t.Id == id
	})
}

func ToggleDone(todos *[]Todo, id int, done bool) {
	*todos = slices.DeleteFunc(*todos, func(t Todo) bool {
		return t.Done == done
	})
}

func EditTask(todos *[]Todo, id int, Task string) {
	Index := slices.IndexFunc(*todos, func(t Todo) bool {
		return t.Id == id
	})

	// Check if the index is valid
	if Index != -1 {
		(*todos)[Index].Task = Task
	}
}

// Init runs the main menu loop for the todo application
func Init() bool {
	for {
		cmd := ""

		// Clear terminal (commented out for now)
		c := exec.Command("clear")
		c.Stdout = os.Stdout

		fmt.Println("Todo Application")
		fmt.Println("Enter a command (q=quit, l=list, a=add, e=edit, r=remove):")
		fmt.Scan(&cmd)

		if cmd == "q" {
			fmt.Println("Goodbye!")
			return false
		} else if cmd == "list" || cmd == "l" {
			PrintList(&Todos)
			return false
		} else if cmd == "add" || cmd == "a" {
			AddTodo(&Todos)
			Init()
			return false
		} else if cmd == "remove" || cmd == "r" {
			fmt.Println("Enter the task ID to remove:")
			fmt.Scan(&cmd)
			id, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Error: Please enter a valid number.")
				return false
			}
			Remove(&Todos, id)
			Init()
			return false
		} else if cmd == "edit" || cmd == "e" {
			fmt.Println("Enter the task ID to Edit:")
			fmt.Scan(&cmd)
			id, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Error: Please enter a valid number.")
				return false
			}
			fmt.Println("Edit the task")
			fmt.Scan(&cmd)
			EditTask(&Todos, id, cmd)
			Init()
			return false
		} else {
			Init()
			return true
		}
	}
}

func main() {

	var t1 = Todo{1, "Im task 1", false}

	Todos = append(Todos, t1)

	Todos = append(Todos, Todo{2, "Im task 2", false})

	fmt.Println("Starting Todo Application...")
	Init()
}
