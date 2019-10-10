package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	createUserInput struct {
		Username string `json:"username"`
	}

	updateUserInput struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
)

func getTodoer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		userID, _ := strconv.Atoi(id)

		result := models.TodoerModel.Read(userID)
		jsonInfo, _ := json.Marshal(result)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonInfo)
	})
}

func createTodoer() http.Handler {
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

		var input createUserInput

		err = json.Unmarshal(body, &input)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		user := models.TodoerModel.Add(input.Username)

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

func updateTodoer() http.Handler {
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

		var input updateUserInput

		err = json.Unmarshal(body, &input)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		user := models.TodoerModel.Update(input.ID, input.Username)

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

func deleteTodoer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		userID, _ := strconv.Atoi(id)

		result := models.TodoerModel.Delete(userID)
		jsonInfo, _ := json.Marshal(result)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonInfo)
	})
}
