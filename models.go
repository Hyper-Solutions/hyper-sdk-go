package hyper

type UtmvcInput struct {
	UserAgent string `json:"userAgent"`

	// SessionIds The value of each cookie that has a name that starts with incap_ses_
	SessionIds []string `json:"sessionIds"`

	// Script is the UTMVC script contents
	Script string `json:"script"`
}
type ReeseInput struct {
	UserAgent string `json:"userAgent"`
	Language  string `json:"language"`
	IP        string `json:"ip"`
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

	// ScriptHash is a sha256 checksum over the akamai script contents
	ScriptHash string `json:"scriptHash"`

	// DynamicValues is required for sites that use the dynamic version of v3, this value can be retrieved by [Session.ParseV3Dynamic]
	DynamicValues string `json:"dynamicValues"`

	Language string `json:"language"`
	IP       string `json:"ip"`
}

type PixelInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`
	HTMLVar   string `json:"htmlVar"`
	ScriptVar string `json:"scriptVar"`
	Language  string `json:"language"`
	IP        string `json:"ip"`
}

type SbsdInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`
	Uuid      string `json:"uuid"`
	PageUrl   string `json:"pageUrl"`
	OCookie   string `json:"o"`
	Script    string `json:"script"`
	Language  string `json:"language"`
	IP        string `json:"ip"`
}

type DynamicInput struct {
	// Script is the akamai script's contents.
	Script string `json:"script"`
}

type apiResponse struct {
	Payload string   `json:"payload"`
	Headers *Headers `json:"headers"`
	Error   string   `json:"error"`
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

	// Language is the first language of your accept-language header, it defaults to "en-US" if left empty.
	Language string `json:"language,omitempty"`
	IP       string `json:"ip,omitempty"`
}

type KasadaPowInput struct {
	// St is the x-kpsdk-st value returned by the /tl POST request
	St int `json:"st"`

	// WorkTime can be used to pre-generate POW strings
	WorkTime *int `json:"workTime,omitempty"`
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

	ParentUrl string `json:"parentUrl"`
	Language  string `json:"language"`
	IP        string `json:"ip"`
}

type DataDomeInterstitialInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	// DeviceLink is the URL that contains the script and starts like this:
	// https://geo.captcha-delivery.com/captcha/?initialCid
	DeviceLink string `json:"deviceLink"`

	// Html is the response body of the GET request to the DeviceLink
	Html string `json:"html"`

	Language string `json:"language,omitempty"`
	IP       string `json:"ip"`
}

type DataDomeTagsInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`

	Cid     string `json:"cid"`
	Ddk     string `json:"ddk"`
	Referer string `json:"referer"`
	Type    string `json:"type"`

	Version  string `json:"version"`
	Language string `json:"language"`
	IP       string `json:"ip"`
}
