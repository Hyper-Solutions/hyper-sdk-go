package hyper

import (
	"context"
	"encoding/base64"
	"fmt"
)

// GenerateKasadaPayload returns the payload to POST to /tl in bytes, and the generated headers
func (s *Session) GenerateKasadaPayload(ctx context.Context, input *KasadaPayloadInput) ([]byte, *KasadaHeaders, error) {
	response, err := sendRequest[*KasadaPayloadInput, *kasadaPayloadOutput](ctx, s, "https://kasada.hypersolutions.co/payload", input)
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
	response, err := sendRequest[*KasadaPowInput, *apiResponse](ctx, s, "https://kasada.hypersolutions.co/cd", input)
	if err != nil {
		return "", err
	}

	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}

// GenerateBotIDHeader returns the x-is-human header value
func (s *Session) GenerateBotIDHeader(ctx context.Context, input *BotIDHeaderInput) (string, error) {
	response, err := sendRequest[*BotIDHeaderInput, *apiResponse](ctx, s, "https://kasada.hypersolutions.co/botid", input)
	if err != nil {
		return "", err
	}

	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}
