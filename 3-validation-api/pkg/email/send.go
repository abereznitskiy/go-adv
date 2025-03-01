package email

import (
	"fmt"
	"go-adv/3-validation-api/pkg/res"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type SendParams struct {
	UserEmail      string
	Link           string
	SmtpUsername   string
	SmtpPassword   string
	ResponseWriter *http.ResponseWriter
}

const (
	NAME      = "Validation api"
	SUBJECT   = "Подтверждение почты"
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = "587"
)

func Send(params SendParams) error {
	e := email.NewEmail()
	e.From = NAME
	e.To = []string{params.UserEmail}
	e.Subject = SUBJECT
	e.Text = []byte(params.Link)

	smtpHost := SMTP_PORT
	smtpPort := SMTP_PORT
	smtpUsername := params.SmtpUsername
	smtpPassword := params.SmtpPassword

	err := e.Send(fmt.Sprintf("%s:%s", smtpHost, smtpPort), smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost))
	if err != nil {
		res.Json(*params.ResponseWriter, 500, err.Error())
		return err
	}

	return nil
}
