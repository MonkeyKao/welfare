package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(email string, verifycode string) error {
	// 設定發件人郵箱和密碼
	from := "cherites0610@gmail.com" // 發件人郵箱
	password := "akkyxqatzdimenml"   // 應用專用密碼

	// 設定郵件伺服器
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// 設定收件人郵箱
	to := []string{email}

	// 設定郵件內容
	subject := "Subject: 哞福利驗證碼\r\n"
	body := fmt.Sprintf("驗證碼為: %s", verifycode)
	message := []byte(subject + "\r\n" + body)

	// 驗證身份並發送郵件
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
