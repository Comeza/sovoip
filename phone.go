package main

import (
	"context"
	"time"

	"github.com/emiago/sipgo"
	"github.com/emiago/sipgox"
	"github.com/rs/zerolog/log"
)

func SetupPhone(config *Config) (*sipgox.Phone, error) {
	uaOptions := sipgo.WithUserAgent(config.Terminal)
	ua, err := sipgo.NewUA(uaOptions)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create SIP user agent")
		return nil, err
	}

	registerOptions := config.RegisterOptions(config.Unregister)
	sipUri := config.DialURI()

	log.Info().Msgf("Registering User URI: %s", sipUri.String())

	// Create new Phone Context
	phone := sipgox.NewPhone(ua)

	// Register Phone
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	err = phone.Register(ctx, sipUri, registerOptions)
	if err != nil {
		log.Warn().Err(err).Msgf("Failed to register SIP phone, Phone might be already registered %+v", err)
		return phone, nil
	}
	log.Info().Msg("Succeffully setup phone")
	return phone, nil
}
