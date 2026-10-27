package hyper

type UtmvcInput struct {
	UserAgent string `json:"userAgent"`

	// SessionIds The value of each cookie that has a name that starts with incap_ses_
	SessionIds []string `json:"sessionIds"`

	// Script is the UTMVC script contents
	Script string `json:"script"`
}
type ReeseInput struct {
	UserAgent      string `json:"userAgent"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
	ScriptUrl      string `json:"scriptUrl"`
	PageUrl        string `json:"pageUrl"`
	Pow            string `json:"pow"`
	Script         string `json:"script"`
}

type SensorInput struct {
	// Abck is the _abck cookie retrieved from your cookiejar.
	// Make sure you always retrieve it fresh from the cookiejar as it gets updated while making requests.
	Abck string `json:"abck"`

	// Bmsz is the bm_sz cookie retrieved from your cookiejar, make sure you always retrieve it fresh from the cookiejar.
	Bmsz string `json:"bmsz"`

	// Version is the akamai version, this will usually be "2"
	Version string `json:"version"`

	// PageUrl is the page url that loaded the akamai script, it is the same URL as the referer header on the sensor posts
	PageUrl string `json:"pageUrl"`

	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// Script is mutually exclusive with [SensorInput.Context], the first sensor request should include the script field.
	// 	Subsequent request should only include the Context.
	Script string `json:"script"`

	// ScriptUrl is the full URL where you are posting sensor data to
	ScriptUrl string `json:"scriptUrl"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
	Context        string `json:"context"`
}

// TelemetryInput is the input of [Session.GenerateTelemetry]: a sensor request for a site that sends an
// Akamai-BM-Telemetry header on its protected API calls instead of posting sensor data.
//
// Only v3 dynamic scripts send the header, so [SensorInput.Version] may be left empty. Script and Context chain
// across headers the same way they do across sensor posts.
type TelemetryInput struct {
	SensorInput

	// J is the "j" from the response of the site's get_params request (see [Session.GetTelemetryParams]), sent
	// as it was returned. Leave it empty when the site makes no get_params request.
	J string `json:"j"`
}

// TelemetryParamsInput is the input of [Session.GetTelemetryParams].
type TelemetryParamsInput struct {
	// PageUrl is the page url that loaded the akamai script.
	PageUrl string `json:"pageUrl"`

	// Script is the akamai script, mutually exclusive with [TelemetryParamsInput.Context].
	Script string `json:"script"`

	// Context is a context returned by an earlier telemetry or telemetry params request, mutually exclusive with
	// [TelemetryParamsInput.Script].
	Context string `json:"context"`
}

type PixelInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent      string `json:"userAgent"`
	HTMLVar        string `json:"htmlVar"`
	ScriptVar      string `json:"scriptVar"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type SbsdInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`
	Uuid      string `json:"uuid"`
	PageUrl   string `json:"pageUrl"`
	OCookie   string `json:"o"`

	// Script is mutually exclusive with [SbsdInput.Context], the first sbsd request should include the script field.
	// 	Subsequent requests should only include the Context.
	// 	A hard block (an SBSD POST with a t parameter) always starts a new chain: send Script and an empty
	// 	Context, never the Context from an earlier request.
	Script string `json:"script"`

	// ScriptUrl is the absolute src of the SBSD script tag, query included.
	// Optional: when set, the fingerprint's resource-timing signal carries the
	// script's own entry the way a browser's does.
	ScriptUrl string `json:"scriptUrl,omitempty"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
	Context        string `json:"context"`
}

type DynamicInput struct {
	// Script is the akamai script's contents.
	Script string `json:"script"`
}

type apiResponse struct {
	Payload  string      `json:"payload"`
	Swhanedl string      `json:"swhanedl,omitempty"`
	Context  string      `json:"context,omitempty"`
	TimeZone string      `json:"timeZone,omitempty"`
	ClientId string      `json:"clientId,omitempty"`
	Headers  ClientHints `json:"headers"`
	Error    string      `json:"error"`
}

// ClientHints are the Accept-CH client hint headers for the device a payload
// was generated for, keyed by lowercase header name, with values ready to send
// verbatim (already serialized the way Chrome sends them; do not re-quote).
//
// The API returns every hint Chrome sends once a site opts in via Accept-CH:
// sec-ch-ua, sec-ch-ua-mobile, sec-ch-ua-platform, sec-ch-ua-arch,
// sec-ch-ua-bitness, sec-ch-ua-model, sec-ch-ua-platform-version,
// sec-ch-ua-full-version, sec-ch-ua-full-version-list, sec-ch-ua-wow64,
// sec-ch-ua-form-factors, sec-ch-device-memory, sec-ch-dpr,
// sec-ch-viewport-width, sec-ch-viewport-height, device-memory, dpr,
// viewport-width, ect, rtt, downlink, sec-ch-prefers-color-scheme,
// sec-ch-prefers-reduced-motion and sec-ch-prefers-reduced-transparency.
// Send only the ones the target site asked for in its Accept-CH response
// header; sending hints a site never requested is not what a browser does.
//
// Nil when the API returned none: Safari user agents, and payloads built by a
// remote browser.
type ClientHints map[string]string

type KasadaPayloadInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// IpsLink is the ips.js script link, parsed from the block page (429 status code)
	IpsLink string `json:"ipsLink"`

	// Script is the ips.js script retrieved using the IpsLink url
	Script string `json:"script"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip,omitempty"`

	// Strict refuses the payload instead of returning one built on a guess when
	// the script has a signal the API cannot identify yet. The request then
	// fails with "Unidentified signals in script; refusing payload because
	// strict was set". Off by default: such a payload usually still works.
	Strict bool `json:"strict,omitempty"`
}

type KasadaPowInput struct {
	// St is the x-kpsdk-st value returned by the /tl POST request
	St     int    `json:"st"`
	Ct     string `json:"ct"`
	Fc     string `json:"fc"`
	Domain string `json:"domain"`
	Script string `json:"script"`
	// WorkTime can be used to pre-generate POW strings
	WorkTime *int `json:"workTime,omitempty"`
}

type BotIDHeaderInput struct {
	Script         string `json:"script"`
	UserAgent      string `json:"userAgent"`
	IP             string `json:"ip"`
	AcceptLanguage string `json:"acceptLanguage"`
}

type kasadaPayloadOutput struct {
	Headers     KasadaHeaders `json:"headers"`
	ClientHints ClientHints   `json:"clientHints"`
	Payload     string        `json:"payload"`
	Error       string        `json:"error"`
}

type KasadaHeaders struct {
	XKpsdkCt string `json:"x-kpsdk-ct"`
	XKpsdkDt string `json:"x-kpsdk-dt"`
	XKpsdkV  string `json:"x-kpsdk-v"`
	XKpsdkR  string `json:"x-kpsdk-r"`
	XKpsdkDv string `json:"x-kpsdk-dv"`
	XKpsdkH  string `json:"x-kpsdk-h"`
	XKpsdkFc string `json:"x-kpsdk-fc"`
	XKpsdkIm string `json:"x-kpsdk-im"`
}

type DataDomeSliderInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// DeviceLink is the URL that contains the script and starts like this:
	// https://geo.captcha-delivery.com/captcha/?initialCid
	DeviceLink string `json:"deviceLink"`

	// Html is the response body of the GET request to the DeviceLink
	Html string `json:"html"`

	// Puzzle is the captcha puzzle image bytes, base64 encoded.
	// The URL that returns the puzzle looks like this:
	// https://dd.prod.captcha-delivery.com/image/2024-xx-xx/hash.jpg
	Puzzle string `json:"puzzle"`

	// Piece is the captcha puzzle piece image bytes, base64 encoded.
	// The URL that returns the puzzle looks like this:
	// https://dd.prod.captcha-delivery.com/image/2024-xx-xx/hash.frag.png
	Piece string `json:"piece"`

	ParentUrl string `json:"parentUrl"`

	// Script is the DataDome challenge bundle the page loads with a
	// <script defer src="..."> tag. Use datadome.ParseChallengeScriptURL on the
	// Html to find its URL, then GET that URL with the same client and pass the
	// response body here.
	//
	// Always set it. The session sends only its ScriptHash, and sends the
	// bundle itself when the API answers that it has not seen it yet.
	Script string `json:"script"`

	// ScriptHash is the lowercase hex sha256 of Script (see ScriptHash). The
	// session fills it in; set it yourself only to send a hash without the
	// bundle, in which case an unknown hash returns ErrScriptRequired.
	ScriptHash string `json:"scriptHash,omitempty"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type DataDomeInterstitialInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// DeviceLink is the URL that contains the script and starts like this:
	// https://geo.captcha-delivery.com/captcha/?initialCid
	DeviceLink string `json:"deviceLink"`

	// Html is the response body of the GET request to the DeviceLink
	Html string `json:"html"`

	// Script is the DataDome challenge bundle the page loads with a
	// <script defer src="..."> tag. Use datadome.ParseChallengeScriptURL on the
	// Html to find its URL, then GET that URL with the same client and pass the
	// response body here.
	//
	// Always set it. The session sends only its ScriptHash, and sends the
	// bundle itself when the API answers that it has not seen it yet.
	Script string `json:"script"`

	// ScriptHash is the lowercase hex sha256 of Script (see ScriptHash). The
	// session fills it in; set it yourself only to send a hash without the
	// bundle, in which case an unknown hash returns ErrScriptRequired.
	ScriptHash string `json:"scriptHash,omitempty"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type DataDomeTagsInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	Cid     string `json:"cid"`
	Ddk     string `json:"ddk"`
	Referer string `json:"referer"`
	Type    string `json:"type"`

	Version        string `json:"version"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}
