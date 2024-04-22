package akamai

import (
	"strconv"
	"strings"
)

// IsCookieValid determines if the provided _abck cookie value is valid, based on Akamai Bot Manager's
// client-side stop signal mechanism using the given request count. If the result is true, the client is ADVISED
// to halt further sensor data submissions. Submitting further would still produce a valid cookie but is unnecessary.
//
// The stop signal mechanism in the Akamai Bot Manager's client-side script informs a client that the cookie received is
// valid and that any additional submissions are superfluous.
//
// However, some applications do not activate the stop signal feature. In such scenarios, the client will continue
// submitting data whenever a trigger event occurs. Under these circumstances, verifying the authenticity of a cookie
// without sending it to a secured endpoint becomes challenging.
func IsCookieValid(cookie string, requestCount int) bool {
	parts := strings.Split(cookie, "~")
	if len(parts) < 2 {
		return false
	}

	requestThreshold, err := strconv.Atoi(parts[1])
	if err != nil {
		requestThreshold = -1
	}

	return requestThreshold != -1 && requestCount >= requestThreshold
}

// IsCookieInvalidated determines if the current session requires more sensors to be sent.
//
// Protected endpoints can invalidate a session by setting a new _abck cookie that ends in '~0~-1~-1' or similar.
// This function returns if such an invalidated cookie is present, if it is present you should be able to make the
// cookie valid again with only 1 sensor post.
func IsCookieInvalidated(cookie string) bool {
	parts := strings.Split(cookie, "~")
	if len(parts) < 4 {
		return false
	}

	signal, err := strconv.Atoi(parts[3])
	if err != nil {
		signal = -1
	}

	return signal > -1
}
