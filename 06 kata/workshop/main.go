package main

import (
	"log"
	"net/http"
	"workshop/database"
)

type modelStruct struct {
	TodoerModel database.TODOER
	TodoModel   database.TODO
}

var models modelStruct

func main() {
	log.Print("Program Start")

	var toderModel database.TODOER = &database.TodoerModel{}
	var todoModel database.TODO = &database.TodoModel{}

	models = modelStruct{toderModel, todoModel}

	startServer()
}

func startServer() {
	mux := http.NewServeMux()
	mux.Handle("/todoer/read", getTodoer())
	mux.Handle("/todoer/create", createTodoer())
	mux.Handle("/todoer/update", updateTodoer())
	mux.Handle("/todoer/delete", deleteTodoer())

	mux.Handle("/todo/readall", getTodosFromUser())
	mux.Handle("/todo/read", getTodo())
	mux.Handle("/todo/create", createTodo())
	mux.Handle("/todo/update", updateTodo())
	mux.Handle("/todo/delete", deleteTodo())

	log.Printf("listening on Port 8888")

	err := http.ListenAndServe(":8888", mux)
	log.Fatal(err)
}
