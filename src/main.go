package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"

	backend "github.com/todo/src/back-end"
	configend "github.com/todo/src/config-end"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	configfilepath := *flag.String("configPath", "\\CACHE_BASEDIR\\todo-app\\todo\\configs\\", "Path to config file")
	err := configend.GetConfigParams(configfilepath)
	if err != nil {
		configend.SetConfigs(configfilepath)
	}
	configend.PrintConfigs()
}

func init() {
	// Specify correct file paths for your templates
	templatePaths := []string{
		configend.CurrentConfig.FrontEndPath + "index.html",
		configend.CurrentConfig.FrontEndPath + "register.html",
		configend.CurrentConfig.FrontEndPath + "home.html",
		configend.CurrentConfig.FrontEndPath + "create.html",
	}

	// Parse the templates
	backend.Templates = template.Must(template.ParseFiles(templatePaths...))
}
func main() {

	umanager := backend.UserManager{}
	umanager.ReadUsersFile() // read the user list from previous sessions

	//Root Page - Login
	http.HandleFunc("/", backend.Indexpage)
	http.HandleFunc("/validateuser", backend.ValidateUser)

	//Home Page - Create, View, Delete Tasks
	http.HandleFunc("/tasks/createtask", backend.CreateTask)
	http.HandleFunc("/tasks/deleteall", backend.DeleteTasks)

	//Registering User
	http.HandleFunc("/register", backend.RegisterPage)
	http.HandleFunc("/registeruser", backend.RegisterUser)
	http.HandleFunc("/send-otp", backend.Sendotp)

	//For - Logout
	http.HandleFunc("/logout", backend.Logout)

	port := ":" + strconv.Itoa(configend.CurrentConfig.Portnum)
	log.Fatal(http.ListenAndServe(port, nil))
}
