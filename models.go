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

type PixelInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent      string `json:"userAgent"`
	HTMLVar        string `json:"htmlVar"`
	ScriptVar      string `json:"scriptVar"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type SbsdInput struct {
	Index int `json:"index"`
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent      string `json:"userAgent"`
	Uuid           string `json:"uuid"`
	PageUrl        string `json:"pageUrl"`
	OCookie        string `json:"o"`
	Script         string `json:"script"`
	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip"`
}

type DynamicInput struct {
	// Script is the akamai script's contents.
	Script string `json:"script"`
}

type apiResponse struct {
	Payload  string   `json:"payload"`
	Swhanedl string   `json:"swhanedl,omitempty"`
	Context  string   `json:"context,omitempty"`
	TimeZone string   `json:"timeZone,omitempty"`
	ClientId string   `json:"clientId,omitempty"`
	Headers  *Headers `json:"headers"`
	Error    string   `json:"error"`
}

type Headers struct {
	DeviceMemory    string `json:"sec-ch-device-memory"`
	Mobile          string `json:"sec-ch-ua-mobile"`
	Arch            string `json:"sec-ch-ua-arch"`
	Platform        string `json:"sec-ch-ua-platform"`
	Model           string `json:"sec-ch-ua-model"`
	FullVersionList string `json:"sec-ch-ua-full-version-list"`
}

type KasadaPayloadInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// IpsLink is the ips.js script link, parsed from the block page (429 status code)
	IpsLink string `json:"ipsLink"`

	// Script is the ips.js script retrieved using the IpsLink url
	Script string `json:"script"`

	AcceptLanguage string `json:"acceptLanguage"`
	IP             string `json:"ip,omitempty"`
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
	Headers KasadaHeaders `json:"headers"`
	Payload string        `json:"payload"`
	Error   string        `json:"error"`
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

	ParentUrl      string `json:"parentUrl"`
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

type TrustDecisionPayloadInput struct {
	UserAgent      string `json:"userAgent"`
	AcceptLanguage string `json:"acceptLanguage"`
	Ip             string `json:"ip"`
	Script         string `json:"script"`
	PageUrl        string `json:"pageUrl"`
	FpUrl          string `json:"fpUrl"`
}

type TrustDecisionSigningInput struct {
	ClientId string `json:"clientId"`
	Path     string `json:"path"`
}

type TrustDecisionDecodeInput struct {
	Result    string `json:"result"`
	RequestId string `json:"requestId"`
}
