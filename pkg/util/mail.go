package util

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"bytes"
	"fmt"
	"github.com/wneessen/go-mail"
	"html/template"
	"os"
	"sync"
)

var (
	mailTemplate *template.Template
	lock         sync.Mutex
)

type data struct {
	IsReplyMail bool
	Subject     string
	Nickname    string
	URL         string
}

func send(to string, subject string, body string) error {
	m := mail.NewMsg()

	if err := m.From(setting.Config.Mail.From); err != nil {
		return err
	}

	if err := m.To(to); err != nil {
		return err
	}

	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, body)
	c, err := mail.NewClient(setting.Config.Mail.Host, mail.WithPort(setting.Config.Mail.Port), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(setting.Config.Mail.Username), mail.WithPassword(setting.Config.Mail.Password), mail.WithSSL())
	if err != nil {
		return err
	}
	return c.DialAndSend(m)
}

// getMailTemplate 获取邮件模板单例
func getMailTemplate() (*template.Template, error) {
	if mailTemplate == nil {
		lock.Lock()
		defer lock.Unlock()
		if mailTemplate == nil {
			b, err := os.ReadFile(setting.Config.Mail.Template.Path)
			if err != nil {
				return nil, err
			}
			mailTemplate, err = template.New("mail").Parse(string(b))
			if err != nil {
				return nil, err
			}
		}
	}
	return mailTemplate, nil
}

func SendReplyMail(aid uint64, nickname string) {
	defer logWithRecover()

	// template
	t, err := getMailTemplate()
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	err = t.Execute(wr, data{
		IsReplyMail: true,
		Subject:     setting.Config.Mail.Template.ReplySubject,
		Nickname:    nickname,
		URL:         fmt.Sprintf(setting.Config.Mail.Template.ReplyURLFormat, aid),
	})
	if err != nil {
		panic(err)
	}

	// send
	if err = send(setting.Config.Mail.Template.AdminEmail, setting.Config.Mail.Template.ReplySubject, wr.String()); err != nil {
		panic(err)
	}
}

func SendReviewMail(nickname string) {
	defer logWithRecover()

	// template
	t, err := getMailTemplate()
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	err = t.Execute(wr, data{
		Subject:  setting.Config.Mail.Template.ReviewSubject,
		Nickname: nickname,
		URL:      setting.Config.Mail.Template.ReviewURL,
	})
	if err != nil {
		panic(err)
	}

	// send
	if err = send(setting.Config.Mail.Template.AdminEmail, setting.Config.Mail.Template.ReviewSubject, wr.String()); err != nil {
		panic(err)
	}
}

func logWithRecover() {
	if err := recover(); err != nil {
		log.Err.Error(fmt.Sprintf("send mail fail: %s", err))
	} else {
		log.Default.Info("send mail success")
	}
}
