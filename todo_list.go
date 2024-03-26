package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name string
	done bool
}

func main() {
	tasks := []Task{}

	for {
		fmt.Println("ToDo List")
		fmt.Println("1. Add")
		fmt.Println("2. Show")
		fmt.Println("3. Mark as Completed")
		fmt.Println("4. Mark as Uncompleted")
		fmt.Println("5. Exit")
		fmt.Println("Option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(&tasks)
		case 2:
			listTasks(tasks)
		case 3:
			markComp(&tasks)
		case 4:
			markUncomp(&tasks)
		case 5:
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func addTask(tasks *[]Task) {
	fmt.Print("Enter the to-do: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskName := scanner.Text()

	*tasks = append(*tasks, Task{name: taskName, done: false})
	fmt.Println("New task added: ", taskName)
}

func listTasks(tasks []Task) {

}

func markComp(tasks *[]Task) {

}

func markUncomp(tasks *[]Task) {

}
