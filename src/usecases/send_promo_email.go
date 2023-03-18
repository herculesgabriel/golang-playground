package usecases

import (
	"errors"
	"herculesgabriel/golang-playground/src/services/mail"
)

type SendPromoEmail struct {
	mailService mail.Mail
}

func NewSendPromoEmail(mail mail.Mail) *SendPromoEmail {
	return &SendPromoEmail{
		mailService: mail,
	}
}

func (s *SendPromoEmail) Execute(to string) error {
	if len(to) < 10 {
		return errors.New("invalid email")
	}

	from := "promo@mail.com"
	body := "<h1>Hello!</h1><p>You've just got a huge discount! :)</p>"

	s.mailService.Send(from, to, body)

	return nil
}
