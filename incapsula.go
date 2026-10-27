package hyper

import (
	"context"
	"fmt"
)

// GenerateReese84Sensor returns the sensor data required to generate valid reese84 cookies using the Hyper Solutions API,
// and the Accept-CH client hints for the device.
func (s *Session) GenerateReese84Sensor(ctx context.Context, input *ReeseInput) (*Reese84Result, error) {
	response, err := sendScriptRequest(ctx, s, "https://incapsula.hypersolutions.co/reese84", input, input.Script, func() *ReeseInput {
		withoutScript := *input
		withoutScript.Script = ""
		return &withoutScript
	})
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &Reese84Result{Payload: response.Payload, ClientHints: response.Headers}, nil
}

// GenerateUtmvcCookie returns the utmvc cookie and the swhanedl value using the Hyper Solutions API.
func (s *Session) GenerateUtmvcCookie(ctx context.Context, input *UtmvcInput) (*UtmvcResult, error) {
	response, err := sendRequest[*UtmvcInput, *apiResponse](ctx, s, "https://incapsula.hypersolutions.co/utmvc", input)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &UtmvcResult{Payload: response.Payload, Swhanedl: response.Swhanedl}, nil
}
