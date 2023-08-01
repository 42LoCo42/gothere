package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"

	"github.com/emersion/go-smtp"
	"github.com/pkg/errors"
)

func main() {
	addr := flag.String("a", ":2525", "address to listen on")
	flag.Parse()

	if err := run(*addr); err != nil {
		log.Fatal(err)
	}
}

type Server struct{}

func (s *Server) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

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

func run(address string) error {
	s := smtp.NewServer(&Server{})

	network := "tcp"
	pretty := network + "://" + address

	l, err := net.Listen(network, address)
	if err != nil {
		return errors.Wrapf(err, "could not listen on %v", pretty)
	}

	log.Printf("listening on %v", pretty)

	if err := s.Serve(l); err != nil {
		return errors.Wrap(err, "could not run server")
	}

	return nil
}
