package utils

import (
	"os"
)

type Env struct {
	// Auth0
	Auth0Domain       string
	Auth0Audience     string
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0CallbackUrl  string

	// JWT
	JwtSecret   string
	JwtAudience string
	JwtIss      string

	// Signal
	SignalServerHost string
	SignalServerPort string

	// Origin Server
	ServerHost string
	ServerPort string
}

func NewEnv() *Env {
	domain := os.Getenv("AUTH0_DOMAIN")
	audience := os.Getenv("AUTH0_AUDIENCE")
	clientid := os.Getenv("AUTH0_CLIENT_ID")
	clientsecret := os.Getenv("AUTH0_CLIENT_SECRET")
	callbackurl := os.Getenv("AUTH0_CALLBACK_URL")

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtAudience := os.Getenv("JWT_AUDIENCE")
	jwtIss := os.Getenv("JWT_ISS")

	signalServerHost := os.Getenv("SIGNAL_SERVER_HOST")
	signalServerPort := os.Getenv("SIGNAL_SERVER_PORT")

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	return &Env{
		Auth0Domain:       domain,
		Auth0Audience:     audience,
		Auth0ClientID:     clientid,
		Auth0ClientSecret: clientsecret,
		Auth0CallbackUrl:  callbackurl,

		JwtSecret:   jwtSecret,
		JwtAudience: jwtAudience,
		JwtIss:      jwtIss,

		SignalServerHost: signalServerHost,
		SignalServerPort: signalServerPort,

		ServerHost: serverHost,
		ServerPort: serverPort,
	}
}
