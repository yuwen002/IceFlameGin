package utils

import "gopkg.in/gomail.v2"

// GoEmailConfig
//
// @Description: 配置信息
// @Author liuxingyu
// @Date 2024-02-15 23:30:03
type GoEmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	From         string
	To           string
	Subject      string
	Body         string
	ContentType  string
}

// SendMail
//
// @Title SendMail
// @Description: 发送邮件
// @Author liuxingyu
// @Date 2024-02-15 23:30:16
// @receiver config
// @return error
func (config *GoEmailConfig) SendMail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", config.To)
	m.SetHeader("Subject", config.Subject)

	// 如果未显式设置 ContentType，则使用默认值 "text/plain"
	contentType := config.ContentType
	if contentType == "" {
		contentType = "text/plain"
	}
	m.SetBody(contentType, config.Body)

	// Create an SMTP client
	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUsername, config.SMTPPassword)

	// Send the email
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
