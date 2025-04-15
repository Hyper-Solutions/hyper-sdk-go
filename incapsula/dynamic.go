package incapsula

import (
	"errors"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var (
	reeseScriptRegex = regexp.MustCompile(`src\s*=\s*"((/[^/]+/\d+)(?:\?.*)?)"`)

	ErrReeseScriptNotFound = errors.New("hyper: reese script not found")
	ErrNotInterruptionPage = errors.New("hyper: not an interruption page")
	ErrInvalidURL          = errors.New("hyper: invalid URL")
)

// ParseDynamicReeseScript parses the sensor path and script path from the given HTML content.
//
// This function searches the provided HTML for a script element containing a specific pattern
// and extracts both the sensor path (shortened path) and script path (the full path).
// It requires that the HTML contains "Pardon Our Interruption" to confirm it's the correct page type.
// It also takes a URL string, extracts the hostname, and appends it to the sensor path.
// Returns the sensor path (with hostname) and script path if found, or appropriate errors otherwise.
func ParseDynamicReeseScript(html io.Reader, urlStr string) (sensorPath string, scriptPath string, err error) {
	// Parse the URL to extract hostname
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", "", ErrInvalidURL
	}
	hostname := parsedURL.Hostname()

	bytes, err := io.ReadAll(html)
	if err != nil {
		return "", "", err
	}

	content := string(bytes)

	// Verify this is an interruption page
	if !strings.Contains(content, "Pardon Our Interruption") {
		return "", "", ErrNotInterruptionPage
	}

	matches := reeseScriptRegex.FindStringSubmatch(content)
	if len(matches) < 3 {
		return "", "", ErrReeseScriptNotFound
	}

	scriptPath = matches[1]
	sensorPath = matches[2]

	// Append the hostname to the sensor path
	return sensorPath + "?d=" + hostname, scriptPath, nil
}
