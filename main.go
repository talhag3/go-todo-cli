package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Todos stores the list of todo items
var Todos = []string{"Task1", "Task2"}

// PrintList displays all todo items with their index numbers
func PrintList(todos *[]string) {
	for i, val := range *todos {
		fmt.Printf("%d - %s\n", i+1, val)
	}
}

// AddTodo appends a new task to the todo list
func AddTodo(todos *[]string) {
	*todos = append(*todos, "Task3")
}

// Remove deletes a task from the todo list at the given index
func Remove(todos *[]string, index int) {
	t := *todos
	*todos = append(t[:index], t[index+1:]...)
}

// Init runs the main menu loop for the todo application
func Init() bool {
	for {
		cmd := ""

		// Clear terminal (commented out for now)
		c := exec.Command("clear")
		c.Stdout = os.Stdout

		fmt.Println("Todo Application")
		fmt.Println("Enter a command (q=quit, l=list, a=add, r=remove):")
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
		} else {
			Init()
			return true
		}
	}
}

func main() {
	fmt.Println("Starting Todo Application...")
	Init()
}
