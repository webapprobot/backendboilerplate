package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/webappbot/backendboilerplate/config"
	"gopkg.in/gomail.v2"
)

func SendMail(subject string, to string, fromName string, html string) {
	env := os.Getenv("env")
	mailerHost, _ := config.GetConfig("mailer", env, "host")
	mailerPort, _ := config.GetConfig("mailer", env, "port")
	// mailerSecure, _ := config.GetConfig("mailer", env, "secure")
	mailerUser, _ := config.GetConfig("mailer", env, "auth", "user")
	mailerUserPassword, _ := config.GetConfig("mailer", env, "auth", "pass")
	from := fmt.Sprintf("%v", mailerUser)

	m := gomail.NewMessage()
	m.SetAddressHeader("From", from, fromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	var port interface{}
	port, _ = strconv.ParseInt(fmt.Sprintf("%v", mailerPort), 10, 32)
	d := gomail.NewPlainDialer(fmt.Sprintf("%v", mailerHost), int(port.(int64)), fmt.Sprintf("%v", mailerUser), fmt.Sprintf("%v", mailerUserPassword))
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
