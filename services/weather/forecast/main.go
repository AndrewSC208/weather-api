package main

import (
	ctr "github.com/AndrewSC208/weather-api/services/weather/forecast/controller"

	"github.com/steady-bytes/draft/pkg/chassis"
	"github.com/steady-bytes/draft/pkg/loggers/zerolog"
)

func main() {
	var (
		logger  = zerolog.New()
		handler = ctr.NewWeatherController(logger)
	)

	defer chassis.New(logger).
		WithRPCHandler(handler).
		Start()
}
