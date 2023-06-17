package main

import (
	"flag"
	"time"

	"github.com/influxdata/telegraf/plugins/common/shim"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/zachstence/winact-telegraf-plugin/plugins/inputs/winact"
)

var pollInterval = flag.Duration("poll_interval", 1*time.Second, "how often to send metrics")
var pollIntervalDisabled = flag.Bool("poll_interval_disabled", false, "how often to send metrics")
var configFile = flag.String("config", "", "path to the config file for this plugin")
var debug = flag.Bool("debug", false, "sets log level to debug")
var err error

func main() {
	flag.Parse()
	if *pollIntervalDisabled {
		*pollInterval = shim.PollIntervalDisabled
	}

	// Set log level based on -debug flag
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug logging enabled")

	shim := shim.New()

	if configFile != nil {
		err = shim.LoadConfig(configFile)
		if err != nil {
			log.Fatal().Err(err).Msg("error loading config")
		}
	}

	if err := shim.Run(*pollInterval); err != nil {
		log.Fatal().Err(err).Msg("winact input plugin failed")
	}
}
