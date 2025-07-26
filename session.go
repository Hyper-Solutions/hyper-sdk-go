package hyper

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// CompressionType represents the compression algorithm to use
type CompressionType string

const (
	CompressionZstd CompressionType = "zstd"
	CompressionGzip CompressionType = "gzip"
)

type Session struct {
	ApiKey      string
	JwtKey      []byte
	AppKey      string
	AppSecret   []byte
	Client      *http.Client
	Compression CompressionType
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
		ApiKey:      apiKey,
		Client:      defaultClient,
		Compression: CompressionZstd,
	}
}

// WithJwtKey adds the JWT Key to the session. If not empty, a signature will be added to each request.
func (s *Session) WithJwtKey(jwt string) *Session {
	s.JwtKey = []byte(jwt)
	return s
}

// WithOrganization adds the organization to the session.
func (s *Session) WithOrganization(key, secret string) *Session {
	s.AppKey = key
	s.AppSecret = []byte(secret)
	return s
}

// WithClient sets a new client that will be used to make requests to the Hyper Solutions API.
func (s *Session) WithClient(client *http.Client) *Session {
	s.Client = client
	return s
}

// WithCompression sets the compression type for requests.
func (s *Session) WithCompression(compression CompressionType) *Session {
	s.Compression = compression
	return s
}

func generateSignature(key string, jwtKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"key": key,
		"exp": time.Now().Add(time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
