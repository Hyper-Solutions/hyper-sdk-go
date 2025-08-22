# Hyper Solutions SDK - Go Library for Bot Protection Bypass (Akamai, Incapsula, Kasada, DataDome)

![Go Version](https://img.shields.io/badge/Go-1.22+-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub Release](https://img.shields.io/github/v/release/Hyper-Solutions/hyper-sdk-go)
![GitHub Stars](https://img.shields.io/github/stars/Hyper-Solutions/hyper-sdk-go)

[![](https://dcbadge.limes.pink/api/server/akamai)](https://discord.gg/akamai)

A powerful **Go SDK** for bypassing modern bot protection systems including **Akamai Bot Manager**, **Incapsula**, **Kasada**, and **DataDome**. Generate valid cookies, solve anti-bot challenges, and automate protected endpoints with ease.

Perfect for **web scraping**, **automation**, and **data collection** from protected websites.

## 🚀 Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/Hyper-Solutions/hyper-sdk-go/v2"
)

func main() {
    session := hyper.NewSession("your-api-key")
    
    // Generate Akamai sensor data
    sensorData, sensorContext, err := session.GenerateSensorData(context.Background(), &hyper.SensorInput{
        // Configure your sensor input
    })
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Generated sensor data: %s", sensorData)
    fmt.Printf("Sensor context: %s", sensorContext)
}
```

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
go get github.com/Hyper-Solutions/hyper-sdk-go/v2
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
sensorData, err := session.GenerateSensorData(ctx, &hyper.SensorInput{
    // Configure sensor parameters
})
if err != nil {
    // Handle error
}
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
pixelData, err := session.GeneratePixelData(ctx, &hyper.PixelInput{
    // Pixel challenge parameters
})
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
sensorData, err := session.GenerateReese84Sensor(ctx, site, userAgent)
if err != nil {
    // Handle error
}
```

### UTMVC Cookie Generation

Generate **UTMVC cookies** for Incapsula protection bypass:

```go
utmvcCookie, err := session.GenerateUtmvcCookie(ctx, &hyper.UtmvcInput{
    Script: "incapsula-script-content",
    SessionIds: []string{"session-id-1", "session-id-2"},
    UserAgent: "Mozilla/5.0 (compatible bot)"
})
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
payload, headers, err := session.GenerateKasadaPayload(ctx, &hyper.KasadaPayloadInput{
    // Kasada payload configuration
})
if err != nil {
    // Handle error
}
```

### Generating POW Data (CD)

Solve **Kasada Proof-of-Work** challenges for x-kpsdk-cd tokens:

```go
powPayload, err := session.GenerateKasadaPow(ctx, &hyper.KasadaPowInput{
    // POW challenge parameters
})
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

## 🎯 DataDome Solutions

Solve **DataDome captchas** including slider challenges and interstitial pages.

### Interstitial Challenge Solving

Bypass **DataDome interstitial pages**:

```go
payload, headers, err := session.GenerateDataDomeInterstitial(ctx, &hyper.DataDomeInterstitialInput{
    // Interstitial parameters
})
if err != nil {
    // Handle error
}
// POST payload to https://geo.captcha-delivery.com/interstitial/
```

### Slider Captcha Solving

Solve **DataDome slider captchas** automatically:

```go
checkUrl, headers, err := session.GenerateDataDomeSlider(ctx, &hyper.DataDomeSliderInput{
    // Slider challenge parameters
})
if err != nil {
    // Handle error
}
// GET request to checkUrl
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

### Getting Help

- Check our [documentation](https://docs.justhyped.dev)
- Join our [Discord community](https://discord.gg/akamai)

## 📄 License

This SDK is licensed under the [MIT License](LICENSE).

---

**Keywords**: Go SDK, Golang, bot protection bypass, web scraping, Akamai bypass, Incapsula bypass, Kasada bypass, DataDome bypass, anti-bot, captcha solver, automation, reverse engineering, bot detection, web automation