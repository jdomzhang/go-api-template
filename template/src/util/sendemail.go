package util

import (
	"bytes"
	"crypto/tls"
	"net"
	"net/smtp"
	"text/template"

	"../../config"
)

const (
	TEMPLATE_SUBJECT = `[APPT: {{.ID}}] New Demo Appointment 有新的试听预约`
	TEMPLATE_BODY    = `
	<html>
	<head />
	<body>
		<p>Hello,</p>
		<p>There is a new demo appointment - [ID: {{.ID}}]<br/>
			<ul>
				<li>Student Age: {{.Age}}</li>
				<li>Mobile: {{.Mobile}}</li>
			</ul>
		</p>
		<p>From Wechat mini-program</p>
		<hr/>
		<p>您好：</p>
		<p>有一个新的试听预约 - [编号：{{.ID}}]<br/>
			<ul>
				<li>幼儿年龄： {{.Age}}</li>
				<li>手机号：{{.Mobile}}</li>
			</ul>
		</p>
		<p>来自微信小程序</p>
	</body>
	</html>
	`
)

type EmailTemplateResult struct {
	Subject string
	Body    string
}

// SendEmail sends email
func SendEmail(to string, data interface{}) error {
	return SendEmails([]string{to}, data)
}

// SendEmails send email to multiple recipients
func SendEmails(to []string, data interface{}) error {
	ssl := config.All["smtp.ssl"] == "true"
	serverAddress := config.All["smtp.serveraddress"]
	// authServer := config.All["smtp.auth.server"]
	authServer, _, _ := net.SplitHostPort(serverAddress)
	authUser := config.All["smtp.auth.user"]
	authPassword := config.All["smtp.auth.password"]
	from := config.All["smtp.sendfrom"]

	// Set up authentication information.
	auth := smtp.PlainAuth("", authUser, authPassword, authServer)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	var msg bytes.Buffer
	msg.WriteString("From: ")
	msg.WriteString(from)
	msg.WriteString("\n")
	for _, v := range to {
		msg.WriteString("To: ")
		msg.WriteString(v)
		msg.WriteString("\n")
	}

	var (
		err   error
		email EmailTemplateResult
	)

	if email, err = parseTemplate(data); err != nil {
		return err
	}

	msg.Write(email.getHTMLMsg())
	if ssl {
		err = sendMailSSL(serverAddress, auth, from, to, msg.Bytes())
	} else {
		err = smtp.SendMail(serverAddress, auth, from, to, msg.Bytes())
	}

	return err
}

func sendMailSSL(serverAddress string, auth smtp.Auth, from string, to []string, message []byte) error {
	host, _, _ := net.SplitHostPort(serverAddress)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", serverAddress, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from); err != nil {
		return err
	}

	for _, v := range to {
		if err = c.Rcpt(v); err != nil {
			return err
		}
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(message)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}

func parseTemplate(data interface{}) (EmailTemplateResult, error) {
	subject := new(bytes.Buffer)
	body := new(bytes.Buffer)
	tmplSubject, _ := template.New("subject").Parse(TEMPLATE_SUBJECT)
	tmplBody, _ := template.New("body").Parse(TEMPLATE_BODY)

	result := EmailTemplateResult{}

	if err := tmplSubject.Execute(subject, data); err != nil {
		return result, err
	}

	if err := tmplBody.Execute(body, data); err != nil {
		return result, err
	}

	result.Subject = subject.String()
	result.Body = body.String()

	return result, nil
}

func (email EmailTemplateResult) getHTMLMsg() []byte {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + email.Subject + "\n"
	msg := []byte(subject + mime + "\n" + email.Body + "\n")

	return msg
}
