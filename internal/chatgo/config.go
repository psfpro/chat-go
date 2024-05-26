package chatgo

import (
	"flag"
	"os"
)

type Config struct {
	serverAddress string
	dsn           string
	openAiApiKey  string
}

func NewConfig() *Config {
	serverAddress := flag.String("a", ":3200", "Server run address")
	dsn := flag.String("d", "postgres://app:pass@localhost:5432/app", "DSN")
	openAiApiKey := flag.String("o", "", "OpenAI API Key")
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		serverAddress = &envRunAddr
	}
	if envDatabaseDsn := os.Getenv("DATABASE_URI"); envDatabaseDsn != "" {
		dsn = &envDatabaseDsn
	}
	if envOpenAiApiKey := os.Getenv("OPENAI_API_KEY"); envOpenAiApiKey != "" {
		openAiApiKey = &envOpenAiApiKey
	}

	return &Config{
		serverAddress: *serverAddress,
		dsn:           *dsn,
		openAiApiKey:  *openAiApiKey,
	}
}
