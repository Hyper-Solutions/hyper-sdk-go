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
	"net/url"
)

// GenerateReese84Sensor returns the sensor data required to generate valid reese84 cookies using the Hyper Solutions API.
func (s *Session) GenerateReese84Sensor(ctx context.Context, site string, input *ReeseInput) (string, error) {
	return sendRequest(ctx, s, "https://incapsula.justhyped.dev/reese84/"+url.PathEscape(site), input)
}

// GenerateUtmvcCookie returns the utmvc cookie using the Hyper Solutions API.
func (s *Session) GenerateUtmvcCookie(ctx context.Context, input *UtmvcInput) (string, string, error) {
	if input.Script == "" {
		return "", "", errors.New("script must be non empty")
	}

	if len(input.SessionIds) == 0 {
		return "", "", errors.New("no session ids set")
	}

	const sensorEndpoint = "https://incapsula.justhyped.dev/utmvc"
	return sendRequestIncapsula(ctx, s, sensorEndpoint, input)
}

func sendRequestIncapsula[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V) (string, string, error) {
	if s.ApiKey == "" {
		return "", "", errors.New("missing api key")
	}

	w := jwriter.Writer{
		Flags:        0,
		Error:        nil,
		Buffer:       buffer.Buffer{},
		NoEscapeHTML: true,
	}

	input.MarshalEasyJSON(&w)

	if w.Error != nil {
		return "", "", w.Error
	}
	payload := w.Buffer.BuildBytes()

	compressed, err := internal.CompressZstd(payload)
	if err != nil {
		return "", "", err
	}
	body := bytes.NewReader(compressed)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("content-encoding", "zstd")
	req.Header.Set("accept-encoding", "zstd")
	req.Header.Set("x-api-key", s.ApiKey)

	if s.JwtKey != nil {
		signature, err := s.generateSignature()
		if err != nil {
			return "", "", err
		}
		req.Header.Set("x-signature", signature)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	respBody, err := internal.DecompressResponse(resp)
	if err != nil {
		return "", "", err
	}

	var response apiResponse
	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
		return "", "", err
	}

	if response.Error != "" {
		return "", "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, response.Swhanedl, nil
}
