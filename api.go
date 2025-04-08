package hyper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/Hyper-Solutions/hyper-sdk-go/internal"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"net/http"
)

func sendRequest[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V) (string, error) {
	if s.ApiKey == "" {
		return "", errors.New("missing api key")
	}

	payloadJSON, err := easyjson.Marshal(input)
	if err != nil {
		return "", err
	}

	requestBodyBytes := payloadJSON
	useCompression := false

	if len(payloadJSON) > 1000 {
		compressedBody, err := internal.CompressZstd(payloadJSON)
		if err != nil {
			return "", fmt.Errorf("failed to compress request body with zstd: %w", err)
		}
		requestBodyBytes = compressedBody
		useCompression = true
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(requestBodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "zstd")
	req.Header.Set("x-api-key", s.ApiKey)

	if useCompression {
		req.Header.Set("content-encoding", "zstd")
	}

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

	respBody, err := internal.DecompressResponse(resp)
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
