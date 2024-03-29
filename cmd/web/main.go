package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/ruhollahh/go-lazy-web-components/api"
)

func main() {
	var cfg api.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Web server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	api := &api.API{
		Config: cfg,
		Logger: logger,
	}

	err := api.Serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
