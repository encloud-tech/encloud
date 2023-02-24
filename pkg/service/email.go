package service

import (
	"bytes"
	"encloud/config"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
	config  types.ConfYaml
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func NewRequest(to []string, subject string, config types.ConfYaml) *Request {
	return &Request{
		to:      to,
		subject: subject,
		config:  config,
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

func (r *Request) BuildMail(timestamp int64) []byte {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", r.from))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(r.to, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", r.subject))

	boundary := "my-boundary-779"
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n",
		boundary))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", r.body))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=\"utf-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString("Content-Disposition: attachment; filename=dek.txt\r\n")
	buf.WriteString("Content-ID: <dek.txt>\r\n\r\n")

	data := thirdparty.ReadFile(config.Assets + "/" + fmt.Sprint(timestamp) + "_dek.txt")

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(b, data)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

	buf.WriteString("--")

	return buf.Bytes()
}

func (r *Request) sendMail(timestamp int64) bool {
	log.Println("config", r.config)
	data := r.BuildMail(timestamp)
	SMTP := fmt.Sprintf("%s:%d", r.config.Email.Server, r.config.Email.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", r.config.Email.Username, r.config.Email.Password, r.config.Email.Server), r.config.Email.From, r.to, data); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *Request) Send(templateName string, items interface{}, timestamp int64) {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendMail(timestamp); ok {
		os.Remove(config.Assets + "/" + fmt.Sprint(timestamp) + "_dek.txt")
		log.Printf("Email has been sent to %s\n", r.to)
	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
	}
}
