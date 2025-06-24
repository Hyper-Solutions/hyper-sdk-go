package hyper

import (
	"context"
	"fmt"
)

func (s *Session) TrustDecisionPayload(ctx context.Context, input *TrustDecisionPayloadInput) (string, string, string, error) {
	response, err := sendRequest[*TrustDecisionPayloadInput, *apiResponse](ctx, s, "https://trustdecision.hypersolutions.co/payload", input)
	if err != nil {
		return "", "", "", err
	}
	if response.Error != "" {
		return "", "", "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.TimeZone, response.ClientId, nil
}

func (s *Session) TrustDecisionSign(ctx context.Context, input *TrustDecisionSigningInput) (string, error) {
	response, err := sendRequest[*TrustDecisionSigningInput, *apiResponse](ctx, s, "https://trustdecision.hypersolutions.co/sign", input)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}

func (s *Session) TrustDecisionDecode(ctx context.Context, input *TrustDecisionDecodeInput) (string, error) {
	response, err := sendRequest[*TrustDecisionDecodeInput, *apiResponse](ctx, s, "https://trustdecision.hypersolutions.co/decode", input)
	if err != nil {
		return "", err
	}
	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}
