package datadome

import (
	"bytes"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/justhyped/OrderedForm"
	"io"
	"regexp"
	"strconv"
)

var (
	ddRegex          = regexp.MustCompile(`var\s+dd\s*=\s*\{\s*([\s\S]*?)\s*}`)
	singleQuoteRegex = regexp.MustCompile(`'([^']*)'`)
	keyRegex         = regexp.MustCompile(`([^"]|^)(\b\w+)\s*: `)
)

// ParseInterstitialDeviceCheckLink parses the device check url (/interstitial/?initialCid...) from a blocked response body
//
// the datadome cookie is the current value of the 'datadome' cookie
func ParseInterstitialDeviceCheckLink(body io.Reader, datadomeCookie, referer string) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	matches := ddRegex.FindSubmatch(bodyBytes)
	if matches == nil || len(matches) < 2 {
		return "", errors.New("DD object not found")
	}

	jsonObject := append([]byte("{"), bytes.TrimSpace(matches[1])...)
	jsonObject = append(jsonObject, []byte("}")...)
	jsonObject = singleQuoteRegex.ReplaceAll(jsonObject, []byte(`"$1"`))
	jsonObject = keyRegex.ReplaceAll(jsonObject, []byte(`$1"$2":`))

	var d dd
	err = jsoniter.Unmarshal(jsonObject, &d)
	if err != nil {
		return "", err
	}

	form := new(OrderedForm.OrderedForm)
	form.Set("initialCid", d.Cid)
	form.Set("hash", d.Hsh)
	form.Set("cid", datadomeCookie)
	form.Set("referer", referer)
	form.Set("s", strconv.FormatInt(d.S, 10))
	form.Set("b", strconv.FormatInt(d.B, 10))
	form.Set("dm", "cd")

	return "https://geo.captcha-delivery.com/interstitial/?" + form.URLEncode(), nil
}

// ParseSliderDeviceCheckLink parses the device check url (/captcha/?initialCid...) from a blocked response body
//
// the datadome cookie is the current value of the 'datadome' cookie
func ParseSliderDeviceCheckLink(body io.Reader, datadomeCookie, referer string) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	matches := ddRegex.FindSubmatch(bodyBytes)
	if matches == nil || len(matches) < 2 {
		return "", errors.New("DD object not found")
	}

	jsonObject := append([]byte("{"), bytes.TrimSpace(matches[1])...)
	jsonObject = append(jsonObject, []byte("}")...)

	jsonObject = singleQuoteRegex.ReplaceAll(jsonObject, []byte(`"$1"`))
	jsonObject = keyRegex.ReplaceAll(jsonObject, []byte(`$1"$2":`))

	var d dd
	err = jsoniter.Unmarshal(jsonObject, &d)
	if err != nil {
		return "", err
	}

	if d.T == "bv" {
		return "", errors.New("proxy blocked")
	}

	form := new(OrderedForm.OrderedForm)
	form.Set("initialCid", d.Cid)
	form.Set("hash", d.Hsh)
	form.Set("cid", datadomeCookie)
	form.Set("t", d.T)
	form.Set("referer", referer)
	form.Set("s", strconv.FormatInt(d.S, 10))
	form.Set("e", d.E)
	form.Set("dm", "cd")

	return "https://geo.captcha-delivery.com/captcha/?" + form.URLEncode(), nil
}

type dd struct {
	Rt  string `json:"rt"`
	Cid string `json:"cid"`
	Hsh string `json:"hsh"`
	B   int64  `json:"b"`
	S   int64  `json:"s"`
	E   string `json:"e"`
	T   string `json:"t"`
}
