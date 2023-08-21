package backend

import (
	"log"
	"net/smtp"

	configend "github.com/todo/src/config-end"
)

// Used to send mail for OTP verification.
func SendMailToDest(mailId string, data string) {
	from := configend.CurrentConfig.SenderEmail
	password := configend.CurrentConfig.SenderEmailPassword
	to := mailId

	smtpHost := configend.CurrentConfig.SMTPHost
	smtpPort := configend.CurrentConfig.SMTPPort

	subject := "HTML Email Test"
	body := `<html>
                <body>
                    <h1>The One-Time-Password for registering User in To-Do</h1>
                    <p>The six digit <strong>OTP :</strong><em>` + data + `</em>.</p>
                </body>
              </html>`

	message := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTML email sent successfully!")
}
