package hyper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/Hyper-Solutions/hyper-sdk-go/internal"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/buffer"
	"github.com/mailru/easyjson/jwriter"
	"net/http"
)

// GenerateDataDomeSlider returns the URL that will return a solved datadome cookie when blocked by captcha, and
// the extra sec-ch-* headers used on consequent requests.
func (s *Session) GenerateDataDomeSlider(ctx context.Context, input *DataDomeSliderInput) (string, *Headers, error) {
	return sendRequestDataDome(ctx, s, "https://datadome.justhyped.dev/slider", input)
}

// GenerateDataDomeInterstitial returns the form data string that is used in the POST request to receive a solved datadome cookie, and
// the extra sec-ch-* headers used on consequent requests.
func (s *Session) GenerateDataDomeInterstitial(ctx context.Context, input *DataDomeInterstitialInput) (string, *Headers, error) {
	return sendRequestDataDome(ctx, s, "https://datadome.justhyped.dev/interstitial", input)
}

// GenerateDataDomeTags returns the tags data string that is used in the POST request to receive a solved datadome cookie.
func (s *Session) GenerateDataDomeTags(ctx context.Context, input *DataDomeTagsInput) (string, error) {
	return sendRequest(ctx, s, "https://datadome.justhyped.dev/tags", input)
}

func sendRequestDataDome[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V) (string, *Headers, error) {
	if s.ApiKey == "" {
		return "", nil, errors.New("missing api key")
	}

	w := jwriter.Writer{
		Flags:        0,
		Error:        nil,
		Buffer:       buffer.Buffer{},
		NoEscapeHTML: true,
	}

	input.MarshalEasyJSON(&w)

	if w.Error != nil {
		return "", nil, w.Error
	}
	payload := w.Buffer.BuildBytes()

	compressed, err := internal.CompressZstd(payload)
	if err != nil {
		return "", nil, err
	}
	body := bytes.NewReader(compressed)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("content-encoding", "zstd")
	req.Header.Set("accept-encoding", "zstd")
	req.Header.Set("x-api-key", s.ApiKey)

	if s.JwtKey != nil {
		signature, err := s.generateSignature()
		if err != nil {
			return "", nil, err
		}
		req.Header.Set("x-signature", signature)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	respBody, err := internal.DecompressResponse(resp)
	if err != nil {
		return "", nil, err
	}

	var response apiResponse
	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
		return "", nil, err
	}

	if response.Error != "" {
		return "", nil, fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.Headers, nil
}
