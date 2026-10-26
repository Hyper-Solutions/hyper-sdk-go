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

// GenerateTelemetry returns the value of the Akamai-BM-Telemetry header that the site's protected API calls carry,
// the context for the next telemetry request, and the Accept-CH client hints for the device, using the Hyper
// Solutions API. Each protected call takes a fresh header built from the previous call's context.
func (s *Session) GenerateTelemetry(ctx context.Context, input *TelemetryInput) (*TelemetryResult, error) {
	response, err := sendRequest[*TelemetryInput, *apiResponse](ctx, s, "https://akm.hypersolutions.co/v1/telemetry", input)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &TelemetryResult{Header: response.Payload, Context: response.Context, ClientHints: response.Headers}, nil
}

// GetTelemetryParams returns the path of the site's get_params request, whose response carries the "j" for
// [TelemetryInput.J], and a context that can replace the script on the first telemetry request, using the Hyper
// Solutions API.
func (s *Session) GetTelemetryParams(ctx context.Context, input *TelemetryParamsInput) (*TelemetryParamsResult, error) {
	response, err := sendRequest[*TelemetryParamsInput, *apiResponse](ctx, s, "https://akm.hypersolutions.co/v1/telemetry/params", input)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &TelemetryParamsResult{Path: response.Payload, Context: response.Context}, nil
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
