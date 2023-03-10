package service

import (
	"bytes"
	"encloud/config"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"encoding/base64"
	"fmt"
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

func (r *Request) BuildMail(cid string, dekType string, timestamp int64) []byte {
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
	buf.WriteString(fmt.Sprintf("\r\n%s", "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\"> <html xmlns=\"http://www.w3.org/1999/xhtml\"> <head> <meta name=\"viewport\" content=\"width=device-width\" /> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" /> <title>Share Content</title> <style type=\"text/css\"> body { margin: 0 auto; padding: 0; min-width: 100%; font-family: sans-serif; } table { margin: 50px 0 50px 0; } .content { height: 100px; font-size: 18px; line-height: 30px; } </style> </head> <body> <table bgcolor=\"#FFFFFF\" width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" > <tr class=\"content\"> <td style=\"padding: 10px\"> <p>Cid: <b>"+cid+"</b></p> </td> </tr> <tr class=\"content\"> <td style=\"padding: 10px\"> <p>DEK Type: <b>"+dekType+"</b></p> </td> </tr> </table> </body> </html>"))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=\"utf-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString("Content-Disposition: attachment; filename=dek.txt\r\n")
	buf.WriteString("Content-ID: <dek.txt>\r\n\r\n")

	data, err := thirdparty.ReadFile(config.Assets + "/" + fmt.Sprint(timestamp) + "_dek.txt")
	if err != nil {
		log.Println(err)
	}

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(b, data)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

	buf.WriteString("--")

	return buf.Bytes()
}

func (r *Request) sendMail(cid string, dekType string, timestamp int64) bool {
	log.Println("config", r.config)
	data := r.BuildMail(cid, dekType, timestamp)
	SMTP := fmt.Sprintf("%s:%d", r.config.Email.Server, r.config.Email.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", r.config.Email.Username, r.config.Email.Password, r.config.Email.Server), r.config.Email.From, r.to, data); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *Request) Send(cid string, dekType string, timestamp int64) {
	if ok := r.sendMail(cid, dekType, timestamp); ok {
		os.Remove(config.Assets + "/" + fmt.Sprint(timestamp) + "_dek.txt")
		log.Printf("Email has been sent to %s\n", r.to)
	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
	}
}
