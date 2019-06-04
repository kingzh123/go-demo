package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	to := []string{"wangz@jfinfo.com"}
	sendEmail(to)
}

//通过 RFC 822 风格的电子邮件发送
func sendEmail(to []string)  {
	//权限验证
	auth := smtp.PlainAuth("", "190432314@qq.com", "zqoajavbeniybiji", "smtp.qq.com")
	//规范协议
	from := "kingz"
	user := "190432314@qq.com"
	subject := "a email from king"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "hello brother!"
	//邮件消息
	msg := []byte(
		"To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "<" + user + ">\r\n" +
		"Subject: " + subject + "\r\n" +
		content_type + "\r\n\r\n" +
		body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}