package akamai

import (
	"errors"
	"io"
	"regexp"
)

var (
	scriptPathExpr = regexp.MustCompile(`<script type="text/javascript"\s+(?:nonce=".*")?\s+src="((?i)[a-z\d/\-_]+)"></script>`)

	ErrScriptPathNotFound = errors.New("hyper-sdk: script path not found")
)

// ParseScriptPath gets the Akamai Bot Manager web SDK path from the given HTML code src.
func ParseScriptPath(reader io.Reader) (string, error) {
	src, err := io.ReadAll(reader)
	if err != nil {
		return "", errors.Join(ErrScriptPathNotFound, err)
	}

	matches := scriptPathExpr.FindSubmatch(src)
	if len(matches) < 2 {
		return "", ErrScriptPathNotFound
	}

	return string(matches[1]), nil
}
