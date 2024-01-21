package api

import "log/slog"

type Config struct {
	Port int
	Env  string
}

type API struct {
	Logger *slog.Logger
	Config Config
}
