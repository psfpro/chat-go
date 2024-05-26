package client

import (
	"flag"
	"os"
)

type Config struct {
	serverAddress string
}

func NewConfig() *Config {
	serverAddress := flag.String("a", ":3200", "Server run address")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		serverAddress = &envRunAddr
	}

	return &Config{
		serverAddress: *serverAddress,
	}
}
