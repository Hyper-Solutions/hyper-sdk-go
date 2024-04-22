package hyper

import (
	"context"
	"errors"
	"net/url"
)

// GenerateReese84Sensor returns the sensor data required to generate valid reese84 cookies using the Hyper Solutions API.
func (s *Session) GenerateReese84Sensor(ctx context.Context, site, userAgent string) (string, error) {
	return sendRequest(ctx, s, "https://incapsula.justhyped.dev/reese84/"+url.PathEscape(site), &UtmvcInput{UserAgent: userAgent})
}

// GenerateUtmvcCookie returns the utmvc cookie using the Hyper Solutions API.
func (s *Session) GenerateUtmvcCookie(ctx context.Context, input *UtmvcInput) (string, error) {
	if input.Script == "" {
		return "", errors.New("script must be non empty")
	}

	if len(input.SessionIds) == 0 {
		return "", errors.New("no session ids set")
	}

	const sensorEndpoint = "https://incapsula.justhyped.dev/utmvc"
	return sendRequest(ctx, s, sensorEndpoint, input)
}
