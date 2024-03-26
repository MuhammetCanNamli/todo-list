package main

import (
	"fmt"
	"os"
)

func main() {

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
			addTask()
		case 2:
			listTasks()
		case 3:
			markComp()
		case 4:
			markUncomp()
		case 5:
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func addTask() {

}

func listTasks() {

}

func markComp() {

}

func markUncomp() {

}
