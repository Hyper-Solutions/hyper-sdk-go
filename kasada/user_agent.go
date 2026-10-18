package kasada

import (
	"crypto/md5"
	"encoding/hex"
)

// GenerateUserAgent returns the Kasada mobile user-agent value: the lowercase
// hex MD5 of "<osVersion> / <bundleID> <appVersion>". On iOS osVersion is
// UIDevice.systemVersion (e.g. "18.0.1") and bundleID is the app bundle id
// (e.g. "com.example.app"); on Android osVersion is Build.VERSION.SDK_INT
// (e.g. "34") and bundleID is the package name. appVersion is the app version
// string (e.g. "9.15.0").
func GenerateUserAgent(osVersion, bundleID, appVersion string) string {
	sum := md5.Sum([]byte(osVersion + " / " + bundleID + " " + appVersion))
	return hex.EncodeToString(sum[:])
}
