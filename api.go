package hyper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/Hyper-Solutions/hyper-sdk-go/v2/internal"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/buffer"
	"github.com/mailru/easyjson/jwriter"
	"net/http"
)

func sendRequest[V easyjson.Marshaler, T easyjson.Unmarshaler](ctx context.Context, s *Session, url string, input V) (response T, err error) {
	if s.ApiKey == "" {
		return response, errors.New("missing api key")
	}

	w := jwriter.Writer{
		Flags:        0,
		Error:        nil,
		Buffer:       buffer.Buffer{},
		NoEscapeHTML: true,
	}

	input.MarshalEasyJSON(&w)

	if w.Error != nil {
		return response, w.Error
	}
	payload := w.Buffer.BuildBytes()

	useCompression := false

	if len(payload) > 1000 {
		compressedBody, err := internal.CompressZstd(payload)
		if err != nil {
			return response, fmt.Errorf("failed to compress request body with zstd: %w", err)
		}
		payload = compressedBody
		useCompression = true
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return response, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept-encoding", "zstd")
	req.Header.Set("x-api-key", s.ApiKey)

	if useCompression {
		req.Header.Set("content-encoding", "zstd")
	}

	if s.JwtKey != nil {
		signature, err := generateSignature(s.ApiKey, s.JwtKey)
		if err != nil {
			return response, err
		}
		req.Header.Set("x-signature", signature)
	}
	if s.AppSecret != nil {
		signature, err := generateSignature(s.AppKey, s.AppSecret)
		if err != nil {
			return response, err
		}
		req.Header.Set("x-app-signature", signature)
		req.Header.Set("x-app-key", s.AppKey)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	respBody, err := internal.DecompressResponse(resp)
	if err != nil {
		return response, err
	}

	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	return response, nil
}
