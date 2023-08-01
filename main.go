package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"

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

func run(address string) error {
	s := smtp.NewServer(&Server{
		Users: []User{{
			Name: "testuser",
			Aliases: map[string]string{
				// "alias": "root@eleonora.gay",
				"alias": "debug",
			},
		}},
	})

	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		return errors.Wrap(err, "could not load certificate")
	}

	s.EnableSMTPUTF8 = true
	s.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "",
	}

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
