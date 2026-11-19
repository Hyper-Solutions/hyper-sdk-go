package datadome

import "regexp"

// A challenge page loads its bundle with the only deferred <script> tag on the
// page, so the bundle is the src of the first tag carrying a defer attribute.
// The file name is not checked: DataDome renames the bundle freely, e.g.
// interstitial.1.36.0.js or captcha.build623.01-c-v1.38.0-2026-11-17-05h07.js.
// The c.js and i.js loaders are not deferred and never match.
var (
	scriptTagRegex = regexp.MustCompile(`(?i)<script\b([^>]*)>`)
	deferAttrRegex = regexp.MustCompile(`(?i)(?:^|\s)defer(?:\s|=|/|$)`)
	srcAttrRegex   = regexp.MustCompile(`(?i)(?:^|\s)src\s*=\s*["']([^"']+)["']`)
)

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
	for _, tag := range scriptTagRegex.FindAllStringSubmatch(html, -1) {
		attrs := tag[1]
		if !deferAttrRegex.MatchString(attrs) {
			continue
		}
		if src := srcAttrRegex.FindStringSubmatch(attrs); src != nil {
			return src[1], true
		}
	}

	return "", false
}
