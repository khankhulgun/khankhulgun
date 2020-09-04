package mailer


import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"github.com/khankhulgun/khankhulgun/tools"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

func (r *Request) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Request) sendMail() bool {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", utils.Config.Mail.Host, utils.Config.Mail.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", utils.Config.Mail.Username, utils.Config.Mail.Password, utils.Config.Mail.Host), utils.Config.Mail.Username, r.to, []byte(body)); err != nil {
		return false
	}
	return true
}

func (r *Request) Send(templateName string, items interface{}) bool {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendMail(); ok {
		return true
	} else {
		return false
	}
}