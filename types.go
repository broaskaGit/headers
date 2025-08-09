package headers

// Core HTTP header types for type-safe header values

// ContentType represents the Content-Type header value
type ContentType string

// Accept represents the Accept header value
type Accept string

// AcceptLanguage represents the Accept-Language header value
type AcceptLanguage string

// AcceptEncoding represents the Accept-Encoding header value
type AcceptEncoding string

// Connection represents the Connection header value
type Connection string

// CacheControl represents the Cache-Control header value
type CacheControl string

// Authorization represents the Authorization header value
type Authorization string

// Pragma represents the Pragma header value
type Pragma string

// XRequestedWith represents the X-Requested-With header value
type XRequestedWith string

// Origin represents the Origin header value
type Origin string

// Referer represents the Referer header value (note: misspelled in HTTP spec)
type Referer string

// Host represents the Host header value
type Host string

// UserAgent represents the User-Agent header value
type UserAgent string

// Upgrade represents the Upgrade header value
type Upgrade string

// Range represents the Range header value
type Range string

// IfModifiedSince represents the If-Modified-Since header value
type IfModifiedSince string

// IfNoneMatch represents the If-None-Match header value
type IfNoneMatch string

// ETag represents the ETag header value
type ETag string

// LastModified represents the Last-Modified header value
type LastModified string

// Location represents the Location header value
type Location string

// SetCookie represents the Set-Cookie header value
type SetCookie string

// Cookie represents the Cookie header value
type Cookie string

// ContentDisposition represents the Content-Disposition header value
type ContentDisposition string

// ContentLength represents the Content-Length header value
type ContentLength string

// TransferEncoding represents the Transfer-Encoding header value
type TransferEncoding string

// TE represents the TE (Transfer Encoding) header value
type TE string

// Security-related header types

// DNT represents the DNT (Do Not Track) header value
type DNT string

// XFrameOptions represents the X-Frame-Options header value
type XFrameOptions string

// XContentTypeOptions represents the X-Content-Type-Options header value
type XContentTypeOptions string

// XCSRFToken represents the X-CSRF-Token header value
type XCSRFToken string

// StrictTransportSecurity represents the Strict-Transport-Security header value
type StrictTransportSecurity string

// ContentSecurityPolicy represents the Content-Security-Policy header value
type ContentSecurityPolicy string

// Proxy and forwarding header types

// XForwardedFor represents the X-Forwarded-For header value
type XForwardedFor string

// XRealIP represents the X-Real-IP header value
type XRealIP string

// XForwardedProto represents the X-Forwarded-Proto header value
type XForwardedProto string

// XForwardedHost represents the X-Forwarded-Host header value
type XForwardedHost string

// XForwardedPort represents the X-Forwarded-Port header value
type XForwardedPort string

// CORS-related header types

// AccessControlAllowOrigin represents the Access-Control-Allow-Origin header value
type AccessControlAllowOrigin string

// AccessControlAllowMethods represents the Access-Control-Allow-Methods header value
type AccessControlAllowMethods string

// AccessControlAllowHeaders represents the Access-Control-Allow-Headers header value
type AccessControlAllowHeaders string

// AccessControlAllowCredentials represents the Access-Control-Allow-Credentials header value
type AccessControlAllowCredentials string

// AccessControlExposeHeaders represents the Access-Control-Expose-Headers header value
type AccessControlExposeHeaders string

// AccessControlMaxAge represents the Access-Control-Max-Age header value
type AccessControlMaxAge string

// AccessControlRequestHeaders represents the Access-Control-Request-Headers header value
type AccessControlRequestHeaders string

// AccessControlRequestMethod represents the Access-Control-Request-Method header value
type AccessControlRequestMethod string

// Sec-Fetch-* header types

// SecFetchDest represents the Sec-Fetch-Dest header value
type SecFetchDest string

// SecFetchMode represents the Sec-Fetch-Mode header value
type SecFetchMode string

// SecFetchSite represents the Sec-Fetch-Site header value
type SecFetchSite string

// SecFetchUser represents the Sec-Fetch-User header value
type SecFetchUser string

// Sec-CH-* header types (Client Hints)

// SecCHUA represents the Sec-CH-UA header value
type SecCHUA string

// SecCHUAMobile represents the Sec-CH-UA-Mobile header value
type SecCHUAMobile string

// SecCHUAPlatform represents the Sec-CH-UA-Platform header value
type SecCHUAPlatform string

// SecCHUAPlatformVersion represents the Sec-CH-UA-Platform-Version header value
type SecCHUAPlatformVersion string

// SecCHUAFullVersionList represents the Sec-CH-UA-Full-Version-List header value
type SecCHUAFullVersionList string

// SecCHUAModel represents the Sec-CH-UA-Model header value
type SecCHUAModel string

// SecCHUAArch represents the Sec-CH-UA-Arch header value
type SecCHUAArch string

// SecCHUABitness represents the Sec-CH-UA-Bitness header value
type SecCHUABitness string

// SecCHUAWoW64 represents the Sec-CH-UA-WoW64 header value
type SecCHUAWoW64 string

// SecCHPrefersColorScheme represents the Sec-CH-Prefers-Color-Scheme header value
type SecCHPrefersColorScheme string

// SecCHPrefersReducedMotion represents the Sec-CH-Prefers-Reduced-Motion header value
type SecCHPrefersReducedMotion string

// SecCHViewportWidth represents the Sec-CH-Viewport-Width header value
type SecCHViewportWidth string

// SecCHDeviceMemory represents the Sec-CH-Device-Memory header value
type SecCHDeviceMemory string

// SecCHDPR represents the Sec-CH-DPR header value
type SecCHDPR string

// SecCHWidth represents the Sec-CH-Width header value
type SecCHWidth string

// HTTP Method types

// Method represents HTTP methods
type Method string

// Custom application header types

// APIKey represents custom API key header values
type APIKey string

// APIVersion represents API version header values
type APIVersion string

// RequestID represents request ID header values
type RequestID string

// Timestamp represents timestamp header values
type Timestamp string

// Signature represents signature header values
type Signature string

// Nonce represents nonce header values
type Nonce string
