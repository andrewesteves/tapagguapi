package config

// LangConfig struct
type LangConfig struct{}

// I18n internationalization
func (l LangConfig) I18n() map[string]string {
	env, err := EnvConfig{}.Vars()
	if err != nil {
		panic(err.Error())
	}
	if env.Idiom.Lang == "pt_br" {
		return l.PtBr()
	}
	return l.EnUs()
}

// EnUs lang
func (l LangConfig) EnUs() map[string]string {
	lang := make(map[string]string)
	lang["success"] = "Operation successfully performed"
	lang["error"] = "Whoops, we could not perform the operation"
	lang["email_taken"] = "This e-mail is already taken"
	lang["email_error"] = "We could not send the e-mail"
	lang["email_invalid"] = "We can't find a user with that e-mail address"
	lang["auth_failed"] = "These credentials do not match our records"
	lang["auth_reset"] = "We have e-mailed your password reset link!"
	lang["token_required"] = "Please provide the authorization token"
	lang["user_inactive"] = "Please confirm your email"
	lang["confirmation_invalid"] = "Invalid user verification link"
	lang["email_send"] = "E-mail successfully sent"
	lang["welcome"] = "Registered, next step is to confirm your email"
	return lang
}

// PtBr lang
func (l LangConfig) PtBr() map[string]string {
	lang := make(map[string]string)
	lang["success"] = "Operação realizada com sucesso"
	lang["error"] = "Ops, não foi possível realizar a operação"
	lang["email_taken"] = "Este e-mail já está cadastrado"
	lang["email_error"] = "Não foi possível enviar o e-mail"
	lang["email_invalid"] = "Não conseguimos encontrar um usuário com esse endereço de e-mail"
	lang["auth_failed"] = "Essas credenciais não correspondem aos nossos registros"
	lang["auth_reset"] = "Enviamos seu link de redefinição de senha por e-mail!"
	lang["token_required"] = "Forneça o token de autorização"
	lang["user_inactive"] = "Por favor, confirme o seu e-mail"
	lang["confirmation_invalid"] = "Link inválido para confirmação de usuário"
	lang["email_send"] = "E-mail enviado com sucesso"
	lang["welcome"] = "Cadastro realizado, próximo passo é confirmar seu e-mail"
	return lang
}
