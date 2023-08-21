package backend

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	configend "github.com/todo/src/config-end"
)

var allUsers []User

type UserManager struct{}

type User struct {
	Username string
	Password string
	Email    string
}

// When the program starts, It has to take all pre-existing users into memory.
// ReadUserFile will read all the users data into memory.
func (u *UserManager) ReadUsersFile() error {
	fileData, err := os.ReadFile(configend.CurrentConfig.ConfigPath + "users.json")
	if err != nil {
		log.Println("Error reading JSON file:", err)
		return err
	}

	err = json.Unmarshal(fileData, &allUsers)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Users Updated in memory!")
	return nil
}

// Used to save the registered user in users.json file
func (u *UserManager) SaveRegisteredUser(tmpUser User) error {
	allUsers = append(allUsers, tmpUser)

	jsonData, err := json.MarshalIndent(allUsers, "", " ")
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(configend.CurrentConfig.ConfigPath+"users.json", jsonData, 0644)
	if err != nil {
		log.Println("Error writing JSON file:", err)
		return err
	}

	log.Println("User Details added to the file successfully!")

	return nil
}

// Used to Validate the user when trying to login by using users.json file
func (u *UserManager) ValidateUser(t User) error {
	var temp []User

	fileData, err := os.ReadFile(configend.CurrentConfig.ConfigPath + "users.json")
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
		if user.Username == t.Username && user.Password == t.Password {
			log.Println("User Found with valid Credintials!")
			return nil
		}
	}

	log.Println("Invalid User Credintials...")
	return errors.New("invalid username or password")
}
