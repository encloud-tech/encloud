package service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"

	mailersend "github.com/mailersend/mailersend-go"
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
	buf.WriteString(fmt.Sprintf("\r\n%s", "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\"><html><head><META http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style></style></head><body><u></u><div><table align=\"center\" bgcolor=\"#ffffff\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"background:#ffffff\" width=\"100%\"><tbody><tr><td><table align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"m_full\" style=\"padding:0 5px\" width=\"570\"><tbody><tr><td height=\"30\" width=\"100%\"></td></tr><tr><td align=\"center\" style=\"padding:0px;text-align:center;font-size:20px;color:#676a6c;line-height:30px;font-weight:600\" valign=\"middle\" width=\"100%\"><img alt=\"\" class=\"m_alignnone m_size-full m_wp-image-627\" src=\"https://ci5.googleusercontent.com/proxy/fSHZkZAidwvW1BU7Z9e3yUysiAqVcgBC5ms7XhEhyY0qguOto5gKzxjpV82IN9Ipep7T4YvKqQEllf0rp_tyL-rtyVq48Re8LMnIF-aCSOIuNkU=s0-d-e1-ft#https://encloud.tech/wp-content/uploads/2023/03/email_header.jpg\" width=\"600px\"></td></tr><tr><td height=\" 30\"></td></tr><tr><td style=\"padding:0px 20px;font-size:14px;color:rgb(103,106,108);line-height:24px\" valign=\"middle\" width=\"100%\"><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">Dear encloud Customer,</span></p><p class=\"m_p1\"><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><span class=\"m_s1\" style=\"font-kerning:none\">Please use <strong>CID</strong> and <strong>DEK Type</strong> and the attached DEK file to retrieve the shared file.</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"><strong>CID: </strong> "+cid+"</span><br><strong>DEK Type: </strong>"+dekType+"<br></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">Please reach out to <a href=\"mailto:contact@encloud.tech\" target=\"_blank\" rel=\"noreferrer\"><span class=\"m_s2\" style=\"font-kerning:none;color:rgb(0,0,233)\">contact@encloud.tech</span></a> for any queries on support or commercials.</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">We are aiming to provide regular updates and launching a premium offering soon so watch out for that announcement.</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">Best,</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">encloud Team</span></p></td></tr><tr><td height=\"30\" width=\"100%\"></td></tr><tr><td height=\"40\" width=\"100%\"></td></tr><tr><td align=\"center\" style=\"padding:0 20px;text-align:center;font-size:16px;color:#aaaaaa;line-height:30px;font-weight:700\" valign=\"middle\" width=\"100%\"><img alt=\"\" class=\"m_wp-image-69 m_aligncenter\" height=\"39\" src=\"https://ci4.googleusercontent.com/proxy/2XCYxeWKVbqeWOLiCNrZaz951BDVfrz7dWhOQmWnjypCDwod_Hh-hkIUVBsUqt1yfDsQckQ1FbQXn705zvvAyA431dU7I3GHlgQKQ048_8EKMr5w8w=s0-d-e1-ft#https://encloud.tech/wp-content/uploads/2023/01/EnCloud_Footer.png\" width=\"251\"></td></tr><tr><td height=\"40\" width=\"100%\"></td></tr></tbody></table></td></tr></tbody></table><br></div></body></html>"))

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

func (r *Request) sendSMTPMail(cid string, dekType string, timestamp int64) bool {
	log.Println("config", r.config)
	data := r.BuildMail(cid, dekType, timestamp)
	SMTP := fmt.Sprintf("%s:%d", r.config.Email.SMTP.Server, r.config.Email.SMTP.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", r.config.Email.SMTP.Username, r.config.Email.SMTP.Password, r.config.Email.SMTP.Server), r.config.Email.From, r.to, data); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *Request) sendMailerSendMail(cid string, dekType string, timestamp int64) bool {
	log.Println("config", r.config)
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	html := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\"><html><head><META http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style></style></head><body><u></u><div><table align=\"center\" bgcolor=\"#ffffff\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"background:#ffffff\" width=\"100%\"><tbody><tr><td><table align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"m_full\" style=\"padding:0 5px\" width=\"570\"><tbody><tr><td height=\"30\" width=\"100%\"></td></tr><tr><td align=\"center\" style=\"padding:0px;text-align:center;font-size:20px;color:#676a6c;line-height:30px;font-weight:600\" valign=\"middle\" width=\"100%\"><img alt=\"\" class=\"m_alignnone m_size-full m_wp-image-627\" src=\"https://ci5.googleusercontent.com/proxy/fSHZkZAidwvW1BU7Z9e3yUysiAqVcgBC5ms7XhEhyY0qguOto5gKzxjpV82IN9Ipep7T4YvKqQEllf0rp_tyL-rtyVq48Re8LMnIF-aCSOIuNkU=s0-d-e1-ft#https://encloud.tech/wp-content/uploads/2023/03/email_header.jpg\" width=\"600px\"></td></tr><tr><td height=\" 30\"></td></tr><tr><td style=\"padding:0px 20px;font-size:14px;color:rgb(103,106,108);line-height:24px\" valign=\"middle\" width=\"100%\"><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">Dear encloud Customer,</span></p><p class=\"m_p1\"><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><span class=\"m_s1\" style=\"font-kerning:none\">Please use <strong>CID</strong> and <strong>DEK Type</strong> and the attached DEK file to retrieve the shared file.</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"><strong>CID: </strong> " + cid + "</span><br><strong>DEK Type: </strong>" + dekType + "<br></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">Please reach out to <a href=\"mailto:contact@encloud.tech\" target=\"_blank\" rel=\"noreferrer\"><span class=\"m_s2\" style=\"font-kerning:none;color:rgb(0,0,233)\">contact@encloud.tech</span></a> for any queries on support or commercials.</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">We are aiming to provide regular updates and launching a premium offering soon so watch out for that announcement.</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\"> </span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">Best,</span></p><p class=\"m_p1\"><span class=\"m_s1\" style=\"font-kerning:none\">encloud Team</span></p></td></tr><tr><td height=\"30\" width=\"100%\"></td></tr><tr><td height=\"40\" width=\"100%\"></td></tr><tr><td align=\"center\" style=\"padding:0 20px;text-align:center;font-size:16px;color:#aaaaaa;line-height:30px;font-weight:700\" valign=\"middle\" width=\"100%\"><img alt=\"\" class=\"m_wp-image-69 m_aligncenter\" height=\"39\" src=\"https://ci4.googleusercontent.com/proxy/2XCYxeWKVbqeWOLiCNrZaz951BDVfrz7dWhOQmWnjypCDwod_Hh-hkIUVBsUqt1yfDsQckQ1FbQXn705zvvAyA431dU7I3GHlgQKQ048_8EKMr5w8w=s0-d-e1-ft#https://encloud.tech/wp-content/uploads/2023/01/EnCloud_Footer.png\" width=\"251\"></td></tr><tr><td height=\"40\" width=\"100%\"></td></tr></tbody></table></td></tr></tbody></table><br></div></body></html>"

	from := mailersend.From{
		Name:  "Encloud",
		Email: r.config.Email.From,
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
		log.Print(err)
		return false
	}

	return true
}

func (r *Request) Send(cid string, dekType string, timestamp int64, emailType string) bool {
	if emailType == "smtp" {
		return r.sendSMTPMail(cid, dekType, timestamp)
	} else {
		return r.sendMailerSendMail(cid, dekType, timestamp)
	}
}
