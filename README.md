# THIS PACKAGE IS FULLY AI GENERATED, SO PLEASE CHECK IT BEFORE USING IT

# HTTP Headers Library for Go

A simple, type-safe HTTP headers library for Go that provides a fluent interface for building HTTP headers with reusable basic header configurations.

## Features

- 🔒 **Type-safe header construction** with custom types for all major HTTP headers
- 🔄 **Reusable HeaderBuilder** with configurable basic headers
- 📝 **Extensive constants** for standard header values
- 🎨 **Fluent API** for readable header construction
- 🛡️ **Security headers** support (CORS, CSP, HSTS, etc.)

## Installation

```bash
go get github.com/mightatnight/headers
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/mightatnight/headers"
)

func main() {
    // Create basic headers that will be included in every build
    basicHeaders := map[string]string{
        "User-Agent":      "MyApp/1.0",
        "Accept":          "application/json",
        "Accept-Language": "en-US,en;q=0.9",
        "Connection":      "keep-alive",
    }

    // Create a reusable builder
    builder := headers.NewHeaderBuilder(basicHeaders)

    // Build headers for JSON API request
    apiHeaders := builder.
        WithContentType(headers.ContentTypeJSON).
        WithBearerToken("your-bearer-token").
        WithCustom("X-Request-ID", "req-123").
        Build()

    fmt.Printf("API headers: %+v\n", apiHeaders)

    // Reset and build different headers (basic headers are preserved)
    formHeaders := builder.
        Reset().
        WithContentType(headers.ContentTypeForm).
        WithBasicAuth("user", "password").
        Build()

    fmt.Printf("Form headers: %+v\n", formHeaders)
}
```

## Core Concept

The `HeaderBuilder` maintains a set of "basic headers" that are included in every build, while allowing you to add specific headers for individual requests.

```go
// Create with your basic headers
basicHeaders := map[string]string{
    "User-Agent":    "MyApp/1.0",
    "Accept":        "application/json",
    "X-API-Version": "v1",
}
builder := headers.NewHeaderBuilder(basicHeaders)

// First request
headers1 := builder.
    WithContentType(headers.ContentTypeJSON).
    WithBearerToken("token1").
    Build()

// Reset and build different headers (basic headers are preserved)
headers2 := builder.
    Reset().
    WithContentType(headers.ContentTypeForm).
    WithBasicAuth("user", "pass").
    Build()
```

## Usage Examples

### Type-Safe Headers

```go
basicHeaders := map[string]string{
    "Accept": "*/*",
    "Connection": "keep-alive",
}
builder := headers.NewHeaderBuilder(basicHeaders)

headers := builder.
    WithContentType(headers.ContentTypeJSONUTF8).
    WithAccept(headers.AcceptJSONXML).
    WithAcceptLanguage(headers.AcceptLanguageEnglish).
    WithConnection(headers.ConnectionKeepAlive).
    WithCacheControl(headers.CacheControlNoCache).
    WithDNT(headers.DNTEnable).
    Build()
```

### Authentication Methods

```go
builder := headers.NewHeaderBuilder(map[string]string{
    "User-Agent": "MyClient/1.0",
})

// Bearer token
bearerHeaders := builder.WithBearerToken("your-jwt-token").Build()

// Basic authentication (automatically base64 encoded)
basicHeaders := builder.Reset().WithBasicAuth("username", "password").Build()

// API Key
apiKeyHeaders := builder.Reset().WithAPIKey("your-api-key").Build()
```

### Security Headers

```go
builder := headers.NewHeaderBuilder(map[string]string{})

secureHeaders := builder.
    WithXFrameOptions(headers.XFrameOptionsDeny).
    WithXContentTypeOptions(headers.XContentTypeOptionsNoSniff).
    WithStrictTransportSecurity(headers.StrictTransportSecurityPreload).
    WithCSRFTokenString("csrf-token-123").
    Build()
```

### CORS Headers

```go
builder := headers.NewHeaderBuilder(map[string]string{})

corsHeaders := builder.
    WithCORSAllowOrigin(headers.AccessControlAllowOriginAll).
    WithCORSAllowMethods(headers.AccessControlAllowMethodsCommon).
    WithCORSAllowHeaders(headers.AccessControlAllowHeadersCommon).
    Build()
```

### Preset Methods

```go
builder := headers.NewHeaderBuilder(map[string]string{
    "User-Agent": "MyApp/1.0",
})

// JSON API headers
jsonHeaders := builder.JSON().Build()

// Form headers
formHeaders := builder.Reset().Form().Build()

// AJAX headers
ajaxHeaders := builder.Reset().AJAX().Build()

// Browser-like headers
browserHeaders := builder.Reset().Browser().Build()
```

## Available Types

The library provides type-safe constants for all major HTTP headers:

### Content Types
```go
headers.ContentTypeJSON
headers.ContentTypeJSONUTF8
headers.ContentTypeXML
headers.ContentTypeForm
headers.ContentTypeMultipart
headers.ContentTypeTextPlain
// ... and many more
```

### Accept Types
```go
headers.AcceptAll
headers.AcceptJSON
headers.AcceptXML
headers.AcceptHTML
headers.AcceptJSONXML
// ... and more
```

### Cache Control
```go
headers.CacheControlNoCache
headers.CacheControlNoStore
headers.CacheControlMaxAge3600
headers.CacheControlPublic
```

### Security Headers
```go
headers.XFrameOptionsDeny
headers.XContentTypeOptionsNoSniff
headers.StrictTransportSecurityPreload
```

## API Reference

### HeaderBuilder Methods

#### Basic Configuration
- `NewHeaderBuilder(basicHeaders map[string]string)` - Create with basic headers
- `Reset()` - Clear options but keep basic headers
- `SetBasicHeader(key, value)` - Set/update a basic header
- `RemoveBasicHeader(key)` - Remove a basic header
- `Build()` - Generate the final headers map

#### Content and Accept
- `WithContentType(ContentType)` / `WithContentTypeString(string)`
- `WithAccept(Accept)` / `WithAcceptString(string)`
- `WithAcceptLanguage(AcceptLanguage)` / `WithAcceptLanguageString(string)`
- `WithAcceptEncoding(AcceptEncoding)` / `WithAcceptEncodingString(string)`

#### Authentication
- `WithAuthorization(Authorization)` / `WithAuthorizationString(string)`
- `WithBearerToken(string)`
- `WithBasicAuth(username, password string)`
- `WithAPIKey(string)`

#### Common Headers
- `WithUserAgent(UserAgent)` / `WithUserAgentString(string)`
- `WithReferer(Referer)` / `WithRefererString(string)`
- `WithOrigin(Origin)` / `WithOriginString(string)`
- `WithConnection(Connection)` / `WithConnectionString(string)`

#### Security & CORS
- `WithXFrameOptions(XFrameOptions)`
- `WithXContentTypeOptions(XContentTypeOptions)`
- `WithCSRFToken(XCSRFToken)` / `WithCSRFTokenString(string)`
- `WithStrictTransportSecurity(StrictTransportSecurity)`
- `WithCORSAllowOrigin(AccessControlAllowOrigin)`
- `WithCORSAllowMethods(AccessControlAllowMethods)`
- `WithCORSAllowHeaders(AccessControlAllowHeaders)`

#### Custom
- `WithCustom(key, value string)` - Add custom header
- `WithSecHeaders()` - Include Sec-CH-* headers

#### Preset Methods
- `JSON()` / `JSONUTF8()` - JSON request headers
- `XML()` / `XMLUTF8()` - XML request headers
- `Form()` / `FormUTF8()` - Form request headers
- `Multipart()` - Multipart form headers
- `AJAX()` - AJAX request headers
- `API()` - API request headers
- `Browser()` - Browser-like headers
- `Mobile()` - Mobile browser headers
- `Bot()` - Bot/crawler headers

## Advanced Usage

### API Client with Consistent Headers

```go
// Set up consistent headers for an API client
apiClient := headers.NewHeaderBuilder(map[string]string{
    "User-Agent":      "MyAPIClient/2.1.0",
    "Accept":          "application/json",
    "X-Client-ID":     "client-12345",
    "X-API-Version":   "v2",
})

// Different endpoint calls
userHeaders := apiClient.JSON().WithBearerToken("user-token").Build()
adminHeaders := apiClient.Reset().JSON().WithBearerToken("admin-token").
    WithCustom("X-Admin-Action", "user-management").Build()
```

### Service-to-Service Communication

```go
serviceBuilder := headers.NewHeaderBuilder(map[string]string{
    "User-Agent":        "ServiceMesh/1.0",
    "Accept":            "application/json",
    "X-Service-Name":    "user-service",
    "X-Service-Version": "v2.1.0",
})

tracedHeaders := serviceBuilder.
    JSON().
    WithCustom("X-Trace-ID", "trace-123").
    WithCustom("X-Span-ID", "span-456").
    WithBearerToken("service-jwt-token").
    Build()
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
