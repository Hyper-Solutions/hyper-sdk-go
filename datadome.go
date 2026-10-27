package hyper

import (
	"context"
	"fmt"
)

// GenerateDataDomeSlider returns the URL that will return a solved datadome cookie when blocked by captcha, and
// the Accept-CH client hints for the device.
func (s *Session) GenerateDataDomeSlider(ctx context.Context, input *DataDomeSliderInput) (*DataDomeSliderResult, error) {
	response, err := sendScriptRequest(ctx, s, "https://datadome.hypersolutions.co/slider", input, input.Script, func() *DataDomeSliderInput {
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

	return &DataDomeSliderResult{Payload: response.Payload, ClientHints: response.Headers}, nil
}

// GenerateDataDomeInterstitial returns the form data string that is used in the POST request to receive a solved datadome cookie, and
// the Accept-CH client hints for the device.
func (s *Session) GenerateDataDomeInterstitial(ctx context.Context, input *DataDomeInterstitialInput) (*DataDomeInterstitialResult, error) {
	response, err := sendScriptRequest(ctx, s, "https://datadome.hypersolutions.co/interstitial", input, input.Script, func() *DataDomeInterstitialInput {
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

	return &DataDomeInterstitialResult{Payload: response.Payload, ClientHints: response.Headers}, nil
}

// GenerateDataDomeTags returns the tags data string that is used in the POST request to receive a solved datadome cookie, and
// the Accept-CH client hints for the device.
func (s *Session) GenerateDataDomeTags(ctx context.Context, input *DataDomeTagsInput) (*DataDomeTagsResult, error) {
	response, err := sendRequest[*DataDomeTagsInput, *apiResponse](ctx, s, "https://datadome.hypersolutions.co/tags", input)
	if err != nil {
		return nil, err
	}

	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return &DataDomeTagsResult{Payload: response.Payload, ClientHints: response.Headers}, nil
}
