package main

import "github.com/emersion/go-smtp"

type Server struct{}

func (s *Server) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}
