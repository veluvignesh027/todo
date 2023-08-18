package backend

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var allUsers []User

type UserManager struct{}

type User struct {
	Name     string `json:"NAME"`
	Password string `json:"PASSWORD"`
}

func (u *UserManager) ValidateUserStruct(s string) (User, error) {
	var temp User

	ut := Util{}
	uname := ut.GetValueForKey(s, "UserName")
	upass := ut.GetValueForKey(s, "AddPassword")
	ucpass := ut.GetValueForKey(s, "ConfirmPassword")

	if upass != ucpass {
		return temp, errors.New("password Didn't matched")
	}

	temp.Name = uname
	temp.Password = upass

	return temp, nil
}

func (u *UserManager) GetUserStruct(s string) (User, error) {
	var temp User

	ut := Util{}
	uname := ut.GetValueForKey(s, "UserName")
	upass := ut.GetValueForKey(s, "Password")

	temp.Name = uname
	temp.Password = upass

	log.Println("Validating: user ", uname, " with password ", upass)

	return temp, nil
}

func (u *UserManager) SaveUser(t User) error {

	allUsers = append(allUsers, t)

	jsonData, err := json.MarshalIndent(allUsers, "", " ")
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile("C:\\CACHE_BASEDIR\\todo-app\\todo\\configs\\users.json", jsonData, 0644)
	if err != nil {
		log.Println("Error writing JSON file:", err)
		return err
	}

	log.Println("User Details added to the file successfully!")

	return nil
}

func (u *UserManager) ValidateUser(t User) error {
	var temp []User

	fileData, err := os.ReadFile("C:\\CACHE_BASEDIR\\todo-app\\todo\\configs\\users.json")
	if err != nil {
		log.Println("Error reading JSON file:", err)
		return err
	}

	err = json.Unmarshal(fileData, &temp)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, user := range temp {
		if user.Name == t.Name && user.Password == t.Password {
			log.Println("User Found!")
			return nil
		}
	}

	log.Println("Invalid User Credintials...")
	return errors.New("invalid username or password")
}
