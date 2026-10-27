# Hyper Solutions SDK - Go Library for Bot Protection Bypass (Akamai, Incapsula, Kasada, DataDome)

![Go Version](https://img.shields.io/badge/Go-1.22+-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub Release](https://img.shields.io/github/v/release/Hyper-Solutions/hyper-sdk-go)
![GitHub Stars](https://img.shields.io/github/stars/Hyper-Solutions/hyper-sdk-go)

[![](https://dcbadge.limes.pink/api/server/akamai)](https://discord.gg/akamai)

A powerful **Go SDK** for bypassing modern bot protection systems including [**Akamai Bot Manager**](https://hypersolutions.co/products/akamai), [**Incapsula**](https://hypersolutions.co/products/incapsula), [**Kasada**](https://hypersolutions.co/products/kasada), and [**DataDome**](https://hypersolutions.co/products/datadome). Generate valid cookies, solve anti-bot challenges, and automate protected endpoints with ease.

Perfect for **web scraping**, **automation**, and **data collection** from protected websites.

## 🔑 Getting API Access

Before using this SDK, you'll need an API key from Hyper Solutions:

1. **Visit [hypersolutions.co](https://hypersolutions.co/?utm_source=github&utm_medium=sdk_readme&utm_campaign=go_sdk_api_access)** to create your account
2. **Choose your plan**:
    - 💳 **Pay-as-you-go**: Perfect for testing and small-scale usage
    - 📊 **Subscription plans**: Cost-effective for high-volume applications
3. **Get your API key** from the dashboard
4. **Start bypassing bot protection** with this SDK!

## 🚀 Quick Start

```go
package main

import (
	"context"
	"fmt"

	"github.com/Hyper-Solutions/hyper-sdk-go/v4"
)

func main() {
	session := hyper.NewSession("your-api-key")

	// Generate Akamai sensor data
	res, err := session.GenerateSensorData(context.Background(), &hyper.SensorInput{
		// Configure your sensor input
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated sensor data: %s", res.Payload)
	fmt.Printf("Sensor context: %s", res.Context)
}
```

Every `Generate*` method returns a result struct (`*hyper.SensorResult`, `*hyper.SbsdResult`, ...), so fields the API adds later arrive without breaking your code.

### Client hints (Accept-CH)

Sensor, SBSD, Reese84, Kasada payload and DataDome results carry `ClientHints`: the `sec-ch-*`, `device-memory`, `dpr`, `viewport-width`, `ect`, `rtt` and `downlink` headers for the device the payload was generated for, keyed by lowercase header name and ready to send verbatim. Send only the ones the site asked for in its `Accept-CH` response header:

```go
for _, name := range strings.Split(resp.Header.Get("Accept-CH"), ",") {
	name = strings.ToLower(strings.TrimSpace(name))
	if value, ok := res.ClientHints[name]; ok {
		req.Header.Set(name, value)
	}
}
```

`ClientHints` is nil when the API returns none (Safari user agents, and payloads built by a remote browser).

## ✨ Features

- 🛡️ **Akamai Bot Manager**: Generate sensor data, handle pixel challenges, validate cookies
- 🔒 **Incapsula Protection**: Generate Reese84 sensors and UTMVC cookies
- ⚡ **Kasada Bypass**: Generate payload data (CT) and POW tokens (CD)
- 🎯 **DataDome Solutions**: Solve tags, slider captchas and interstitial challenges
- 🔧 **Easy Integration**: Simple Go API with comprehensive documentation
- ⚙️ **Flexible Configuration**: Custom HTTP clients and JWT key support

## 📦 Installation

Install the Hyper Solutions SDK for Go using:

```bash
go get github.com/Hyper-Solutions/hyper-sdk-go/v4
```

## 📋 Table of Contents

- [Quick Start](#-quick-start)
- [Installation](#-installation)
- [Basic Usage](#-basic-usage)
- [Akamai Bot Manager](#-akamai-bot-manager)
- [Incapsula Protection](#-incapsula-protection)
- [Kasada Bypass](#-kasada-bypass)
- [DataDome Solutions](#-datadome-solutions)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)
- [License](#-license)

## 🔧 Basic Usage

### Creating a Session

Initialize the SDK with your API key to start bypassing bot protection:

```go
// Basic session
session := hyper.NewSession("your-api-key")

// Advanced session with custom configuration
session := hyper.NewSession("your-api-key").
    WithJwtKey("your-jwt-key").
    WithClient(customHTTPClient)
```

## 🛡️ Akamai Bot Manager

Bypass **Akamai Bot Manager** protection with sensor data generation, cookie validation, and challenge solving.

### Generating Sensor Data

Generate sensor data for valid **Akamai cookies** and bot detection bypass:

```go
res, err := session.GenerateSensorData(ctx, &hyper.SensorInput{
    // Configure sensor parameters
})
if err != nil {
    // Handle error
}
// res.Payload is the sensor data, res.Context goes into the next SensorInput
```

### SBSD

Generate **SBSD** payloads. The first request of a session sends the script, later ones send the returned context. A hard block (an SBSD POST with a `t` parameter) always sends the script with an empty context, never one from an earlier request:

```go
res, err := session.GenerateSbsdData(ctx, &hyper.SbsdInput{
    // ... other parameters
    Script:    script,
    ScriptUrl: scriptUrl, // optional: absolute src of the SBSD script tag, query included
})
if err != nil {
    // Handle error
}
// POST res.Payload; pass res.Context in the next SbsdInput instead of Script
```

### Parsing Script Path

Extract Akamai Bot Manager script paths from HTML for **reverse engineering**:

```go
scriptPath, err := akamai.ParseScriptPath(reader)
if err != nil {
    // Handle error
}
```

### Handling Sec-Cpt Challenges

Solve **sec-cpt challenges** with built-in parsing and payload generation:

- `ParseSecCptChallenge`: Parse sec-cpt challenges from HTML
- `ParseSecCptChallengeFromJson`: Parse from JSON responses
- `GenerateSecCptPayload`: Generate challenge response payloads
- `Sleep`: Handle challenge timing requirements
- `SleepWithContext`: Context-aware challenge delays

### Cookie Validation

Validate **Akamai _abck cookies** and session states:

- `IsCookieValid`: Check cookie validity for request counts
- `IsCookieInvalidated`: Determine if more sensors are needed

### Pixel Challenge Solving

Handle **Akamai pixel challenges** for advanced bot detection bypass:

```go
res, err := session.GeneratePixelData(ctx, &hyper.PixelInput{
    // Pixel challenge parameters
})
// res.Payload is the pixel data
if err != nil {
    // Handle error
}
```

**Pixel parsing functions**:
- `ParsePixelHtmlVar`: Extract pixel variables from HTML
- `ParsePixelScriptURL`: Get pixel script and POST URLs
- `ParsePixelScriptVar`: Parse dynamic pixel values

## 🔒 Incapsula Protection

Bypass **Incapsula bot detection** with Reese84 sensors and UTMVC cookie generation.

### Generating Reese84 Sensors

Create **Reese84 sensor data** for Incapsula bypass:

```go
res, err := session.GenerateReese84Sensor(ctx, &hyper.ReeseInput{
    // Reese84 parameters
})
// res.Payload is the sensor, res.ClientHints the Accept-CH headers
if err != nil {
    // Handle error
}
```

### UTMVC Cookie Generation

Generate **UTMVC cookies** for Incapsula protection bypass:

```go
res, err := session.GenerateUtmvcCookie(ctx, &hyper.UtmvcInput{
    Script: "incapsula-script-content",
    SessionIds: []string{"session-id-1", "session-id-2"},
    UserAgent: "Mozilla/5.0 (compatible bot)",
})
// res.Payload is the ___utmvc cookie, res.Swhanedl the swhanedl value
if err != nil {
    // Handle error
}
```

### Script Path Parsing

Parse **UTMVC script paths** for Incapsula integration:

```go
scriptPath, err := incapsula.ParseUtmvcScriptPath(scriptReader)
if err != nil {
    // Handle error
}

// Generate submit path
submitPath := incapsula.GetUtmvcSubmitPath()
```

## ⚡ Kasada Bypass

Defeat **Kasada Bot Manager** with payload generation and POW solving.

### Generating Payload Data (CT)

Create **x-kpsdk-ct tokens** for Kasada bypass:

```go
res, err := session.GenerateKasadaPayload(ctx, &hyper.KasadaPayloadInput{
    // Kasada payload configuration

    // Strict refuses a payload built on a signal the API cannot identify yet,
    // instead of returning one built on a guess. Optional, off by default.
    Strict: false,
})
// POST res.Payload to /tl with the x-kpsdk-* values in res.Headers.
// res.ClientHints are the Accept-CH headers: send only the ones the site asked for.
if err != nil {
    // Handle error
}
```

### Generating POW Data (CD)

Solve **Kasada Proof-of-Work** challenges for x-kpsdk-cd tokens:

```go
res, err := session.GenerateKasadaPow(ctx, &hyper.KasadaPowInput{
    // POW challenge parameters
})
// res.Payload is the x-kpsdk-cd value
if err != nil {
    // Handle error
}
```

### Script Path Extraction

Extract **Kasada script paths** from blocked pages (HTTP 429):

```go
scriptPath, err := kasada.ParseScriptPath(reader)
if err != nil {
    // Handle error
}
// Returns: /ips.js?timestamp=...
```

### Mobile User-Agent Generation

Generate the **Kasada mobile `user-agent`** value sent on the native app flow. The mobile SDK derives it as the lowercase hex MD5 of `"<osVersion> / <bundleID> <appVersion>"`:

```go
userAgent := kasada.GenerateUserAgent("18.0.1", "com.example.app", "9.15.0")
// iOS:     osVersion = UIDevice.systemVersion, bundleID = CFBundleIdentifier
// Android: osVersion = Build.VERSION.SDK_INT,  bundleID = package name
// Returns: 7016adc4ba8d433533954ebad33be424
```

Available in `hyper-sdk-go` v3.1.0 and later.

## 🎯 DataDome Solutions

Solve **DataDome captchas** including slider challenges and interstitial pages.

### Interstitial Challenge Solving

Bypass **DataDome interstitial pages**:

```go
res, err := session.GenerateDataDomeInterstitial(ctx, &hyper.DataDomeInterstitialInput{
    // Interstitial parameters
})
if err != nil {
    // Handle error
}
// POST res.Payload to https://geo.captcha-delivery.com/interstitial/
```

### Slider Captcha Solving

Solve **DataDome slider captchas** automatically:

```go
res, err := session.GenerateDataDomeSlider(ctx, &hyper.DataDomeSliderInput{
    // Slider challenge parameters
})
if err != nil {
    // Handle error
}
// GET request to res.Payload (the check URL)
```

### DeviceLink URL Parsing

Extract **DataDome device check URLs**:

```go
// Interstitial device links
deviceLink, err := datadome.ParseInterstitialDeviceCheckLink(reader, datadomeCookie, referer)
if err != nil {
    // Handle error
}

// Slider device links  
deviceLink, err := datadome.ParseSliderDeviceCheckLink(reader, datadomeCookie, referer)
if err != nil {
    // Handle error
}
```

### Challenge Script Parsing

DataDome serves some challenge pages with the bundle **inlined in the HTML** and others with it in **its own file**, switching between the two per request. When a page uses its own file, fetch it and pass it as `Script`:

```go
// After GET-ing the deviceLink and reading the body into html
var script string
if scriptURL, ok := datadome.ParseChallengeScriptURL(html); ok {
    // GET scriptURL with the same client, proxy and headers, then read the body
    script = scriptBody
}

res, err := session.GenerateDataDomeInterstitial(ctx, &hyper.DataDomeInterstitialInput{
    // ... other parameters
    Html:   html,
    Script: script, // empty when the page inlines the bundle
})
```

The same applies to `DataDomeSliderInput`. `Script` is optional and leaving it empty keeps working, but the payload is then built against a pinned script version instead of the one the page actually served.

### Getting Help

- Check our [documentation](https://docs.hypersolutions.co/)
- Join our [Discord community](https://discord.gg/akamai)

## 📄 License

This SDK is licensed under the [MIT License](LICENSE).

---

**Keywords**: Go SDK, Golang, bot protection bypass, web scraping, Akamai bypass, Incapsula bypass, Kasada bypass, DataDome bypass, anti-bot, captcha solver, automation, reverse engineering, bot detection, web automation