package utils

import (
	"NewApp/global"
	"NewApp/pkg/email"
	"math/rand"
	"time"
)

// MailSendCode
// 邮箱验证码发送
func MailSendCode(mail, code string) error {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	err := defailtMailer.SendMail(
		[]string{mail},
		"验证码发送",
		"你的验证码为：<h1>"+code+"</h1>",
	)
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
