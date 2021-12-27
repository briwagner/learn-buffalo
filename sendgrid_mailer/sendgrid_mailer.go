package sendgrid_mailer

import (
	"errors"
	"fmt"
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/events"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func init() {
	_, err := events.Listen(func(e events.Event) {
		if e.Kind == "learnbuffalo:user:register" {
			username, err := e.Payload.Pluck("username")
			if err != nil {
				log.Print(err.Error())
				return
			}

			err = Send(fmt.Sprintf("%v", username), "Account created", fmt.Sprintf("New account created for %s", username))
			if err != nil {
				log.Printf("sendgrid error %s", err.Error())
			}
		}
	})
	if err != nil {
		log.Print(err.Error())
	}
}

// Register loads the package when the app starts.
func Register() error {
	secret := envy.Get("SENDGRID", "")
	if secret == "" {
		return errors.New("sendgrid key not set")
	}
	admin := envy.Get("ADMIN_MAIL", "")
	if admin == "" {
		return errors.New("admin email not set")
	}
	return nil
}

// Send forwards email with default from account.
func Send(rec string, subj string, msg string) error {
	secret := envy.Get("SENDGRID", "")
	if secret == "" {
		return errors.New("sendgrid key not set")
	}
	admin := envy.Get("ADMIN_MAIL", "")
	if secret == "" {
		return errors.New("sendgrid key not set")
	}
	from := mail.NewEmail("Learn Buffalo", admin)
	subject := subj
	to := mail.NewEmail("New User", rec)
	plainTextContent := msg
	htmlContent := ""
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(secret)
	response, err := client.Send(message)
	if err != nil {
		return err
	} else {
		log.Printf("Mail delivered %s", response.Body)
		return nil
	}
}
