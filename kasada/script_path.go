package kasada

import (
	"errors"
	"io"
	"regexp"
	"strings"
)

var (
	scriptPathExpr = regexp.MustCompile(`<script\s+src="([^"]+)"`)

	ErrScriptPathNotFound = errors.New("hyper-sdk: script path not found")
)

// ParseScriptPath gets the Kasada ips.js script path from a blocked response body (status code 429)
func ParseScriptPath(reader io.Reader) (string, error) {
	src, err := io.ReadAll(reader)
	if err != nil {
		return "", errors.Join(ErrScriptPathNotFound, err)
	}

	matches := scriptPathExpr.FindSubmatch(src)
	if len(matches) < 2 {
		return "", ErrScriptPathNotFound
	}

	blockLink := string(matches[1])
	blockLink = strings.ReplaceAll(blockLink, "&amp;", "&")

	return blockLink, nil
}
