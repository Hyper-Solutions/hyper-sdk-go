package hyper

import "time"

// Each Generate* method returns one of these. Fields the API adds later become
// new fields here instead of new return values, so they never break callers.

// SensorResult is returned by [Session.GenerateSensorData].
type SensorResult struct {
	// Payload is the sensor data to post.
	Payload string
	// Context must be sent back on the next sensor request of the session.
	Context string
	// ClientHints are the Accept-CH headers for the device the payload describes.
	ClientHints ClientHints
}

// TelemetryResult is returned by [Session.GenerateTelemetry].
type TelemetryResult struct {
	// Header is the ready-to-send value of the Akamai-BM-Telemetry header.
	Header string
	// Context must be sent back on the next telemetry request of the session.
	Context string
	// ClientHints are the Accept-CH headers for the device the header describes.
	ClientHints ClientHints
}

// TelemetryParamsResult is returned by [Session.GetTelemetryParams].
type TelemetryParamsResult struct {
	// Path is the path of the site's get_params request, exactly as the browser
	// requests it. Send a GET to it on the page's origin; the "j" in its JSON
	// response goes into [TelemetryInput.J].
	Path string
	// Context can replace the script on the first [Session.GenerateTelemetry] request.
	Context string
}

// PixelResult is returned by [Session.GeneratePixelData].
type PixelResult struct {
	// Payload is the pixel data to post.
	Payload string
}

// SbsdResult is returned by [Session.GenerateSbsdData].
type SbsdResult struct {
	// Payload is the sbsd body to post.
	Payload string
	// Context must be sent back on the next sbsd request of the session.
	Context string
	// ClientHints are the Accept-CH headers for the device the payload describes.
	ClientHints ClientHints
}

// Reese84Result is returned by [Session.GenerateReese84Sensor].
type Reese84Result struct {
	// Payload is the reese84 sensor to post.
	Payload string
	// ClientHints are the Accept-CH headers for the device the payload
	// describes. Nil for Safari user agents.
	ClientHints ClientHints
}

// UtmvcResult is returned by [Session.GenerateUtmvcCookie].
type UtmvcResult struct {
	// Payload is the ___utmvc cookie value.
	Payload string
	// Swhanedl is the swhanedl value for the utmvc request.
	Swhanedl string
}

// KasadaPayloadResult is returned by [Session.GenerateKasadaPayload].
type KasadaPayloadResult struct {
	// Payload is the body to POST to /tl.
	Payload []byte
	// Headers are the x-kpsdk-* headers to send with that POST.
	Headers KasadaHeaders
	// ClientHints are the Accept-CH headers for the device the payload
	// describes. Unlike Headers, send only the ones the site asked for.
	ClientHints ClientHints
}

// KasadaVersionResult is returned by [Session.GetKasadaVersion].
type KasadaVersionResult struct {
	// Version is the x-kpsdk-v value, such as "j-1.2.797". Send it verbatim
	// on the /fp request and in the x-kpsdk-v header.
	Version string
}

// UsageResult is returned by [Session.GetUsage].
type UsageResult struct {
	// BalanceEuro is the pay-as-you-go balance in euros.
	BalanceEuro float64 `json:"balanceEuro"`
	// Plans are the plans of the API key, one per product.
	Plans []UsagePlan `json:"plans"`
}

// UsagePlan is the quota of one plan in a [UsageResult].
type UsagePlan struct {
	// Product is the product the plan covers, such as "akamai" or "datadome".
	Product string `json:"product"`
	// RequestsUsed is the number of requests used from the plan.
	RequestsUsed int64 `json:"requestsUsed"`
	// RequestsQuota is the number of requests the plan allows.
	RequestsQuota int64 `json:"requestsQuota"`
	// RequestsRemaining is RequestsQuota minus RequestsUsed.
	RequestsRemaining int64 `json:"requestsRemaining"`
	// ExpiresAt is when the plan expires.
	ExpiresAt time.Time `json:"expiresAt"`
}

// KasadaPowResult is returned by [Session.GenerateKasadaPow].
type KasadaPowResult struct {
	// Payload is the x-kpsdk-cd value.
	Payload string
}

// BotIDHeaderResult is returned by [Session.GenerateBotIDHeader].
type BotIDHeaderResult struct {
	// Payload is the x-is-human header value.
	Payload string
}

// DataDomeSliderResult is returned by [Session.GenerateDataDomeSlider].
type DataDomeSliderResult struct {
	// Payload is the URL that returns a solved datadome cookie.
	Payload string
	// ClientHints are the Accept-CH headers for the device the payload describes.
	ClientHints ClientHints
}

// DataDomeInterstitialResult is returned by [Session.GenerateDataDomeInterstitial].
type DataDomeInterstitialResult struct {
	// Payload is the form body to POST for a solved datadome cookie.
	Payload string
	// ClientHints are the Accept-CH headers for the device the payload describes.
	ClientHints ClientHints
}

// DataDomeTagsResult is returned by [Session.GenerateDataDomeTags].
type DataDomeTagsResult struct {
	// Payload is the tags body to POST.
	Payload string
	// ClientHints are the Accept-CH headers for the device the payload describes.
	ClientHints ClientHints
}
