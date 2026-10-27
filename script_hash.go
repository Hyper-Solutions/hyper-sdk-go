package hyper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/mailru/easyjson"
)

// ErrScriptRequired is returned when a request sent only a ScriptHash and the
// API has not seen that script yet. Send the request again with the script.
var ErrScriptRequired = errors.New("api returned with: script_required")

// ScriptHash returns the hash the API identifies a challenge script by: the
// lowercase hex sha256 of the script exactly as the site served it.
func ScriptHash(script string) string {
	sum := sha256.Sum256([]byte(script))
	return hex.EncodeToString(sum[:])
}

// sendScriptRequest sends a request that carries a challenge script. Once the
// API has accepted a script from this session, the script is replaced by its
// hash (withHash builds that request). If the API no longer knows the hash it
// answers 422 script_required, and the request is sent once more with the
// script.
func sendScriptRequest[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V, script, scriptHash string, withHash func(hash string) V) (*apiResponse, error) {
	if script == "" {
		response, status, err := sendRequestStatus[V, *apiResponse](ctx, s, url, input)
		if err == nil && scriptHash != "" && isScriptRequired(status, response) {
			return nil, ErrScriptRequired
		}
		return response, err
	}

	hash := ScriptHash(script)
	if _, known := s.knownScripts.Load(hash); known {
		response, status, err := sendRequestStatus[V, *apiResponse](ctx, s, url, withHash(hash))
		if err != nil || !isScriptRequired(status, response) {
			return response, err
		}
		s.knownScripts.Delete(hash)
	}

	response, err := sendRequest[V, *apiResponse](ctx, s, url, input)
	if err == nil && response.Error == "" {
		s.knownScripts.Store(hash, struct{}{})
	}
	return response, err
}

func isScriptRequired(status int, response *apiResponse) bool {
	return status == http.StatusUnprocessableEntity && response != nil && response.Error == "script_required"
}
