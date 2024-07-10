package hyper

type UtmvcInput struct {
	UserAgent string `json:"userAgent"`

	// SessionIds The value of each cookie that has a name that starts with incap_ses_
	SessionIds []string `json:"sessionIds"`

	// Script is the UTMVC script contents
	Script string `json:"script"`
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
}

type PixelInput struct {
	// UserAgent must be a Chrome Windows User-Agent.
	UserAgent string `json:"userAgent"`
	HTMLVar   string `json:"htmlVar"`
	ScriptVar string `json:"scriptVar"`
}

type apiResponse struct {
	Payload string `json:"payload"`
	Error   string `json:"error"`
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
