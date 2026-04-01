package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
)

// Todo represents a single todo item
type Todo struct {
	ID   int
	Task string
	Done bool
}

// Todos stores all todo items
var Todos = []Todo{}

// nextID keeps track of the next available ID for new todos
var nextID = 1

// PrintList displays all todo items with their details
func PrintList(todos *[]Todo) {
	fmt.Println("\n--------------")
	fmt.Println("All Todos")
	fmt.Println("--------------\n")
	for _, todo := range *todos {
		status := "[ ]"
		if todo.Done {
			status = "[x]"
		}
		fmt.Printf("%s ID: %d | Task: %s\n", status, todo.ID, todo.Task)
		fmt.Println("------------------")
	}
}

// AddTodo creates and appends a new task to the todo list
func AddTodo(todos *[]Todo, task string) {
	newTodo := Todo{
		ID:   nextID,
		Task: task,
		Done: false,
	}
	*todos = append(*todos, newTodo)
	nextID++
}

// Remove deletes a task from the todo list by ID
func Remove(todos *[]Todo, id int) {
	*todos = slices.DeleteFunc(*todos, func(t Todo) bool {
		return t.ID == id
	})
}

// ToggleDone flips the Done status of a task by ID
func ToggleDone(todos *[]Todo, id int) {
	index := slices.IndexFunc(*todos, func(t Todo) bool {
		return t.ID == id
	})
	if index != -1 {
		(*todos)[index].Done = !(*todos)[index].Done
	}
}

// EditTask updates the task title by ID
func EditTask(todos *[]Todo, id int, newTask string) {
	index := slices.IndexFunc(*todos, func(t Todo) bool {
		return t.ID == id
	})
	if index != -1 {
		(*todos)[index].Task = newTask
	}
}

// Init runs the main menu loop for the todo application
func Init() bool {
	for {
		cmd := ""

		c := exec.Command("clear")
		c.Stdout = os.Stdout

		fmt.Println("Todo Application")
		fmt.Println("Enter a command (q=quit, l=list, a=add, e=edit, r=remove, t=toggle):")
		fmt.Scan(&cmd)

		if cmd == "q" {
			fmt.Println("Goodbye!")
			return false
		} else if cmd == "list" || cmd == "l" {
			PrintList(&Todos)
			Init()
			return false
		} else if cmd == "add" || cmd == "a" {
			fmt.Println("Enter the task:")
			var task string
			fmt.Scan(&task)
			AddTodo(&Todos, task)
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
			fmt.Println("Enter the task ID to edit:")
			fmt.Scan(&cmd)
			id, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Error: Please enter a valid number.")
				return false
			}
			fmt.Println("Enter the new task:")
			var newTask string
			fmt.Scan(&newTask)
			EditTask(&Todos, id, newTask)
			Init()
			return false
		} else if cmd == "toggle" || cmd == "t" {
			fmt.Println("Enter the task ID to toggle:")
			fmt.Scan(&cmd)
			id, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Error: Please enter a valid number.")
				return false
			}
			ToggleDone(&Todos, id)
			Init()
			return false
		} else {
			Init()
			return true
		}
	}
}

func main() {
	Todos = append(Todos, Todo{nextID, "Learn Go structs", false})
	nextID++
	Todos = append(Todos, Todo{nextID, "Practice Go maps", false})
	nextID++

	fmt.Println("Starting Todo Application...")
	Init()
}
