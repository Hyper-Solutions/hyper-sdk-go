package incapsula

import (
	"errors"
	"io"
	"math/rand"
	"regexp"
	"strconv"
)

var (
	scriptRegex = regexp.MustCompile(`src="(/_Incapsula_Resource\?[^"]*)"`)

	ErrScriptNotFound = errors.New("hyper: utmvc script not found")
)

// ParseUtmvcScriptPath parses the UTMVC script path from the given script content.
//
// This function searches the provided script content for a specific pattern matching the UTMVC script path
// using a precompiled regular expression. It extracts and returns the first match if found.
func ParseUtmvcScriptPath(script io.Reader) (string, error) {
	bytes, err := io.ReadAll(script)
	if err != nil {
		return "", err
	}

	match := scriptRegex.FindSubmatch(bytes)
	if len(match) < 2 {
		return "", ErrScriptNotFound
	}

	return string(match[1]), nil
}

// GetUtmvcSubmitPath generates a UTMVC submit path with a unique random query parameter.
//
// This function constructs a submit path for the UTMVC script by appending a random floating-point number as a query
// parameter. The random number is used to ensure the uniqueness of the request.
func GetUtmvcSubmitPath() string {
	return "/_Incapsula_Resource?SWKMTFSR=1&e=" + strconv.FormatFloat(rand.Float64(), 'g', -1, 64)
}
