package hyper

import (
	"context"
)

// GenerateSensorData returns the sensor data required to generate valid akamai cookies using the Hyper Solutions API.
func (s *Session) GenerateSensorData(ctx context.Context, input *SensorInput) (string, error) {
	return sendRequest(ctx, s, "https://akm.justhyped.dev/sensor", input)
}

// GeneratePixelData returns the pixel data using the Hyper Solutions API.
func (s *Session) GeneratePixelData(ctx context.Context, input *PixelInput) (string, error) {
	return sendRequest(ctx, s, "https://akm.justhyped.dev/pixel", input)
}
