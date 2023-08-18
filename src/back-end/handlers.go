package backend

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Handler struct{}

// Handler for "/"
func (h *Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\CACHE_BASEDIR\\todo-app\\todo\\src\\front-end\\index.html")
}

// Handler for "/home"
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\CACHE_BASEDIR\\todo-app\\todo\\src\\front-end\\home.html")
}

// Handler for "/create"
func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\CACHE_BASEDIR\\todo-app\\todo\\src\\front-end\\create.html")
}

// Handler for "/submit"
func (h *Handler) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	ut := Util{}
	str := ut.ReadBodyToString(r)

	ok := ut.SaveToContentFile(str)
	if ok {
		log.Println("Saved in content file!")
	} else {
		log.Println("Can not save the content..")
	}

	time.Sleep(time.Second * 2)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// Handler for "/view"
func (h *Handler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	mbyte, err := os.ReadFile("C:\\CACHE_BASEDIR\\todo-app\\todo\\content-store\\contents.json")
	if err != nil {
		log.Println(err)
	}
	w.Write(mbyte)
}

// Handler for "/delete"
func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "")
}

// Handler for "/validate"
func (h *Handler) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	ut := Util{}

	s := ut.ReadBodyToString(r)

	um := UserManager{}
	tmp, err := um.GetUserStruct(s)
	if err != nil {
		log.Println(err)
	}

	err = um.ValidateUser(tmp)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Invalid user...Try login with valid user."))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\CACHE_BASEDIR\\todo-app\\todo\\src\\front-end\\getuser.html")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	ut := Util{}
	s := ut.ReadBodyToString(r)

	um := UserManager{}
	temp, err := um.ValidateUserStruct(s)
	if err != nil {
		log.Println(err)
	}

	err = um.SaveUser(temp)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second * 2)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
