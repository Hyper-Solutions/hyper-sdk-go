package hyper

import (
	"context"
	"fmt"
)

// GenerateReese84Sensor returns the sensor data required to generate valid reese84 cookies using the Hyper Solutions API.
func (s *Session) GenerateReese84Sensor(ctx context.Context, input *ReeseInput) (string, error) {
	response, err := sendRequest[*ReeseInput, *apiResponse](ctx, s, "https://incapsula.hypersolutions.co/reese84", input)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}

// GenerateUtmvcCookie returns the utmvc cookie using the Hyper Solutions API.
func (s *Session) GenerateUtmvcCookie(ctx context.Context, input *UtmvcInput) (string, string, error) {
	response, err := sendRequest[*UtmvcInput, *apiResponse](ctx, s, "https://incapsula.hypersolutions.co/utmvc", input)
	if err != nil {
		return "", "", err
	}
	if response.Error != "" {
		return "", "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.Swhanedl, nil
}
