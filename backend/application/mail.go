package application

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/BambooTuna/letustalk/backend/config"
	"net/smtp"
	"strings"
)

type ActivateAccountMailer struct {
	UserName                   string
	Password                   string
	SmtpHost                   string
	SmtpPort                   string
	FromAddress                string
	ActivateAccountAPIEndpoint string
}

func (a ActivateAccountMailer) Send(toAddress []string, subject, message string) error {
	auth := smtp.PlainAuth("", a.UserName, a.Password, a.SmtpHost)
	return smtp.SendMail(fmt.Sprintf("%s:%s", a.SmtpHost, a.SmtpPort), auth, a.FromAddress, toAddress, []byte("Subject: "+a.encodeSubject(subject)+"\r\n"+"\r\n"+message+"\r\n"))
}

func (a ActivateAccountMailer) SendActivateCode(code, mailAddress string) error {
	message := fmt.Sprintf(
		"会員登録ありがとうございます！\n"+
			"以下のアドレスにアクセスしてアカウントを有効にしてください。(アカウントを有効にするまで予約はできません！)\n"+
			"なおこのメールに心当たりのない場合は破棄していただけると幸いです。\n\n\n"+
			"%s/%s", a.ActivateAccountAPIEndpoint, code)
	return a.Send([]string{mailAddress}, "アカウントのアクティベートをお願いします", message)
}

func (a ActivateAccountMailer) utf8Split(utf8string string, length int) []string {
	var resultString []string
	var buffer bytes.Buffer
	for k, c := range strings.Split(utf8string, "") {
		buffer.WriteString(c)
		if k%length == length-1 {
			resultString = append(resultString, buffer.String())
			buffer.Reset()
		}
	}
	if buffer.Len() > 0 {
		resultString = append(resultString, buffer.String())
	}
	return resultString
}

func (a ActivateAccountMailer) encodeSubject(subject string) string {
	var buffer bytes.Buffer
	for _, line := range a.utf8Split(subject, 13) {
		buffer.WriteString(" =?utf-8?B?")
		buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		buffer.WriteString("?=\r\n")
	}
	return buffer.String()
}

func ActivateAccountMailerFromConfig() ActivateAccountMailer {
	return ActivateAccountMailer{
		UserName:                   config.FetchEnvValue("SMTP_USER_NAME", "username"),
		Password:                   config.FetchEnvValue("SMTP_PASSWORD", "password"),
		SmtpHost:                   config.FetchEnvValue("SMTP_HOST", "127.0.0.1"),
		SmtpPort:                   config.FetchEnvValue("SMTP_PORT", "587"),
		FromAddress:                config.FetchEnvValue("SMTP_FROM_ADDRESS", "info@example.com"),
		ActivateAccountAPIEndpoint: config.FetchEnvValue("ACTIVATE_ACCOUNT_API_ENDPOINT", "http://localhost:8080/v1/account/activate"),
	}
}
