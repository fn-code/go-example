package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hello there")
}

func send(body string) {
	from := "funproject@funproject.space"
	pass := "4dm1nfnp.space"
	to := "ludyyn@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("mail.funproject.space:587",
		smtp.PlainAuth("", from, pass, "mail.funproject.space"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %v", err)
		return
	}

	log.Print("sent, success")
}
