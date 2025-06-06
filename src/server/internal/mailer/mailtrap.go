package mailer

import (
	"bytes"
	"errors"

	"text/template"

	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	fromEmail string
	username  string
	password  string
}

func NewMailTrapClient(username, password, fromEmail string) (mailtrapClient, error) {
	if username == "" {
		return mailtrapClient{}, errors.New("username is required")
	}

	if password == "" {
		return mailtrapClient{}, errors.New("password is required")
	}

	if fromEmail == "" {
		return mailtrapClient{}, errors.New("from email is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		username:  username,
		password:  password,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// Template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())

	message.AddAlternative("text/html", body.String())

	if isSandbox {
		dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, m.username, m.password)

		if err := dialer.DialAndSend(message); err != nil {
			return -1, err
		}
	} else {
		dialer := gomail.NewDialer("smtp.mailtrap.io", 587, m.username, m.password)

		if err := dialer.DialAndSend(message); err != nil {
			return -1, err
		}
	}

	return 200, nil
}
