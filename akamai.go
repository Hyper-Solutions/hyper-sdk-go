package hyper

import (
	"context"
	"fmt"
)

// GenerateSensorData returns the sensor data required to generate valid akamai cookies using the Hyper Solutions API,
// the context for the next sensor request, and the Accept-CH client hints for the device.
func (s *Session) GenerateSensorData(ctx context.Context, input *SensorInput) (*SensorResult, error) {
	response, err := sendRequest[*SensorInput, *apiResponse](ctx, s, "https://akm.hypersolutions.co/v2/sensor", input)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &SensorResult{Payload: response.Payload, Context: response.Context, ClientHints: response.Headers}, nil
}

// GeneratePixelData returns the pixel data using the Hyper Solutions API.
func (s *Session) GeneratePixelData(ctx context.Context, input *PixelInput) (*PixelResult, error) {
	response, err := sendRequest[*PixelInput, *apiResponse](ctx, s, "https://akm.hypersolutions.co/pixel", input)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &PixelResult{Payload: response.Payload}, nil
}

// GenerateSbsdData returns the sbsd payload, the context for the next sbsd request, and the Accept-CH client hints
// for the device, using the Hyper Solutions API.
func (s *Session) GenerateSbsdData(ctx context.Context, input *SbsdInput) (*SbsdResult, error) {
	response, err := sendRequest[*SbsdInput, *apiResponse](ctx, s, "https://akm.hypersolutions.co/sbsd", input)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &SbsdResult{Payload: response.Payload, Context: response.Context, ClientHints: response.Headers}, nil
}
