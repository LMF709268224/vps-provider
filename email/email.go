package email

import (
	"fmt"
	"net/smtp"
	"net/textproto"

	"vps-provider/config"

	"github.com/jordan-wright/email"
)

type emailData struct {
	SendTo  string
	Subject string
	Tittle  string
	Content string
}

func sendEmail(data emailData) error {
	cfg := config.Cfg.Email
	message := &email.Email{
		To:      []string{data.SendTo},
		From:    fmt.Sprintf("%s <%s>", cfg.Name, cfg.Username),
		Subject: data.Subject,
		Text:    []byte(data.Tittle),
		HTML:    []byte(data.Content),
		Headers: textproto.MIMEHeader{},
	}

	// smtp.PlainAuth：the first param can be empty，the second param should be the email account，the third param is the secret of the email
	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SMTPHost)

	return message.Send(addr, auth)
}
