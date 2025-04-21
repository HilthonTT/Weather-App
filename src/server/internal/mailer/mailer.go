package mailer

import "embed"

const (
	FromName            = "Weather"
	maxRetries          = 3
	UserWelcomeTemplate = "email_verification.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}
