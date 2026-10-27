package hyper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
)

// scriptHash is the hash the API identifies a challenge script by: the
// lowercase hex sha256 of the script exactly as the site served it.
func scriptHash(script string) string {
	sum := sha256.Sum256([]byte(script))
	return hex.EncodeToString(sum[:])
}

// withScriptHash is a request whose script is replaced by its hash: input,
// marshaled without its script, plus a "scriptHash" field. It keeps the hash
// out of the public input types, so callers always pass the script.
type withScriptHash[V easyjson.Marshaler] struct {
	input V
	hash  string
}

func (w withScriptHash[V]) MarshalEasyJSON(out *jwriter.Writer) {
	inner := jwriter.Writer{NoEscapeHTML: out.NoEscapeHTML}
	w.input.MarshalEasyJSON(&inner)
	body, err := inner.BuildBytes()
	if err != nil {
		out.Error = err
		return
	}
	// body is a JSON object; reopen it to add the hash.
	out.Raw(body[:len(body)-1], nil)
	if len(body) > 2 {
		out.RawByte(',')
	}
	out.RawString(`"scriptHash":`)
	out.String(w.hash)
	out.RawByte('}')
}

// sendScriptRequest sends a request that carries a challenge script with only
// the script's hash in place of the script (withoutScript returns input with
// the script cleared). The API answers 422 script_required only for a script
// no client has sent yet, and the request is then sent once more with the
// script, which the API stores for every later request.
func sendScriptRequest[V easyjson.Marshaler](ctx context.Context, s *Session, url string, input V, script string, withoutScript func() V) (*apiResponse, error) {
	if script == "" {
		return sendRequest[V, *apiResponse](ctx, s, url, input)
	}

	hashed := withScriptHash[V]{input: withoutScript(), hash: scriptHash(script)}
	response, status, err := sendRequestStatus[withScriptHash[V], *apiResponse](ctx, s, url, hashed)
	if err != nil || !isScriptRequired(status, response) {
		return response, err
	}
	return sendRequest[V, *apiResponse](ctx, s, url, input)
}

func isScriptRequired(status int, response *apiResponse) bool {
	return status == http.StatusUnprocessableEntity && response != nil && response.Error == "script_required"
}
