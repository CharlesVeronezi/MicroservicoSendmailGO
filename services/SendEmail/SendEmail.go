package sendemail

import (
	"fmt"
	"net/smtp"
)

func EnviarEmail(conteudo string) error {
	// Dados do remetente.
	from := "remetente@gmail.com"
	password := "senha-de-app"

	// Endereço de email do destinatário.
	to := []string{
		"distanatario@gmail.com",
	}

	// Configuração do servidor SMTP.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Mensagem.
	message := []byte(conteudo)

	// Autenticação.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Enviando email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		return err // Retorna o erro, se houver.
	}

	fmt.Println("Email enviado")
	return nil // Retorna nil se não houver erro.
}
