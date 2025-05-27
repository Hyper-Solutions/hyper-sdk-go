package akamai

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/mailru/easyjson"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Hyper-Solutions/hyper-sdk-go/v2/akamai/internal"
	"github.com/Hyper-Solutions/orderedobject"
	jsoniter "github.com/json-iterator/go"
)

var (
	secDurationExpr  = regexp.MustCompile(`data-duration=(\d+)`)
	secChallengeExpr = regexp.MustCompile(`challenge="(.*?)"`)
	secPageExpr      = regexp.MustCompile(`data-duration=\d+\s+src="([^"]+)"`)

	ErrSecCptParsing       = errors.New("hyper-sdk: error parsing sec-cpt")
	ErrSecCptInvalidCookie = errors.New("hyper-sdk: malformed sec_cpt cookie")
)

type SecCptChallenge struct {
	ChallengePath string

	duration      int
	challengeData *secCptChallengeData
}

//easyjson:json
type secCptChallengeData struct {
	Token      string `json:"token"`
	Timestamp  int    `json:"timestamp"`
	Nonce      string `json:"nonce"`
	Difficulty int    `json:"difficulty"`
	Count      int    `json:"count"`
	Timeout    int    `json:"timeout"`
	CPU        bool   `json:"cpu"`
	VerifyURL  string `json:"verify_url"`
}

//easyjson:json
type secCptApiResponse struct {
	SecCpChallenge     string `json:"sec-cp-challenge"`
	Provider           string `json:"provider"`
	BrandingURLContent string `json:"branding_url_content"`
	ChlgDuration       int    `json:"chlg_duration"`
	Token              string `json:"token"`
	Timestamp          int    `json:"timestamp"`
	Nonce              string `json:"nonce"`
	Difficulty         int    `json:"difficulty"`
	Timeout            int    `json:"timeout"`
	CPU                bool   `json:"cpu"`
}

// ParseSecCptChallenge parses a sec-cpt challenge from an io.Reader.
//
// The function extracts the challenge data, duration, and challenge path from the provided HTML content.
// It returns a *SecCptChallenge struct containing the parsed information and any error encountered during parsing.
//
// Example usage:
//
//	html := `<iframe id="sec-cpt-if" provider="crypto" class="crypto" challenge="..." data-key="" data-duration=5 src="/_sec/cp_challenge/ak-challenge-4-3.htm"></iframe>`
//	challenge, err := ParseSecCptChallenge(strings.NewReader(html))
//	if err != nil {
//	    // Handle the error
//	}
//
// Parameters:
//   - reader: An io.Reader containing the HTML content with the sec-cpt challenge.
//
// Returns:
//   - *SecCptChallenge: A pointer to a SecCptChallenge struct containing the parsed challenge data, duration, and challenge path.
//   - error: An error encountered during parsing, or nil if parsing was successful.
//
// Errors:
//   - ErrSecCptParsing: Returned when there is an error parsing the sec-cpt challenge data.
//   - Other errors may be returned by the underlying io.Reader or JSON unmarshaling.
func ParseSecCptChallenge(html io.Reader) (*SecCptChallenge, error) {
	src, err := io.ReadAll(html)
	if err != nil {
		return nil, errors.Join(ErrSecCptParsing, err)
	}

	challengeData, err := parseSecCptChallengeData(src)
	if err != nil {
		return nil, err
	}

	duration, err := parseSecCptDuration(src)
	if err != nil {
		return nil, err
	}

	challengePath, err := parseSecCptChallengePath(src)
	if err != nil {
		return nil, err
	}

	return &SecCptChallenge{
		challengeData: challengeData,
		duration:      duration,
		ChallengePath: challengePath,
	}, nil
}

// ParseSecCptChallengeFromJson parses a sec-cpt challenge from a JSON payload.
//
// The function takes an io.Reader containing the JSON payload of the sec-cpt challenge
// and unmarshals it into a secCptApiResponse struct. It then extracts the necessary
// information from the struct to create a SecCptChallenge struct.
//
// Example usage:
//
//	jsonPayload := `{"sec-cp-challenge":"true","provider":"crypto","branding_url_content":"/_sec/cp_challenge/crypto_message-4-3.htm","chlg_duration":30,"token":"AAQAAAAJ____9z_ZPsdHbk36hg2f6np2sGJDXmkwGmBiMBr_DDEmSWfi8Zt7BdtjWrNd9KD4DS_vim0VnK2wsa8tIC7XWsCshkvDF9J9Rf5EFwBU00c6SMXTaSNSTcDR-HVFGp3uAa67Mb3I6HeifXbjALcEomjcnwa9ZNQdDWuTAUTgNGbYw09A8AXIuP9DNv3QktUx488FV38Rm6xBXr66-MmD05hsBhucIYpLS_VCJVs9OFPnWsksPJ19ibw2K3fabfJbzIdB3Xv3J0kzLQ0gY7bpLRXK1oAcUTxNNsy-LQGe_lyV6INQ4ojPLGJpOTk","timestamp":1713283747,"nonce":"ebccdb479fcb92636fbc","difficulty":15000,"timeout":1000,"cpu":false}`
//	challenge, err := ParseSecCptChallengeFromJson(strings.NewReader(jsonPayload))
//	if err != nil {
//	    // Handle the error
//	}
//
// Parameters:
//   - payload: An io.Reader containing the JSON payload of the sec-cpt challenge.
//
// Returns:
//   - *SecCptChallenge: A pointer to a SecCptChallenge struct containing the parsed challenge data, duration, and challenge path.
//   - error: An error encountered during parsing, or nil if parsing was successful.
//
// Errors:
//   - Any error returned by the JSON unmarshaling process.
func ParseSecCptChallengeFromJson(payload io.Reader) (*SecCptChallenge, error) {
	var apiResponse secCptApiResponse
	if err := easyjson.UnmarshalFromReader(payload, &apiResponse); err != nil {
		return nil, err
	}

	return &SecCptChallenge{
		challengeData: &secCptChallengeData{
			Token:      apiResponse.Token,
			Timestamp:  apiResponse.Timestamp,
			Nonce:      apiResponse.Nonce,
			Difficulty: apiResponse.Difficulty,
			Timeout:    apiResponse.Timeout,
		},
		duration:      apiResponse.ChlgDuration,
		ChallengePath: apiResponse.BrandingURLContent,
	}, nil
}

func parseSecCptChallengeData(src []byte) (*secCptChallengeData, error) {
	challengeMatches := secChallengeExpr.FindSubmatch(src)
	if len(challengeMatches) < 2 {
		return nil, ErrSecCptParsing
	}

	decodedChallenge := make([]byte, base64.StdEncoding.DecodedLen(len(challengeMatches[1])))
	n, err := base64.StdEncoding.Decode(decodedChallenge, challengeMatches[1])
	if err != nil {
		return nil, err
	}

	var cd secCptChallengeData
	if err := easyjson.Unmarshal(decodedChallenge[:n], &cd); err != nil {
		return nil, err
	}

	return &cd, nil
}

func parseSecCptDuration(src []byte) (int, error) {
	durationMatches := secDurationExpr.FindSubmatch(src)
	if len(durationMatches) < 2 {
		return 0, ErrSecCptParsing
	}

	duration, err := strconv.Atoi(string(durationMatches[1]))
	if err != nil {
		return 0, errors.Join(ErrSecCptParsing, err)
	}

	return duration, nil
}

func parseSecCptChallengePath(src []byte) (string, error) {
	pageMatches := secPageExpr.FindSubmatch(src)
	if len(pageMatches) < 2 {
		return "", ErrSecCptParsing
	}

	return string(pageMatches[1]), nil
}

// GenerateSecCptPayload generates the payload for the sec-cpt challenge.
//
// The function takes the sec_cpt cookie value as input and extracts the necessary information
// to generate the payload. It generates the answers for the challenge using the `generateSecCptAnswers`
// function and creates an ordered object containing the token and answers.
//
// Example usage:
//
//	secCptCookie := "..."
//	payload, err := challenge.GenerateSecCptPayload(secCptCookie)
//	if err != nil {
//	    // Handle the error
//	}
//	// Use the generated payload
//	fmt.Println(string(payload))
//
// Parameters:
//   - secCptCookie: A string representing the value of the sec_cpt cookie.
//
// Returns:
//   - []byte: The generated payload as a byte slice.
//   - error: An error encountered during payload generation, or nil if generation was successful.
//
// Errors:
//   - errors.New("error parsing sec_cpt cookie"): Returned when the sec_cpt cookie is not in the expected format.
//   - Other errors may be returned by the underlying JSON marshaling.
func (s *SecCptChallenge) GenerateSecCptPayload(secCptCookie string) ([]byte, error) {
	sec, _, found := strings.Cut(secCptCookie, "~")
	if !found {
		return nil, ErrSecCptInvalidCookie
	}

	answers := generateSecCptAnswers(sec, s.challengeData)

	payload := orderedobject.NewObject[any](2)
	payload.Set("token", s.challengeData.Token)
	payload.Set("answers", answers)

	return jsoniter.Marshal(payload)
}

// Sleep sleeps for the duration specified in the sec-cpt challenge.
//
// The function uses the `duration` field of the `SecCptChallenge` struct to determine
// the number of seconds to sleep. It blocks the current goroutine for the specified duration.
//
// Example usage:
//
//	challenge, err := ParseSecCptChallenge(html)
//	if err != nil {
//	    // Handle the error
//	}
//	challenge.Sleep()
//
// Parameters:
//   - None
//
// Returns:
//   - None
func (s *SecCptChallenge) Sleep() {
	time.Sleep(time.Second * time.Duration(s.duration))
}

// SleepWithContext sleeps for the duration specified in the sec-cpt challenge or until the provided context is done.
//
// The function uses the `duration` field of the `SecCptChallenge` struct to determine
// the number of seconds to sleep. It creates a timer with the specified duration and waits for either
// the timer to expire or the provided context to be done. If the context is done before the timer expires,
// the timer is stopped to prevent it from firing.
//
// Example usage:
//
//	challenge, err := ParseSecCptChallenge(html)
//	if err != nil {
//	    // Handle the error
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	challenge.SleepWithContext(ctx)
//
// Parameters:
//   - ctx: A context.Context that can be used to cancel the sleep operation.
//
// Returns:
//   - None
func (s *SecCptChallenge) SleepWithContext(ctx context.Context) {
	timer := time.NewTimer(time.Second * time.Duration(s.duration))
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
	}
}

func generateSecCptAnswers(sec string, challengeData *secCptChallengeData) []string {
	answers := make([]string, challengeData.Count)
	challenge := sec + strconv.Itoa(challengeData.Timestamp) + challengeData.Nonce
	hash := sha256.New()
	var hashBytes [sha256.Size]byte

	for i := range answers {
		initialPart := []byte(challenge + strconv.Itoa(challengeData.Difficulty+i))

		buf := make([]byte, len(initialPart)+64)
		copy(buf, initialPart)

		for {
			answerLen := internal.FloatToStringRadix(rand.Float64(), 16, buf[len(initialPart):])
			hash.Reset()
			hash.Write(buf[:len(initialPart)+answerLen])
			hash.Sum(hashBytes[:0])

			var output int
			for _, v := range hashBytes {
				output = int(int32(uint32((output<<8)|int(v)) >> 0))
				output %= challengeData.Difficulty + i
			}

			if output != 0 {
				continue
			}

			answers[i] = string(buf[len(initialPart) : len(initialPart)+answerLen])
			break
		}
	}

	return answers
}
