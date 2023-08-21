package configend

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type ConfigStruct struct {
	ContentStore        string `json:"ContentStorePath"`
	FrontEndPath        string `json:"FrontEndPath"`
	ImagesPath          string `json:"ImagesPath"`
	Portnum             int    `json:"PortNumber"`
	SenderEmail         string `json:"SenderMail"`
	SenderEmailPassword string `json:"SenderMailPassword"`
	ConfigPath          string `json:"ConfigFilePath"`
	SMTPHost            string `json:"SMTP-HOST"`
	SMTPPort            string `json:"SMTP-PORT"`
}

var CurrentConfig ConfigStruct

func SetConfigs() {
	CurrentConfig.ContentStore = *flag.String("contentStore", "\\build\\todo\\content-store\\", "Path to content store")
	CurrentConfig.FrontEndPath = *flag.String("frontEndPath", "\\build\\todo\\src\\front-end\\", "Path to front end")
	CurrentConfig.ImagesPath = *flag.String("imagesPath", "\\build\\todo\\images\\", "Path to images")
	CurrentConfig.Portnum = *flag.Int("portnum", 7000, "Port number")
	CurrentConfig.SenderEmail = *flag.String("senderEmail", "veluvignesh027@gmail.com", "Sender email")
	CurrentConfig.SenderEmailPassword = *flag.String("senderEmailPassword", "ofnainecktydvywf", "Sender email password")
	CurrentConfig.ConfigPath = *flag.String("configPath", "\\build\\todo\\configs\\", "Path to config file")
	CurrentConfig.SMTPHost = *flag.String("smtphost", "smtp.gmail.com", "smpt host url")
	CurrentConfig.SMTPPort = *flag.String("smtpport", "587", "smpt host port number")
	flag.Parse()

	if !checkFile(CurrentConfig.ContentStore) {
		log.Println("Error config store.Exiting()")
	}
	if !checkFile(CurrentConfig.ConfigPath) {
		log.Println("Error config config path.Exiting()")
	}
	if !checkFile(CurrentConfig.ImagesPath) {
		log.Println("Error finding config images path...")
	}
	if !checkFile(CurrentConfig.FrontEndPath) {
		log.Println("Error config front end .Exiting()")
	}
}
func checkFile(str string) bool {
	_, err := os.Stat(str)
	return err == nil
}
func PrintConfigs() {
	log.Println("Config Path : ", CurrentConfig.ConfigPath)
	log.Println("Content Store Path : ", CurrentConfig.ContentStore)
	log.Println("Front-End Path : ", CurrentConfig.FrontEndPath)
	log.Println("Images Path : ", CurrentConfig.ImagesPath)
	log.Println("Port Number : ", CurrentConfig.Portnum)
	log.Println("Sender Mail ID  : ", CurrentConfig.SenderEmail)
	log.Println("Sender Mail Password  : ", CurrentConfig.SenderEmailPassword)
	log.Println("SMTP Host URL  : ", CurrentConfig.SMTPHost)
	log.Println("Sender Port Number  : ", CurrentConfig.SMTPPort)

	log.Println("Initialized all configurations.!")

	fp, err := os.OpenFile(CurrentConfig.ConfigPath+"configurations.json", os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(-2)
	}
	defer fp.Close()

	mbyte, err := json.Marshal(CurrentConfig)
	if err != nil {
		log.Println(err)
		os.Exit(-2)
	}

	_, err = fp.Write(mbyte)
	if err != nil {
		log.Println(err)
		os.Exit(-2)
	}

	log.Println("Config Details saved in ", fp.Name())
}
