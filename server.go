package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/caarlos0/env/v11"
	"github.com/things-go/go-socks5"
)

type ProxyUser struct {
	Username string `env:"USERNAME" envDefault:""`
	Password string `env:"PASSWORD" envDefault:""`
}

type ProxyParams struct {
	Credentials     []ProxyUser `envPrefix:"PROXY_CREDS"`
	Port            int         `env:"PROXY_PORT" envDefault:"1080"`
	AllowedDestFqdn string      `env:"ALLOWED_DEST_FQDN" envDefault:""`
}

func main() {
	// Working with app params
	cfg := ProxyParams{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	fmt.Printf("Proxy config: %+v\n", cfg)

	// Initialize socks5 options
	opts := []socks5.Option{
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	}

	if len(cfg.Credentials) > 0 {
		creds := make(socks5.StaticCredentials, len(cfg.Credentials))
		for _, c := range cfg.Credentials {
			creds[c.Username] = c.Password
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		authMethods := []socks5.Authenticator{cator}
		opts = append(opts, socks5.WithAuthMethods(authMethods))
	}

	if cfg.AllowedDestFqdn != "" {
		rules := PermitDestAddrPattern(cfg.AllowedDestFqdn)
		opts = append(opts, socks5.WithRule(rules))
	}

	server := socks5.NewServer(opts...)

	log.Printf("Start listening proxy service on port %d\n", cfg.Port)
	if err := server.ListenAndServe("tcp", ":"+strconv.Itoa(cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
