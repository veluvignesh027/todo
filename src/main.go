package main

import (
	"errors"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	backend "github.com/todo/src/back-end"
	configend "github.com/todo/src/config-end"
)

var templates *template.Template
var tempString string

type CurrentUser struct {
	Name string
	File string
	Time time.Time
}

var nowUsing CurrentUser //holds the current user details

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)

	configend.SetConfigs()
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
	templates = template.Must(template.ParseFiles(templatePaths...))
}

func main() {

	umanager := backend.UserManager{}
	umanager.ReadUsersFile() // read the user list from previous sessions

	//Root Page - Login
	http.HandleFunc("/", indexpage)
	http.HandleFunc("/validateuser", validateUser)

	//Home Page - Create, View, Delete Tasks
	http.HandleFunc("/tasks/createtask", CreateTask)
	http.HandleFunc("/tasks/deleteall", DeleteTasks)

	//Registering User
	http.HandleFunc("/register", registerPage)
	http.HandleFunc("/registeruser", registerUser)
	http.HandleFunc("/send-otp", sendotp)

	port := ":" + strconv.Itoa(configend.CurrentConfig.Portnum)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Func Handle to create new task and save it into memory
func CreateTask(w http.ResponseWriter, r *http.Request) {

}

// Func Handle to Delete All the tasks
func DeleteTasks(w http.ResponseWriter, r *http.Request) {
}

// Func to send a OTP to the registered e-mail
func sendotp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email not provided", http.StatusBadRequest)
		return
	}
	randomNumber := rand.Intn(900000) + 100000
	tempString = strconv.Itoa(randomNumber)
	backend.SendMailToDest(email, strconv.Itoa(randomNumber))

}

// Function to validate the user - if true creates a file for the user task data
// if not send unauthoraized
func validateUser(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	upass := r.FormValue("password")
	log.Println(uname)
	log.Println(upass)
	tmpUser := backend.User{
		Username: uname,
		Password: upass,
	}
	umgr := backend.UserManager{}
	err := umgr.ValidateUser(tmpUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid User Credintials.", http.StatusUnauthorized)
		return
	}
	log.Println("Valid user, Logging in...")

	// Adding current user details
	nowUsing.Name = uname
	nowUsing.File = configend.CurrentConfig.ContentStore + uname + "-tasks.json"
	nowUsing.Time = time.Now()

	//check for user tasks file
	_, err = os.Stat(nowUsing.File)
	if errors.Is(err, os.ErrNotExist) {
		fp, err := os.Create(nowUsing.File)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Tasks File ", fp.Name(), "created for user ", nowUsing.Name)
	}

	err = templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Func Handle for index page (root)
func indexpage(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

// Func Handle for Regitering new user page
func registerPage(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

// Func Handle for Register the user (/registeruser)
func registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	cpassword := r.FormValue("confirm_password")
	email := r.FormValue("email")
	otp := r.FormValue("otp")

	if cpassword != password {
		log.Println("passwords didn't matched!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if otp != tempString {
		log.Println("OTP didn't matched!")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	user := backend.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	uMgr := backend.UserManager{}

	err := uMgr.SaveRegisteredUser(user)
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
