package chatgo

import (
	"crypto/rsa"
	"flag"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type Config struct {
	serverAddress string
	dsn           string
	openAiApiKey  string
	jwtPrivateKey *rsa.PrivateKey
	jwtPublicKey  *rsa.PublicKey
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

	privatePEMData, err := os.ReadFile("config/jwt/private.pem")
	if err != nil {
		panic(err)
	}
	jwtPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEMData)
	if err != nil {
		panic(err)
	}
	publicPEMData, err := os.ReadFile("config/jwt/public.pem")
	if err != nil {
		panic(err)
	}
	jwtPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEMData)

	return &Config{
		serverAddress: *serverAddress,
		dsn:           *dsn,
		openAiApiKey:  *openAiApiKey,
		jwtPrivateKey: jwtPrivateKey,
		jwtPublicKey:  jwtPublicKey,
	}
}
