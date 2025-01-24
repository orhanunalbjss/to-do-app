package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	completed bool
}

var tasks []Task

func addTask(taskName string) {
	task := Task{TaskName: taskName, completed: false}
	tasks = append(tasks, task)

	fmt.Println("Added task:", taskName)
}

func listTasks() {
	for i, task := range tasks {
		status := "n"
		if task.completed {
			status = "d"
		}

		fmt.Printf("%2d: %s [%s]\n", i+1, task.TaskName, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].completed = true
		fmt.Println("Marked task as completed:", tasks[index-1].TaskName)
	} else {
		fmt.Println("Invalid index:", index)
	}
}

func editTask(index int, newTaskName string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].TaskName = newTaskName
		fmt.Println("Edited task", tasks[index-1].TaskName)
	} else {
		fmt.Println("Invalid index:", index)
	}
}

func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		taskName := tasks[index-1].TaskName
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Deleted task:", taskName)
	} else {
		fmt.Println("Invalid index:", index)
	}
}

func main() {
	var indexInput int
	var taskName string

	fmt.Println("Options")
	fmt.Println("1. Add a new task")
	fmt.Println("2. List tasks")
	fmt.Println("3. Edit task")
	fmt.Println("4. Delete task")
	fmt.Println("5. Mark completed")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter choice (1, 2, 3, 4, 5, 6): ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice:", choice)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName = scanner.Text()
			addTask(taskName)
		case 2:
			listTasks()
		case 3:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName = scanner.Text()
			editTask(indexInput, taskName)
		case 4:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			deleteTask(indexInput)
		case 5:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			markCompleted(indexInput)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice:", choice)
		}
	}
}
