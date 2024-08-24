package main

import (
	"time"

	"github.com/emiago/sipgo/sip"
	"github.com/emiago/sipgox"
)

type Config struct {
	// The terminal of the application
	Terminal string
	// The password used for registering this terminal
	Password string
	// Username of this terminal
	Username string
	// Server where SIP is managed
	Server string
	// Terminal to call
	TargetUser string
	// Whether to unregister previously registered terminals
	Unregister bool
}

func (c *Config) DialAuth() sipgox.DialOptions {
	return sipgox.DialOptions{
		Username: c.Username,
		Password: c.Password,
	}
}

func (c *Config) AnswerOptions(onCall func(invRequest *sip.Request) int) sipgox.AnswerOptions {
	return sipgox.AnswerOptions{
		Ringtime: time.Second,
		Username: c.Username,
		Password: c.Password,
		OnCall:   onCall,
	}
}

func (c *Config) RegisterOptions(unregistger bool) sipgox.RegisterOptions {
	return sipgox.RegisterOptions{
		Username:      c.Username,
		Password:      c.Password,
		UnregisterAll: unregistger,
	}
}

func (c *Config) DialURI() sip.Uri {
	return sip.Uri{
		User: c.TargetUser,
		Host: c.Server,
	}
}
