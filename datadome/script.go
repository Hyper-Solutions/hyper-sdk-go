package datadome

import "regexp"

// challengeScriptRegex matches the tag a challenge page uses to load the
// bundle from ct.captcha-delivery.com, capturing the URL. The name in front of
// the version is left open so it keeps matching whatever DataDome calls the
// bundle, e.g. interstitial.1.33.0.202609141.js or captcha.1.34.0.202609141.js.
var challengeScriptRegex = regexp.MustCompile(`<script[^>]+src\s*=\s*["']([^"']*/[a-z_-]+\.\d+\.\d+\.\d+\.[^"']*\.js)["']`)

// ParseChallengeScriptURL reports the challenge bundle a device check or
// captcha page loads with a <script defer src="..."> tag, and whether the page
// uses one at all.
//
// DataDome serves some challenge pages with the bundle inlined in the HTML and
// others with it in its own file, switching between the two per request, so
// this has to be checked on every challenge rather than configured once.
//
// When it reports true, GET the returned URL with the same client, proxy and
// headers you used for the challenge page, and pass the response body as the
// Script field of DataDomeInterstitialInput or DataDomeSliderInput. When it
// reports false the bundle is already in the HTML and Script stays empty.
//
//	deviceLink, err := datadome.ParseInterstitialDeviceCheckLink(body, cookie, referer)
//	// ... GET deviceLink, read the body into html ...
//
//	var script string
//	if scriptURL, ok := datadome.ParseChallengeScriptURL(html); ok {
//	    // fetch scriptURL with your own client, read the body into script
//	}
//
//	payload, headers, err := session.GenerateDataDomeInterstitial(ctx, &hyper.DataDomeInterstitialInput{
//	    UserAgent:      userAgent,
//	    DeviceLink:     deviceLink,
//	    Html:           html,
//	    Script:         script,
//	    AcceptLanguage: acceptLanguage,
//	    IP:             ip,
//	})
func ParseChallengeScriptURL(html string) (string, bool) {
	matches := challengeScriptRegex.FindStringSubmatch(html)
	if len(matches) != 2 {
		return "", false
	}

	return matches[1], true
}
