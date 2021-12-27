package sendgrid_mailer

import (
	"testing"
)

func Test_Send(t *testing.T) {
	err := Send("admin@learnbuffalo-notReal.com", "Test msg", "Test msg")
	if err != nil {
		t.Fatalf("mailer Send returns no error; got %s", err.Error())
	}
}
