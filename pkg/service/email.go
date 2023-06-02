package service

import (
	"bufio"
	"context"
	"encloud/config"
	"encloud/pkg/types"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	mailersend "github.com/mailersend/mailersend-go"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
	config  types.ConfYaml
}

func NewRequest(to []string, subject string, config types.ConfYaml) *Request {
	return &Request{
		to:      to,
		subject: subject,
		config:  config,
	}
}

func (r *Request) sendMail(cid string, dekType string, timestamp int64) bool {
	log.Println("config", r.config)
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	html := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\"> <html xmlns=\"http://www.w3.org/1999/xhtml\"> <head> <meta name=\"viewport\" content=\"width=device-width\" /> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" /> <title>Share Content</title> <style type=\"text/css\"> body { margin: 0 auto; padding: 0; min-width: 100%; font-family: sans-serif; } table { margin: 50px 0 50px 0; } .content { height: 100px; font-size: 18px; line-height: 30px; } </style> </head> <body> <table bgcolor=\"#FFFFFF\" width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" > <tr class=\"content\"> <td style=\"padding: 10px\"> <p>Cid: <b>" + cid + "</b></p> </td> </tr> <tr class=\"content\"> <td style=\"padding: 10px\"> <p>DEK Type: <b>" + dekType + "</b></p> </td> </tr> </table> </body> </html>"

	from := mailersend.From{
		Name:  "Encloud",
		Email: r.from,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "",
			Email: r.to[0],
		},
	}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(r.subject)
	message.SetHTML(html)

	// Open file on disk.
	f, _ := os.Open(config.Assets + "/" + fmt.Sprint(timestamp) + "_dek.txt")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	fmt.Println(encoded)

	attachment := mailersend.Attachment{Filename: "dek.txt", Content: encoded}

	message.AddAttachment(attachment)

	_, err := ms.Email.Send(ctx, message)
	if err != nil {
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
