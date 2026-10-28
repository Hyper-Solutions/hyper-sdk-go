package kasada

import (
	"encoding/base64"
	"encoding/json"
)

// dvKey is the key the Kasada iOS SDK XORs the x-kpsdk-dv header with.
const dvKey = "9d9a1f73-a488-464c-9f73-8c87d018363c"

// GenerateDV returns the x-kpsdk-dv header value the Kasada iOS SDK (i-2.x)
// sends: base64 of {"osVersion":…,"hardwareId":…,"appId":…,"appVersion":…}
// XORed with the SDK's key.
//
// osVersion is UIDevice.systemVersion (e.g. "18.0.1") and hardwareID is the
// model identifier from utsname.machine (e.g. "iPhone13,2" for an iPhone 12),
// not the marketing name. appID is the app's bundle id and appVersion its
// version string, the same values GenerateUserAgent takes, so a user-agent and
// a dv built from one set of inputs describe the same app and device.
//
// Send the same value as the dv field of KasadaPayloadInput, so the payload is
// generated on a device with that model's screen and iOS release.
func GenerateDV(osVersion, hardwareID, appID, appVersion string) string {
	plain, _ := json.Marshal(struct {
		OsVersion  string `json:"osVersion"`
		HardwareID string `json:"hardwareId"`
		AppID      string `json:"appId"`
		AppVersion string `json:"appVersion"`
	}{osVersion, hardwareID, appID, appVersion})
	return base64.StdEncoding.EncodeToString(xorDV(plain))
}

// xorDV XORs b with the first min(len(b), 36) bytes of the key, repeated.
func xorDV(b []byte) []byte {
	key := dvKey
	if len(b) < len(key) {
		key = key[:len(b)]
	}
	out := make([]byte, len(b))
	for i := range b {
		out[i] = b[i] ^ key[i%len(key)]
	}
	return out
}
