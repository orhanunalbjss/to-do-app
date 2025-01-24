package main

import (
	"bufio"
	"fmt"
	datastore "local.com/to-do-app/data-store"
	"os"
	"strconv"
)

func main() {
	var id int
	var item, newItem string

	fmt.Println("Options")
	fmt.Println("1. Add a new todo")
	fmt.Println("2. List todos")
	fmt.Println("3. Edit todo")
	fmt.Println("4. Delete todo")
	fmt.Println("5. Mark Completed")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter choice (1, 2, 3, 4, 5, 6): ")
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid choice:", choice)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter item: ")
			scanner.Scan()
			item = scanner.Text()
			datastore.AddTodo(item)
		case 2:
			datastore.ListTodos()
		case 3:
			fmt.Print("Enter id: ")
			scanner.Scan()
			id, _ = strconv.Atoi(scanner.Text())
			fmt.Print("Enter new item: ")
			scanner.Scan()
			newItem = scanner.Text()
			datastore.UpdateTodo(id, newItem)
		case 4:
			fmt.Print("Enter id: ")
			scanner.Scan()
			id, _ = strconv.Atoi(scanner.Text())
			datastore.DeleteTodo(id)
		case 5:
			fmt.Print("Enter id: ")
			scanner.Scan()
			id, _ = strconv.Atoi(scanner.Text())
			datastore.MarkCompleted(id)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice:", choice)
		}
	}
}
