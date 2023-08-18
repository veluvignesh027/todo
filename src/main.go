package main

import (
	"log"
	"net/http"
	"os"

	backend "github.com/todo/src/back-end"
)

var h backend.Handler

func main() {
	log.SetFlags(log.Lshortfile)

	//HandleFunc for all handles
	http.HandleFunc("/", h.RootHandler)
	http.HandleFunc("/home", h.HomeHandler)
	http.HandleFunc("/tasks/create", h.CreateHandler)
	http.HandleFunc("/view", h.ViewHandler)
	http.HandleFunc("/delete", h.DeleteHandler)
	http.HandleFunc("/submit", h.SubmitHandler)
	http.HandleFunc("/validate", h.ValidateHandler)
	http.HandleFunc("/usercreate", h.CreateUserHandler)
	http.HandleFunc("/adduser", h.AddUserHandler)

	log.Println("Server Listening on port 9000.....")
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		log.Println("Error", err)
		os.Exit(-1)
	}

}
