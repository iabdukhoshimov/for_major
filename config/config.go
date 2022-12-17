package config

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

// Config ...
type Config struct {
	Project struct {
		ServiceName string `env:"SERVICE_NAME,default=editory_user_service"`
		Environment string `env:"ENVIRONMENT,default=test"`
		ServiceHost string `env:"SERVICE_HOST,default=localhost"`
		RPCPort     string `env:"RPC_PORT,default=:9106"`
		LogLevel    string `env:"LOG_LEVEL,default=debug"`
	}

	Postgres struct {
		PostgresHost           string `env:"POSTGRES_HOST,default=localhost"`
		PostgresPort           int    `env:"POSTGRES_PORT,default=5432"`
		PostgresDatabase       string `env:"POSTGRES_DATABASE,default=user_service"`
		PostgresUser           string `env:"POSTGRES_USER,default=postgres"`
		PostgresPassword       string `env:"POSTGRES_PASSWORD,default=postgres"`
		PostgresMaxConnections int32  `env:"POSTGRES_MAX_CONNECTIONS,default=30"`
	}

	AccountSid  string `env:"ACCOUNT_SID,default=ACc55c86913dd36efd423bc742889e2ffa"`
	AuthToken   string `env:"AUTH_TOKEN,default=ee4e3c0602d65541d017dfbb055579cc"`
	PhoneNumber string `env:"PHONE_NUMBER,default=+19033293426"`
	SMSSender   string `env:"SMS_SENDER,default=play_mobile"`

	PlayMobile struct {
		PlayMobileLogin      string `env:"PLAY_MOBILE_LOGIN,default="`
		PlayMobilePassword   string `env:"PLAY_MOBILE_PASSWORD,default="`
		PlayMobileUrl        string `env:"PLAY_MOBILE_URL,default="`
		PlayMobileOriginator string `env:"PLAY_MOBILE_ORIGINATOR,default="`
	}
}

// Load loads environment vars and inflates Config
func Load() Config {
	var (
		cfg Config
	)

	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
