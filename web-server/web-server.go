package main

import (
	"encoding/json"
	"fmt"
	datastore "local.com/to-do-app/data-store"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		for _, v := range datastore.Todos {
			_, err := fmt.Fprintf(w, v.String())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	router.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {
		var item datastore.Todo
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		datastore.AddTodo(item.Item)
	})

	router.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		var newTodo datastore.Todo
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&newTodo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		datastore.UpdateTodo(id, newTodo.Item)
	})

	router.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		datastore.DeleteTodo(id)
	})

	router.HandleFunc("PATCH /todos/{id}/mark-completed", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		datastore.MarkTodoCompleted(id)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
