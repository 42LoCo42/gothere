package main

import (
	"io"
	"log"
	"os"

	"github.com/emersion/go-smtp"
	"github.com/pkg/errors"
)

type Session struct{}

func (s *Session) AuthPlain(user, pass string) error {
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Print("from: ", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Print("rcpt: ", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return errors.Wrap(err, "could not receive data")
	}

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
