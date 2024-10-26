package utils

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

func SendEmail(email string, verifycode string) error {
	message := mail.NewMsg()
	if err := message.From("walfare@gamil.com"); err != nil {
		return err
	}
	if err := message.To(email); err != nil {
		return err
	}

	message.Subject("哞福利驗證碼")
	message.SetBodyString(mail.TypeTextPlain, fmt.Sprintf("驗證碼為:%s", verifycode))

	client, err := mail.NewClient("smtp.gmail.com", mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("cherites0610@gmail.com"), mail.WithPassword("jpnxkodatrsxvjvm"))
	if err != nil {
		return err
	}

	if err := client.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
