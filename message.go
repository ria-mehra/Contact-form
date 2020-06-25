package main

import (
  "regexp"
  "strings"

  "github.com/go-mail/mail"
)


var rxEmail = regexp.MustCompile(".+@.+\\..+")

type Message struct {
  Email   string
  Content string
  Errors  map[string]string
}

func (msg *Message) Validate() bool {
  msg.Errors = make(map[string]string)

  match := rxEmail.Match([]byte(msg.Email))
  if match == false {
    msg.Errors["Email"] = "Please enter a valid email address"
  }

  if strings.TrimSpace(msg.Content) == "" {
    msg.Errors["Content"] = "Please enter a message"
  }

  return len(msg.Errors) == 0
}

func (msg *Message) Deliver() error {
  email := mail.NewMessage()
  email.SetHeader("To", "admin@example.com")
  email.SetHeader("From", "server@example.com")
  email.SetHeader("Reply-To", msg.Email)
  email.SetHeader("Subject", "New message via Contact Form")
  email.SetBody("text/plain", msg.Content)

  username := "ab568d53cd4a61"
  password := "507d69fe7b6d6e"

  return mail.NewDialer("smtp.mailtrap.io", 25, username, password).DialAndSend(email)
}  