package hyper

import (
	"context"
)

// GenerateDataDomeSlider returns the URL that will return a solved datadome cookie when blocked by captcha.
func (s *Session) GenerateDataDomeSlider(ctx context.Context, input *DataDomeSliderInput) (string, error) {
	return sendRequest(ctx, s, "https://datadome.justhyped.dev/slider", input)
}

// GenerateDataDomeInterstitial returns the form data string that is used in the POST request to receive a solved datadome cookie.
func (s *Session) GenerateDataDomeInterstitial(ctx context.Context, input *DataDomeInterstitialInput) (string, error) {
	return sendRequest(ctx, s, "https://datadome.justhyped.dev/interstitial", input)
}
