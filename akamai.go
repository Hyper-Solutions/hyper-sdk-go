package hyper

import (
	"context"
)

// GenerateSensorData returns the sensor data required to generate valid akamai cookies using the Hyper Solutions API.
func (s *Session) GenerateSensorData(ctx context.Context, input *SensorInput) (string, error) {
	return sendRequest(ctx, s, "https://akm.justhyped.dev/sensor", input)
}

// ParseV3Dynamic returns the dynamic values for a v3 dynamic script
func (s *Session) ParseV3Dynamic(ctx context.Context, input *DynamicInput) (string, error) {
	return sendRequest(ctx, s, "https://akm.justhyped.dev/v3dynamic", input)
}

// GeneratePixelData returns the pixel data using the Hyper Solutions API.
func (s *Session) GeneratePixelData(ctx context.Context, input *PixelInput) (string, error) {
	return sendRequest(ctx, s, "https://akm.justhyped.dev/pixel", input)
}

// GenerateSbsdData returns the sbsd payload using the Hyper Solutions API.
func (s *Session) GenerateSbsdData(ctx context.Context, input *SbsdInput) (string, error) {
	return sendRequest(ctx, s, "https://akm.justhyped.dev/sbsd", input)
}
