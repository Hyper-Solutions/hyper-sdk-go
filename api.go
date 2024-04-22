package hyper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
)

func sendRequest[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V) (string, error) {
	if s.ApiKey == "" {
		return "", errors.New("missing api key")
	}

	payload, err := easyjson.Marshal(input)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "gzip")
	req.Header.Set("x-api-key", s.ApiKey)

	if s.JwtKey != nil {
		signature, err := s.generateSignature()
		if err != nil {
			return "", err
		}
		req.Header.Set("x-signature", signature)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response apiResponse
	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
		return "", err
	}

	if response.Error != "" {
		return "", fmt.Errorf("api returned with: %s", response.Error)
	}

	return response.Payload, nil
}
