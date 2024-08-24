package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	username := flag.String("user", "", "SIP Username")
	terminal := flag.String("terminal", "", "SIP Auth Id")
	password := flag.String("pw", "", "Password")
	server := flag.String("server", "", "SIP Server")
	target := flag.String("target", "", "SIP Target User")
	unregister := flag.Bool("unregister", false, "Unregister old terminal")

	flag.Parse()

	config := Config{
		Username:   *username,
		Terminal:   *terminal,
		Password:   *password,
		Server:     *server,
		TargetUser: *target,
		Unregister: *unregister,
	}

	SetupLogger()

	log.Info().Msg("Setting up phone")

	phone, err := SetupPhone(&config)
	defer phone.Close()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup phone")
		return
	}

	// Actually do stuff
	ctx, kill := context.WithCancel(context.Background())

	log.Info().Msgf("Calling %s from %s", config.DialURI().User, config.Terminal)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-sig
		log.Warn().Msg("Killing context (hangup)")
		kill()
	}()

	_, err = phone.Dial(ctx, config.DialURI(), config.DialAuth())

	log.Info().Msg("Exiting")
}
