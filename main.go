package main

import (
	todo "go_final_project/todo"
	"os"
)

func main() {
	_ = todo.DbOpen()

	port := ":" + os.Getenv("TODO_PORT")
	if port == ":" {
		port = ":7540"
	}

	todo.StartServer(port)

}
