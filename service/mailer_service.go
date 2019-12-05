package service

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"strings"

	"github.com/andrewesteves/tapagguapi/config"
)

const (
	hostname = "smtp.gmail.com"
	username = "tapaggu@gmail.com"
	password = "T4d3m4is"
	from     = "andrew@digitalnativa.com.br"
)

// Mailer struct
type Mailer struct{}

// Send an SMTP mail
func (m Mailer) Send(to []string, template string, data []string) error {
	var encodedBody bytes.Buffer
	var msg string
	body := ""
	env, err := config.EnvConfig{}.Vars()
	if err != nil {
		panic(err.Error())
	}

	setHeader(&body, to, env.Mail.From)
	if template == "welcome" {
		body += "Subject: Boa vindas!\r\n"
		msg = setBodyWelcome()
	} else if template == "recover" {
		body += "Subject: Recuperação de senha\r\n"
		msg = setBodyRecover(data)
	}
	message := quotedprintable.NewWriter(&encodedBody)
	message.Write([]byte(msg))
	message.Close()
	body += encodedBody.String()
	setFooter(&body)

	auth := smtp.PlainAuth("", env.Mail.Username, env.Mail.Password, env.Mail.Hostname)
	err = smtp.SendMail(env.Mail.Hostname+":587", auth, env.Mail.From, to, []byte(body))
	if err != nil {
		log.Println(err.Error())
		if template == "recover" {
			return errors.New(config.LangConfig{}.I18n()["auth_reset"])
		}
		return errors.New(config.LangConfig{}.I18n()["email_error"])
	}
	return nil
}

func setHeader(body *string, to []string, from string) {
	header := make(map[string]string)
	header["From"] = from
	header["To"] = strings.Join(to, ",")
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html;charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"
	for key, value := range header {
		*body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
}

func setBodyWelcome() string {
	body := "Seja bem vindo!<br><br>"
	body += "Precisamos confirmar o seu e-mail, por favor, acesse o link abaixo.<br><br>"
	body += "[...]<br><br>"
	return body
}

func setBodyRecover(data []string) string {
	body := "Olá, tudo bem?<br><br>"
	body += "Recebemos sua solicitação de recuperação de senha.<br>"
	body += fmt.Sprintf("Acesse o link https://tapaggu.com/reset?email=%s&token=%s e atualize.<br><br>", data[0], data[1])
	return body
}

func setFooter(body *string) {
	*body += "Att,<br>"
	*body += "TaPaggu!"
}
