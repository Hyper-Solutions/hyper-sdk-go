package hyper

import (
	"context"
	"errors"
	"fmt"
	"github.com/Hyper-Solutions/hyper-sdk-go/v4/internal"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

const usageURL = "https://api.hypersolutions.co/usage"

type usageOutput struct {
	UsageResult
	Error string `json:"error"`
}

// GetUsage returns the euro balance and the per-plan request quota of the session's API key, using the Hyper
// Solutions API. It is not counted as usage.
func (s *Session) GetUsage(ctx context.Context) (*UsageResult, error) {
	if s.ApiKey == "" {
		return nil, errors.New("missing api key")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, usageURL, nil)
	if err != nil {
		return nil, err
	}
	if err := s.setAuthHeaders(req); err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, release, err := internal.DecompressResponse(resp)
	defer release()
	if err != nil {
		return nil, err
	}

	var response usageOutput
	if err := jsoniter.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("decoding usage response (status %d): %w", resp.StatusCode, err)
	}

	if response.Error != "" {
		return nil, fmt.Errorf("api returned with: %s", response.Error)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned with status code: %d", resp.StatusCode)
	}

	return &response.UsageResult, nil
}
