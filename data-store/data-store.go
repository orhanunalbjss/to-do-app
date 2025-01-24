package data_store

import (
	"fmt"
)

type Todo struct {
	ID        int
	Item      string
	Completed bool
}

var Todos = []Todo{
	{ID: 1, Item: "Buy a new backpack", Completed: false},
	{ID: 2, Item: "Book flight tickets for holiday", Completed: false},
	{ID: 3, Item: "Book car servicing", Completed: false},
}

func (todo Todo) String() string {
	return fmt.Sprintf("%2d: %s [%s]\n", todo.ID, todo.Item, getStatus(todo))
}

func getStatus(todo Todo) string {
	if todo.Completed {
		return "y"
	}
	return "n"
}

func AddTodo(item string) {
	todo := Todo{ID: len(Todos) + 1, Item: item, Completed: false}
	Todos = append(Todos, todo)

	fmt.Println("Todo added")
}

func ListTodos() {
	for _, todo := range Todos {
		fmt.Printf("%2d: %s [%s]\n", todo.ID, todo.Item, getStatus(todo))
	}
}

func UpdateTodo(id int, newItem string) {
	if id >= 1 && id <= len(Todos) {
		Todos[id-1].Item = newItem
		fmt.Println("Todo updated")
	} else {
		fmt.Println("Invalid id:", id)
	}
}

func DeleteTodo(id int) {
	if id >= 1 && id <= len(Todos) {
		Todos = append(Todos[:id-1], Todos[id:]...)
		fmt.Println("Todo deleted")
	} else {
		fmt.Println("Invalid id:", id)
	}
}

func MarkTodoCompleted(id int) {
	if id >= 1 && id <= len(Todos) {
		Todos[id-1].Completed = true
		fmt.Println("Todo marked as completed")
	} else {
		fmt.Println("Invalid id:", id)
	}
}
