package services

import (
    "fmt"
    "os"
    "strconv"

    "github.com/go-gomail/gomail"
)

func SendEmail(to, subject, body string) error {
    port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

    m := gomail.NewMessage()
    m.SetHeader("From", os.Getenv("EMAIL_FROM"))
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer(
        os.Getenv("SMTP_HOST"),
        port,
        os.Getenv("SMTP_USER"),
        os.Getenv("SMTP_PASS"),
    )

    if err := d.DialAndSend(m); err != nil {
        return fmt.Errorf("error al enviar el correo: %w", err)
    }

    return nil
}
