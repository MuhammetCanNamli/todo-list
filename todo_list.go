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
		fmt.Println("\n-_-_-_-ToDo List-_-_-_-")
		fmt.Println("-----------------------")
		fmt.Println("1. Add")
		fmt.Println("2. Show")
		fmt.Println("3. Mark as Completed")
		fmt.Println("4. Mark as Uncompleted")
		fmt.Println("5. Exit")
		fmt.Println("-----------------------")
		fmt.Print("Option: ")

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
			fmt.Println("Exiting the program...\n")
			os.Exit(0)
		default:
			fmt.Println("\nInvalid option!")
		}
	}
}

func addTask(tasks *[]Task) {
	fmt.Print("\nEnter the to-do: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskName := scanner.Text()

	*tasks = append(*tasks, Task{name: taskName, done: false})
	fmt.Println("New task added: ", taskName)
}

func listTasks(tasks []Task) {
	fmt.Println("\nThings to do: ")
	for i, task := range tasks {
		status := "Incomplete"
		if task.done {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.name, status)
	}
}

func markComp(tasks *[]Task) {
	listTasks(*tasks)
	fmt.Print("Enter the number of the task you want to mark as completed: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum <= 0 || taskNum > len(*tasks) {
		fmt.Println("Invalid task number!")
		return
	}

	if (*tasks)[taskNum-1].done {
		fmt.Println("This task is already marked as completed!")
		return
	}

	(*tasks)[taskNum-1].done = true
	fmt.Println("Task marked as completed.")
}

func markUncomp(tasks *[]Task) {
	listTasks(*tasks)
	fmt.Print("Enter the number of the task you want to mark as uncompleted: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum <= 0 || taskNum > len(*tasks) {
		fmt.Println("Invalid task number!")
		return
	}

	if !((*tasks)[taskNum-1].done) {
		fmt.Println("This task is already marked as incomplete!")
		return
	}

	(*tasks)[taskNum-1].done = false
	fmt.Println("Task marked as uncompleted.")
}
