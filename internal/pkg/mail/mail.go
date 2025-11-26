package mail

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body []byte) (bool, error) {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", os.Getenv("MAIL_FROM"))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)

	// ==== INI PENTING ====
	mailer.SetBody("text/html", string(body))
	// ======================

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	dialer := gomail.NewDialer(
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_FROM"),
		os.Getenv("MAIL_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return false, err
	}

	return true, nil
}

// func ForgotPassword(dto dtos.ForgotPasswordMailDto) error {
// 	_, filename, _, _ := runtime.Caller(0)

// 	// Path ke folder project
// 	basePath := filepath.Dir(filepath.Dir(filename))
// 	templatePath := filepath.Join(basePath, "templates", "reset_password.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		return err
// 	}

// 	var body bytes.Buffer
// 	if err := t.Execute(&body, dto); err != nil {
// 		return err
// 	}

// 	_, err = SendMail(dto.Email, "Forgot Password", body.Bytes())
// 	return err
// }

// func RegisteredAccount(to string, dto dtos.RegisteredAccountMailDto) error {
// 	wd, _ := os.Getwd()
// 	filePath := filepath.Join(wd, "pkg", "templates", "registered_account.html")

// 	t, err := template.ParseFiles(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	var body bytes.Buffer
// 	err = t.Execute(&body, dto)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = SendMail(to, "Registrasi Akun", nil)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func AcceptedArticle(to string, dto dtos.RegisteredAccountMailDto) error {
// 	t, err := template.ParseFiles("templates/registered_account.html")
// 	if err != nil {
// 		return err
// 	}

// 	var body bytes.Buffer
// 	err = t.Execute(&body, dto)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = SendMail(to, "Registrasi Akun", nil)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
