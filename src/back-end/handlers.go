package backend

import (
	"errors"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	configend "github.com/todo/src/config-end"
)

var Templates *template.Template
var tempString string
var store = sessions.NewCookieStore([]byte("AbCdEfGhIjK"))

type CurrentUser struct {
	Name string
	File string
	Time time.Time
}

// var nowUsing CurrentUser //holds the current user details

// Func Handle to create new task and save it into memory
func CreateTask(w http.ResponseWriter, r *http.Request) {
	cMgr := ContentManager{}

	session, _ := store.Get(r, "session-name")
	username, ok := session.Values["username"].(string)
	if !ok {
		return
	}
	log.Println("Creating Task for Session ", session, " ", session.Values["username"])

	task := ContentTasks{
		TaskName: r.FormValue("task"),
		Date:     r.FormValue("date"),
		Time:     r.FormValue("time"),
	}
	cMgr.SaveTasksInFile(task, username)
	Templates.ExecuteTemplate(w, "/home", task.TaskName+"Saved!")

}

// Func Handle to Delete All the tasks
func DeleteTasks(w http.ResponseWriter, r *http.Request) {
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	delete(session.Values, "username")
	session.Save(r, w)
}

// Func to send a OTP to the registered e-mail
func Sendotp(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email not provided", http.StatusBadRequest)
		return
	}
	randomNumber := rand.Intn(900000) + 100000
	tempString = strconv.Itoa(randomNumber)
	SendMailToDest(email, strconv.Itoa(randomNumber))

}

// Function to validate the user - if true creates a file for the user task data
// if not send unauthoraized
func ValidateUser(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	upass := r.FormValue("password")
	log.Println(uname)
	log.Println(upass)
	tmpUser := User{
		Username: uname,
		Password: upass,
	}
	umgr := UserManager{}
	err := umgr.ValidateUser(tmpUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid User Credintials.", http.StatusUnauthorized)
		return
	}
	log.Println("Valid user, Logging in...")

	//set a session
	session, _ := store.New(r, "session-name")
	session.Values["username"] = uname
	session.Save(r, w)

	log.Println("New Session creted ", session, " ", session.Values["username"])
	// Adding current user details
	uName := uname
	uFile := configend.CurrentConfig.ContentStore + uname + "-tasks.json"
	uTime := time.Now()

	log.Println("For User : ", uName, "  ", uTime)
	//check for user tasks file
	_, err = os.Stat(uFile)
	if errors.Is(err, os.ErrNotExist) {
		fp, err := os.Create(uFile)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Tasks File ", fp.Name(), "created for user ", uName)
	}

	err = Templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Func Handle for index page (root)
func Indexpage(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

// Func Handle for Regitering new user page
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

// Func Handle for Register the user (/registeruser)
func RegisterUser(w http.ResponseWriter, r *http.Request) {
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
	user := User{
		Username: username,
		Password: password,
		Email:    email,
	}

	uMgr := UserManager{}

	err := uMgr.SaveRegisteredUser(user)
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
