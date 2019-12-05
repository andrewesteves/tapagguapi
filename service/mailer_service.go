package service

import (
	"fmt"
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
func (m Mailer) Send(to []string, template string) error {
	body := ""
	env, err := config.Env{}.Vars()
	if err != nil {
		panic(err.Error())
	}
	setHeader(&body, to, env.Mail.From)
	if template == "welcome" {
		setBodyWelcome(&body)
	} else if template == "recover" {
		setBodyRecover(&body)
	}
	setFooter(&body)
	auth := smtp.PlainAuth("", env.Mail.Username, env.Mail.Password, env.Mail.Hostname)
	err = smtp.SendMail(env.Mail.Hostname+":587", auth, env.Mail.From, to, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func setHeader(body *string, to []string, from string) {
	header := make(map[string]string)
	header["From"] = from
	header["To"] = strings.Join(to, ",")
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain;charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"
	for key, value := range header {
		*body += fmt.Sprintf("%s: %s\r\n", key, value)
	}
}

func setBodyWelcome(body *string) {
	*body += "Subject: Boa vindas!\r\n"
	*body += "Seja bem vindo!\r\n\r\n"
	*body += "Precisamos confirmar o seu e-mail, por favor, acesse o link abaixo.\r\n"
	*body += "[...]\r\n\r\n"
}

func setBodyRecover(body *string) {
	*body += "Subject: Recuperação de senha\r\n"
	*body += "Olá, tudo bem?\r\n\r\n"
	*body += "Recebemos sua solicitação de recuperação de senha.\r\n"
	*body += "Acesse o link [...] e atualize.\r\n\r\n"
}

func setFooter(body *string) {
	*body += "Att,\r\n"
	*body += "TaPaggu!"
}
