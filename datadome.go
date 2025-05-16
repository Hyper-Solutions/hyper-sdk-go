package hyper

import (
	"context"
	"fmt"
)

// GenerateDataDomeSlider returns the URL that will return a solved datadome cookie when blocked by captcha, and
// the extra sec-ch-* headers used on consequent requests.
func (s *Session) GenerateDataDomeSlider(ctx context.Context, input *DataDomeSliderInput) (string, *Headers, error) {
	response, err := sendRequest[*DataDomeSliderInput, *apiResponse](ctx, s, "https://datadome.justhyped.dev/slider", input)
	if err != nil {
		return "", nil, err
	}

	if response.Error != "" {
		return "", nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.Headers, nil
}

// GenerateDataDomeInterstitial returns the form data string that is used in the POST request to receive a solved datadome cookie, and
// the extra sec-ch-* headers used on consequent requests.
func (s *Session) GenerateDataDomeInterstitial(ctx context.Context, input *DataDomeInterstitialInput) (string, *Headers, error) {
	response, err := sendRequest[*DataDomeInterstitialInput, *apiResponse](ctx, s, "https://datadome.justhyped.dev/interstitial", input)
	if err != nil {
		return "", nil, err
	}

	if response.Error != "" {
		return "", nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.Headers, nil
}

// GenerateDataDomeTags returns the tags data string that is used in the POST request to receive a solved datadome cookie.
func (s *Session) GenerateDataDomeTags(ctx context.Context, input *DataDomeTagsInput) (string, error) {
	response, err := sendRequest[*DataDomeTagsInput, *apiResponse](ctx, s, "https://datadome.justhyped.dev/tags", input)
	if err != nil {
		return "", err
	}

	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}
