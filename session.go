package hyper

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type Session struct {
	ApiKey string
	JwtKey []byte
	Client *http.Client
}

// Default optimized HTTP client for concurrent requests
var defaultClient = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
		IdleConnTimeout:     30 * time.Second,
	},
}

// NewSession creates a new Session that can be used to make requests to the Hyper Solutions API.
func NewSession(apiKey string) *Session {
	return &Session{
		ApiKey: apiKey,
		Client: defaultClient,
	}
}

// WithJwtKey adds the JWT Key to the session. If not empty, a signature will be added to each request.
func (s *Session) WithJwtKey(jwt string) *Session {
	s.JwtKey = []byte(jwt)
	return s
}

// WithClient sets a new client that will be used to make requests to the Hyper Solutions API.
func (s *Session) WithClient(client *http.Client) *Session {
	s.Client = client
	return s
}

func (s *Session) generateSignature() (string, error) {
	claims := jwt.MapClaims{
		"key": s.ApiKey,
		"exp": time.Now().Add(time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.JwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
