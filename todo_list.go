package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Task struct {
	Name     string
	Done     bool
	Deadline time.Time
	Category string
	Tags     []string
}

const taskFile = "task.gob"

func main() {
	tasks := loadTasks()

	for {
		fmt.Println("\n-_-_-_-ToDo List-_-_-_-")
		fmt.Println("-----------------------")
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. Show Task/Tasks")
		fmt.Println("4. Edit Deadline")
		fmt.Println("5. Mark as Completed")
		fmt.Println("6. Mark as Uncompleted")
		fmt.Println("7. Sort Tasks")
		fmt.Println("8. Filter Tasks by Category")
		fmt.Println("9. Filter Tasks by Tag")
		fmt.Println("10. Save")
		fmt.Println("11. Delete Save File")
		fmt.Println("12. Exit")
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
			editDeadline(&tasks)
		case 5:
			markComp(&tasks)
		case 6:
			markUncomp(&tasks)
		case 7:
			sortTasks(&tasks)
		case 8:
			filterByCategory(&tasks)
		case 9:
			filterByTag(&tasks)
		case 10:
			saveTasks(tasks)
		case 11:
			deleteSave(&tasks)
		case 12:
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter the to-do: ")
		scanner.Scan()
		taskName := scanner.Text()

		if taskName == "" {
			fmt.Println("Task name cannot be empty! Please enter a valid task name.")
			continue
		}

		exists := false
		for _, task := range *tasks {
			if task.Name == taskName {
				exists = true
				break
			}
		}

		if exists {
			fmt.Printf("Task '%s' already exists! Please add a different task.\n", taskName)
			continue
		}

		var deadlineInput string
		for {
			fmt.Print("Enter the deadline (YYYY-MM-DD): ")
			scanner.Scan()
			deadlineInput = scanner.Text()

			deadline, err := time.Parse("2006-01-02", deadlineInput)
			if err != nil {
				fmt.Println("Invalid deadline format! Please enter the deadline in the format YYYY-MM-DD.")
				continue
			}

			fmt.Print("Enter the category (optional): ")
			scanner.Scan()
			category := scanner.Text()

			fmt.Print("Enter comma-separated tags (optional): ")
			scanner.Scan()
			tagsInput := scanner.Text()
			tags := strings.Split(tagsInput, ",")

			*tasks = append(*tasks, Task{Name: taskName, Done: false, Deadline: deadline, Category: category, Tags: tags})
			fmt.Println("New task added: ", taskName)
			return
		}
	}
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
		if !task.Done && !task.Deadline.IsZero() {
			fmt.Printf(" (Deadline: %s)", task.Deadline.Format("2006-01-02"))
		}
		if task.Category != "" {
			fmt.Printf(" [Category : %s]", task.Category)
		}
		if len(task.Tags) > 0 {
			fmt.Printf(" [Tags: %s]", strings.Join(task.Tags, ", "))
		}
		fmt.Println()
	}
}

func editDeadline(tasks *[]Task) {
	listTasks(*tasks)
	fmt.Print("\nEnter the number of the task you want to edit the deadline for: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum <= 0 || taskNum > len(*tasks) {
		fmt.Println("Invalid task number")
		return
	}

	taskIndex := taskNum - 1
	fmt.Printf("Current deadline for task '%s': %s\n", (*tasks)[taskIndex].Name, (*tasks)[taskIndex].Deadline.Format("2006-01-02"))

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the new deadline (YYYY-MM-DD): ")
		scanner.Scan()
		newDeadlineInput := scanner.Text()

		newDeadline, err := time.Parse("2006-01-02", newDeadlineInput)
		if err != nil {
			fmt.Println("Invalid deadline format! Please enter he deadline in the format YYYY-MM-DD.")
			continue
		}

		(*tasks)[taskIndex].Deadline = newDeadline
		fmt.Printf("Deadline for task '%s' updated to: %s\n", (*tasks)[taskIndex].Name, newDeadline.Format("2006-01-02"))
		break
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

func sortTasks(tasks *[]Task) {
	fmt.Println("\nSort Tasks by: ")
	fmt.Println("1. Name")
	fmt.Println("2. Deadline")
	fmt.Println("3. Category")
	fmt.Println("4. Tags")
	fmt.Print("Option: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		sort.Slice(*tasks, func(i, j int) bool {
			return (*tasks)[i].Name < (*tasks)[j].Name
		})
		fmt.Println("Tasks sorted by name.")
	case 2:
		sort.Slice(*tasks, func(i, j int) bool {
			return (*tasks)[i].Deadline.Before((*tasks)[j].Deadline)
		})
		fmt.Println("Tasks sorted by deadline")
	case 3:
		sort.Slice(*tasks, func(i, j int) bool {
			return (*tasks)[i].Category < (*tasks)[j].Category
		})
		fmt.Println("Tasks sorted by category.")
	case 4:
		sort.Slice(*tasks, func(i, j int) bool {
			return strings.Join((*tasks)[i].Tags, ", ") < strings.Join((*tasks)[j].Tags, ", ")
		})
		fmt.Println("Tasks sorted by tags.")
	default:
		fmt.Println("\nInvalid option!")
	}

	listTasks(*tasks)
}

func filterByCategory(tasks *[]Task) {
	fmt.Print("\nEnter category to filter tasks by: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	category := scanner.Text()

	var filteredTasks []Task
	for _, task := range *tasks {
		if task.Category == category {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == 0 {
		fmt.Printf("No tasks found under category '%s'.\n", category)
		return
	}

	listTasks(filteredTasks)
}

func filterByTag(tasks *[]Task) {
	fmt.Print("\nEnter category to tag tasks by: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tag := scanner.Text()

	var filteredTasks []Task
	for _, task := range *tasks {
		for _, t := range task.Tags {
			if t == tag {
				filteredTasks = append(filteredTasks, task)
				break
			}
		}
	}

	if len(filteredTasks) == 0 {
		fmt.Printf("No tasks found under category '%s'.\n", tag)
		return
	}

	listTasks(filteredTasks)
}

func loadTasks() []Task {
	file, err := os.OpenFile(taskFile, os.O_RDONLY, 0666)
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
	file, err := os.Create(taskFile)
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
			err := os.Remove(taskFile)
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
