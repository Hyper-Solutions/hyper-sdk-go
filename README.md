# Hyper Solutions SDK

## Installation

To use the Hyper Solutions SDK in your Go project, you need to install it using the following command:

```
go get github.com/Hyper-Solutions/hyper-sdk-go
```

## Usage

### Creating a Session

To start using the SDK, you need to create a new `Session` instance by providing your API key:

```go
session := hyper.NewSession("your-api-key")
```

You can also optionally set a JWT key and a custom HTTP client:

```go
session := hyper.NewSession("your-api-key").
    WithJwtKey("your-jwt-key").
    WithClient(customHTTPClient)
```

## Akamai

The Akamai package provides functions for interacting with Akamai Bot Manager, including generating sensor data, parsing script path, parsing pixel challenges, and handling sec-cpt challenges.

### Generating Sensor Data

To generate sensor data required for generating valid Akamai cookies, use the `GenerateSensorData` function:

```go
sensorData, err := session.GenerateSensorData(ctx, &hyper.SensorInput{
    // Set the required input fields
})
if err != nil {
    // Handle the error
}
```

### Parsing Script Path

To parse the Akamai Bot Manager script path from the given HTML code, use the `ParseScriptPath` function:

```go
scriptPath, err := akamai.ParseScriptPath(reader)
if err != nil {
    // Handle the error
}
```

### Handling Sec-Cpt Challenges

The Akamai package provides functions for handling sec-cpt challenges:

- `ParseSecCptChallenge`: Parses a sec-cpt challenge from an `io.Reader`.
- `ParseSecCptChallengeFromJson`: Parses a sec-cpt challenge from an `io.Reader`.
- `GenerateSecCptPayload`: Generates a sec-cpt payload using the provided sec-cpt cookie.
- `Sleep`: Sleeps for the duration specified in the sec-cpt challenge.
- `SleepWithContext`: Sleeps for the duration specified in the sec-cpt challenge, this is context aware.

### Validating Cookies

The Akamai package provides functions for validating cookies:

- `IsCookieValid`: Determines if the provided `_abck` cookie value is valid based on the given request count.
- `IsCookieInvalidated`: Determines if the current session requires more sensors to be sent.


### Generating Pixel Data

To generate pixel data, use the `GeneratePixelData` function:

```go
pixelData, err := session.GeneratePixelData(ctx, &hyper.PixelInput{
    // Set the required input fields
})
if err != nil {
    // Handle the error
}
```

### Parsing Pixel Challenges

The Akamai package provides functions for parsing pixel challenges:

- `ParsePixelHtmlVar`: Parses the required pixel challenge variable from the given HTML code.
- `ParsePixelScriptURL`: Parses the script URL of the pixel challenge script and the URL to post a generated payload to from the given HTML code.
- `ParsePixelScriptVar`: Parses the dynamic value from the pixel script.
## Incapsula

The Incapsula package provides functions for interacting with Incapsula, including generating Reese84 sensor data, UTMVC cookies, and parsing UTMVC script paths.

### Generating Reese84 Sensor

To generate sensor data required for generating valid Reese84 cookies, use the `GenerateReese84Sensor` function:

```go
sensorData, err := session.GenerateReese84Sensor(ctx, site, userAgent)
if err != nil {
    // Handle the error
}
```

### Generating UTMVC Cookie

To generate the UTMVC cookie using the Hyper Solutions API, use the `GenerateUtmvcCookie` function:

```go
utmvcCookie, err := session.GenerateUtmvcCookie(ctx, &hyper.UtmvcInput{
    Script: "your-script",
    SessionIds: []string{"session-id-1", "session-id-2"},
    UserAgent: "user-agent-here"
})
if err != nil {
    // Handle the error
}
```

### Parsing UTMVC Script Path

To parse the UTMVC script path from a given script content, use the `ParseUtmvcScriptPath` function:

```go
scriptPath, err := incapsula.ParseUtmvcScriptPath(scriptReader)
if err != nil {
    // Handle the error
}
```

### Generating UTMVC Submit Path

To generate a unique UTMVC submit path with a random query parameter, use the `GetUtmvcSubmitPath` function:

```go
submitPath := incapsula.GetUtmvcSubmitPath()
```

## Contributing

If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License

This SDK is licensed under the [MIT License](LICENSE).
