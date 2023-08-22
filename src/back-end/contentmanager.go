package backend

import (
	"encoding/json"
	"log"
	"os"

	configend "github.com/todo/src/config-end"
)

type ContentManager struct{}

type ContentTasks struct {
	TaskName string
	Date     string
	Time     string
}

// When the program starts, It has to take all pre-existing users into memory.
// ReadUserFile will read all the users data into memory.
func (c *ContentManager) ReadTaskFile(user string) ([]ContentTasks, error) {
	allContent := []ContentTasks{}

	fileData, err := os.ReadFile(configend.CurrentConfig.ContentStore + user + "-tasks.json")
	if err != nil {
		log.Println("Error reading JSON file:", err)
		return allContent, err
	}
	if len(fileData) != 0 {
		err = json.Unmarshal(fileData, &allContent)
		if err != nil {
			log.Println(err)
			return allContent, err
		}
	} else {
		log.Println("Empty File Cannot Read")
		return allContent, nil
	}

	log.Println("Tasks Updated in memory!")
	return allContent, nil
}

// Used to save the registered user in users.json file
func (c *ContentManager) SaveTasksInFile(tmpTask ContentTasks, user string) error {

	allContent, err := c.ReadTaskFile(user)
	if err != nil {
		log.Println(err)
		return err
	}

	allContent = append(allContent, tmpTask)

	jsonData, err := json.MarshalIndent(allContent, "", " ")
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(configend.CurrentConfig.ContentStore+user+"-tasks.json", jsonData, 0644)
	if err != nil {
		log.Println("Error writing JSON file:", err)
		return err
	}

	log.Println("Task Added to the file successfully!")
	return nil
}
