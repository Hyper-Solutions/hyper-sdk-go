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

// sendScriptRequest sends a request that carries a challenge script with the
// script replaced by its hash (withHash builds that request). The API answers
// 422 script_required only for a script it has never seen, and the request is
// then sent once more with the script, which the API stores for every later
// request.
func sendScriptRequest[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V, script, scriptHash string, withHash func(hash string) V) (*apiResponse, error) {
	if script == "" {
		response, status, err := sendRequestStatus[V, *apiResponse](ctx, s, url, input)
		if err == nil && scriptHash != "" && isScriptRequired(status, response) {
			return nil, ErrScriptRequired
		}
		return response, err
	}

	response, status, err := sendRequestStatus[V, *apiResponse](ctx, s, url, withHash(ScriptHash(script)))
	if err != nil || !isScriptRequired(status, response) {
		return response, err
	}
	return sendRequest[V, *apiResponse](ctx, s, url, input)
}

func isScriptRequired(status int, response *apiResponse) bool {
	return status == http.StatusUnprocessableEntity && response != nil && response.Error == "script_required"
}
