package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"workshop/database"
)

type (
	createTodoInput struct {
		Description string `json:"description"`
		Creator     int    `json:"creator"`
	}

	updateTodoInput struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
	}
)

func getTodo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		userID, _ := strconv.Atoi(id)

		result := models.TodoModel.Read(userID)
		jsonInfo, _ := json.Marshal(result)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonInfo)
	})
}

func getTodosFromUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("creator")
		userID, _ := strconv.Atoi(id)

		result := models.TodoModel.ReadAll(userID)
		jsonInfo, _ := json.Marshal(result)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonInfo)
	})
}

func createTodo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var input createTodoInput

		err = json.Unmarshal(body, &input)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		user := models.TodoModel.Add(database.Todo{Description: input.Description, Creator: input.Creator})

		output, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

func updateTodo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var input updateTodoInput

		err = json.Unmarshal(body, &input)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		user := models.TodoModel.Update(database.Todo{ID: input.ID, Description: input.Description})

		output, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	})
}

func deleteTodo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		todoID, _ := strconv.Atoi(id)

		result := models.TodoModel.Delete(todoID)
		jsonInfo, _ := json.Marshal(result)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonInfo)
	})
}
