package akamai

import (
	"errors"
	"io"
	"regexp"
)

var (
	scriptPathExpr = regexp.MustCompile(`<script type="text/javascript"\s+(?:nonce=".*")?\s+src="((?i)[a-z\d/\-_]+)"></script>`)
	sbsdScriptPathExpr = regexp.MustCompile(`<script type="text/javascript"\s+(?:nonce="[^"]*")?\s+src="((?i)[a-z\d/\-_]+)\?v=([^"'&]+)"`)
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

// ParseSbsdScriptPath get the Akamai Bot Manager web Sbsd SDK path from the given HTML code src.
// path, v-param, error
func ParseSbsdScriptPath(reader io.Reader) (string,string,error) {
	src, err := io.ReadAll(reader)
	if err != nil {
		return "", "", errors.Join(ErrScriptPathNotFound, err)
	}
	matches := regex.FindSubmatch(body)
		if len(matches) < 3 {
			return "", "", ErrScriptPathNotFound
		}
	return string(matches[1]), string(matches[2]), nil
}
