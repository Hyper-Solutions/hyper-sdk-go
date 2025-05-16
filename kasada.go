package hyper

import (
	"context"
	"encoding/base64"
	"fmt"
)

// GenerateKasadaPayload returns the payload to POST to /tl in bytes, and the generated headers
func (s *Session) GenerateKasadaPayload(ctx context.Context, input *KasadaPayloadInput) ([]byte, *KasadaHeaders, error) {
	response, err := sendRequest[*KasadaPayloadInput, *kasadaPayloadOutput](ctx, s, "https://kasada.justhyped.dev/payload", input)
	if err != nil {
		return nil, nil, err
	}

	if response.Error != "" {
		return nil, nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	decodedPayload, err := base64.StdEncoding.DecodeString(response.Payload)
	if err != nil {
		return nil, nil, err
	}

	return decodedPayload, &response.Headers, nil
}

// GenerateKasadaPow returns the x-kpsdk-cd value
func (s *Session) GenerateKasadaPow(ctx context.Context, input *KasadaPowInput) (string, error) {
	response, err := sendRequest[*KasadaPowInput, *apiResponse](ctx, s, "https://kasada.justhyped.dev/cd", input)
	if err != nil {
		return "", err
	}

	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}
