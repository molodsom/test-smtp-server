package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter SMTP server address (e.g., mail.example.com:587): ")
	smtpServer, _ := reader.ReadString('\n')
	smtpServer = strings.TrimSpace(smtpServer)

	fmt.Print("Enter your email: ")
	from, _ := reader.ReadString('\n')
	from = strings.TrimSpace(from)

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Enter email subject: ")
	subject, _ := reader.ReadString('\n')
	subject = strings.TrimSpace(subject)

	fmt.Print("Enter email body: ")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body)

	// Read recipient email addresses
	var recipients []string
	fmt.Println("Enter recipient email addresses (double Enter to finish):")
	for {
		fmt.Print("Email: ")
		recipient, _ := reader.ReadString('\n')
		recipient = strings.TrimSpace(recipient)
		if recipient == "" {
			break
		}
		recipients = append(recipients, recipient)
	}

	// SMTP settings
	auth := smtp.PlainAuth("", from, password, strings.Split(smtpServer, ":")[0])

	// Send email to each recipient individually
	for _, recipient := range recipients {
		// Prepare email
		msg := fmt.Sprintf("From: %s\r\n", from) +
			fmt.Sprintf("To: \r\n") + // To field left empty
			fmt.Sprintf("Bcc: %s\r\n", recipient) +
			fmt.Sprintf("Subject: %s\r\n", subject) +
			"\r\n" + body

		// Send email
		err := smtp.SendMail(smtpServer, auth, from, []string{recipient}, []byte(msg))
		if err != nil {
			log.Println("Error sending email to", recipient, ":", err)
		} else {
			fmt.Println("Email successfully sent to", recipient)
		}
	}
}
