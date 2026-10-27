package hyper

import (
	"context"
	"encoding/base64"
	"fmt"
)

// GenerateKasadaPayload returns the payload to POST to /tl in bytes, the x-kpsdk-* headers to send with it, and the
// Accept-CH client hints for the device.
func (s *Session) GenerateKasadaPayload(ctx context.Context, input *KasadaPayloadInput) (*KasadaPayloadResult, error) {
	response, err := sendRequest[*KasadaPayloadInput, *kasadaPayloadOutput](ctx, s, "https://kasada.hypersolutions.co/payload", input)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	decodedPayload, err := base64.StdEncoding.DecodeString(response.Payload)
	if err != nil {
		return nil, err
	}

	return &KasadaPayloadResult{Payload: decodedPayload, Headers: response.Headers, ClientHints: response.ClientHints}, nil
}

// GetKasadaVersion returns the x-kpsdk-v value of a site's p.js, using the Hyper Solutions API. It is a helper and is
// not counted as usage.
func (s *Session) GetKasadaVersion(ctx context.Context, input *KasadaVersionInput) (*KasadaVersionResult, error) {
	response, err := sendRequest[*KasadaVersionInput, *apiResponse](ctx, s, "https://kasada.hypersolutions.co/version", input)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &KasadaVersionResult{Version: response.Payload}, nil
}

// GenerateKasadaPow returns the x-kpsdk-cd value
func (s *Session) GenerateKasadaPow(ctx context.Context, input *KasadaPowInput) (*KasadaPowResult, error) {
	response, err := sendRequest[*KasadaPowInput, *apiResponse](ctx, s, "https://kasada.hypersolutions.co/cd", input)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &KasadaPowResult{Payload: response.Payload}, nil
}

// GenerateBotIDHeader returns the x-is-human header value
func (s *Session) GenerateBotIDHeader(ctx context.Context, input *BotIDHeaderInput) (*BotIDHeaderResult, error) {
	response, err := sendRequest[*BotIDHeaderInput, *apiResponse](ctx, s, "https://kasada.hypersolutions.co/botid", input)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &BotIDHeaderResult{Payload: response.Payload}, nil
}
