package hyper

type UtmvcInput struct {
	UserAgent  string   `json:"userAgent"`
	SessionIds []string `json:"sessionIds"`
	Script     string   `json:"script"`
}

type SensorInput struct {
	Abck       string `json:"abck"`
	Bmsz       string `json:"bmsz"`
	Version    string `json:"version"`
	PageUrl    string `json:"pageUrl"`
	UserAgent  string `json:"userAgent"`
	ScriptHash string `json:"scriptHash"`
}

type PixelInput struct {
	UserAgent string `json:"userAgent"`
	HTMLVar   string `json:"htmlVar"`
	ScriptVar string `json:"scriptVar"`
}

type apiResponse struct {
	Payload string `json:"payload"`
	Error   string `json:"error"`
}
