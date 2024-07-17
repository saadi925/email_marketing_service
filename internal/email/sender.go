package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strings"
)

type SmtpConfig struct {
	Server   string
	Port     string
	Username string
	Email    string
	Password string
}

func NewSmtpConfig() *SmtpConfig {
	return &SmtpConfig{
		Server:   os.Getenv("SMTP_SERVER"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("EMAIL_USERNAME"),
		Email:    os.Getenv("SENDER_EMAIL"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}
}

func (config *SmtpConfig) SendEmail(toEmail, subject, templateFile string, data interface{}) error {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Server)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.Server,
	}

	client, err := smtp.Dial(config.Server + ":" + config.Port)
	if err != nil {
		return fmt.Errorf("error connecting to SMTP server: %w", err)
	}
	defer client.Close()

	if err = client.StartTLS(tlsconfig); err != nil {
		return fmt.Errorf("error starting TLS: %w", err)
	}

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("error authenticating with SMTP server: %w", err)
	}

	if err := client.Mail(config.Email); err != nil {
		return fmt.Errorf("error setting sender: %w", err)
	}

	if err := client.Rcpt(toEmail); err != nil {
		return fmt.Errorf("error setting recipient: %w", err)
	}

	dataWriter, err := client.Data()
	if err != nil {
		return fmt.Errorf("error starting email data: %w", err)
	}
	defer dataWriter.Close()

	tplBytes, err := os.ReadFile(templateFile)
	if err != nil {
		return fmt.Errorf("error reading email template: %w", err)
	}

	tpl, err := template.New("emailTemplate").Parse(string(tplBytes))
	if err != nil {
		return fmt.Errorf("error parsing email template: %w", err)
	}

	var emailBody bytes.Buffer
	err = tpl.Execute(&emailBody, data)
	if err != nil {
		return fmt.Errorf("error executing email template: %w", err)
	}

	headers := map[string]string{
		"From":         config.Email,
		"To":           toEmail,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=\"UTF-8\"",
	}

	var emailHeaders []string
	for key, value := range headers {
		emailHeaders = append(emailHeaders, fmt.Sprintf("%s: %s", key, value))
	}

	emailMessage := strings.Join(emailHeaders, "\r\n") + "\r\n\r\n" + emailBody.String()

	if _, err = dataWriter.Write([]byte(emailMessage)); err != nil {
		return fmt.Errorf("error writing email message: %w", err)
	}

	client.Quit()
	return nil
}
