package hyper

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
)

// GenerateKasadaPayload returns the payload to POST to /tl in bytes, and the generated headers
func (s *Session) GenerateKasadaPayload(ctx context.Context, input *KasadaPayloadInput) ([]byte, *KasadaHeaders, error) {
	if s.ApiKey == "" {
		return nil, nil, errors.New("missing api key")
	}

	payload, err := easyjson.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://kasada.justhyped.dev/payload", bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "gzip")
	req.Header.Set("x-api-key", s.ApiKey)

	if s.JwtKey != nil {
		signature, err := s.generateSignature()
		if err != nil {
			return nil, nil, err
		}
		req.Header.Set("x-signature", signature)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var response kasadaPayloadOutput
	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
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
	return sendRequest(ctx, s, "https://kasada.justhyped.dev/cd", input)
}
