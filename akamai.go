package hyper

import (
	"context"
	"fmt"
)

// GenerateSensorData returns the sensor data required to generate valid akamai cookies using the Hyper Solutions API.
func (s *Session) GenerateSensorData(ctx context.Context, input *SensorInput) (string, string, error) {
	response, err := sendRequest[*SensorInput, *apiResponse](ctx, s, "https://akm.justhyped.dev/v2/sensor", input)
	if err != nil {
		return "", "", err
	}
	if response.Error != "" {
		return "", "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.Context, nil
}

// ParseV3Dynamic returns the dynamic values for a v3 dynamic script
func (s *Session) ParseV3Dynamic(ctx context.Context, input *DynamicInput) (string, error) {
	response, err := sendRequest[*DynamicInput, *apiResponse](ctx, s, "https://akm.justhyped.dev/v3dynamic", input)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}

// GeneratePixelData returns the pixel data using the Hyper Solutions API.
func (s *Session) GeneratePixelData(ctx context.Context, input *PixelInput) (string, error) {
	response, err := sendRequest[*PixelInput, *apiResponse](ctx, s, "https://akm.justhyped.dev/pixel", input)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}

// GenerateSbsdData returns the sbsd payload using the Hyper Solutions API.
func (s *Session) GenerateSbsdData(ctx context.Context, input *SbsdInput) (string, error) {
	response, err := sendRequest[*SbsdInput, *apiResponse](ctx, s, "https://akm.justhyped.dev/sbsd", input)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}
