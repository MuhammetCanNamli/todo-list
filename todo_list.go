package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	tasks := loadTasks()

	for {
		fmt.Println("\n-_-_-_-ToDo List-_-_-_-")
		fmt.Println("-----------------------")
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. Show Task/Tasks")
		fmt.Println("4. Mark as Completed")
		fmt.Println("5. Mark as Uncompleted")
		fmt.Println("6. Save")
		fmt.Println("7. Delete Save File")
		fmt.Println("8. Exit")
		fmt.Println("-----------------------")
		fmt.Print("Option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(&tasks)
		case 2:
			deleteTask(&tasks)
		case 3:
			listTasks(tasks)
		case 4:
			markComp(&tasks)
		case 5:
			markUncomp(&tasks)
		case 6:
			saveTasks(tasks)
		case 7:
			deleteSave(&tasks)
		case 8:
			fmt.Println("\nExiting the program...")
			var confirm string
			for true {
				fmt.Print("Program state will be saved. Do you want to continue? (Y/N) ")
				fmt.Scanln(&confirm)
				if confirm == "Y" || confirm == "y" {
					saveTasks(tasks)
					fmt.Println("Program state saved.")
					break
				} else if confirm == "N" || confirm == "n" {
					fmt.Println("Program state not saved.")
					break
				} else {
					fmt.Println("Incorrect keystroke made!")
				}
			}
			fmt.Println()
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

	*tasks = append(*tasks, Task{Name: taskName, Done: false})
	fmt.Println("New task added: ", taskName)
}

func deleteTask(tasks *[]Task) {
	listTasks(*tasks)
	fmt.Print("\nEnter the number of the task you want to delete: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum <= 0 || taskNum > len(*tasks) {
		fmt.Println("Invalid task number!")
		return
	}

	taskToRemove := taskNum - 1
	fmt.Printf("Removing task: %s\n", (*tasks)[taskToRemove].Name)
	*tasks = append((*tasks)[:taskToRemove], (*tasks)[taskToRemove+1:]...)
}

func listTasks(tasks []Task) {
	fmt.Println("\nThings to do: ")
	for i, task := range tasks {
		status := "Incomplete"
		if task.Done {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.Name, status)
	}
}

func markComp(tasks *[]Task) {
	listTasks(*tasks)
	fmt.Print("\nEnter the number of the task you want to mark as completed: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum <= 0 || taskNum > len(*tasks) {
		fmt.Println("Invalid task number!")
		return
	}

	if (*tasks)[taskNum-1].Done {
		fmt.Println("This task is already marked as completed!")
		return
	}

	(*tasks)[taskNum-1].Done = true
	fmt.Println("Task marked as completed.")
}

func markUncomp(tasks *[]Task) {
	listTasks(*tasks)
	fmt.Print("\nEnter the number of the task you want to mark as uncompleted: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum <= 0 || taskNum > len(*tasks) {
		fmt.Println("Invalid task number!")
		return
	}

	if !((*tasks)[taskNum-1].Done) {
		fmt.Println("This task is already marked as incomplete!")
		return
	}

	(*tasks)[taskNum-1].Done = false
	fmt.Println("Task marked as uncompleted.")
}

func loadTasks() []Task {
	file, err := os.OpenFile("tasks.gob", os.O_RDONLY, 0666)
	if err != nil {
		return []Task{}
	}
	defer file.Close()

	var tasks []Task
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return []Task{}
	}
	return tasks
}

func saveTasks(tasks []Task) {
	file, err := os.Create("tasks.gob")
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks: ", err)
	}

	fmt.Println("\nThe tasks have been saved successfully.")
}

func deleteSave(tasks *[]Task) {
	var confirm string
	for true {
		fmt.Print("\nAre you sure you want to delete save file? (Y/N) ")
		fmt.Scanln(&confirm)
		if confirm == "Y" || confirm == "y" {
			err := os.Remove("tasks.gob")
			if err != nil {
				fmt.Println("Error deleting task file: ", err)
				return
			}
			fmt.Println("Save deleted.")
			*tasks = []Task{}
			break
		} else if confirm == "N" || confirm == "n" {
			fmt.Println("Operation cancelled.")
			break
		} else {
			fmt.Println("Incorrect keystroke made!")
		}
	}
}
