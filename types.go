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

// UserAgent represents the User-Agent header value
type UserAgent string

// Upgrade represents the Upgrade header value
type Upgrade string

// Range represents the Range header value
type Range string

// ContentDisposition represents the Content-Disposition header value
type ContentDisposition string

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

// StrictTransportSecurity represents the Strict-Transport-Security header value
type StrictTransportSecurity string

// CORS-related header types

// AccessControlAllowOrigin represents the Access-Control-Allow-Origin header value
type AccessControlAllowOrigin string

// AccessControlAllowMethods represents the Access-Control-Allow-Methods header value
type AccessControlAllowMethods string

// AccessControlAllowHeaders represents the Access-Control-Allow-Headers header value
type AccessControlAllowHeaders string

// AccessControlAllowCredentials represents the Access-Control-Allow-Credentials header value
type AccessControlAllowCredentials string

// AccessControlMaxAge represents the Access-Control-Max-Age header value
type AccessControlMaxAge string

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

// SecCHUAMobile represents the Sec-CH-UA-Mobile header value
type SecCHUAMobile string

// SecCHUAPlatform represents the Sec-CH-UA-Platform header value
type SecCHUAPlatform string

// SecCHPrefersColorScheme represents the Sec-CH-Prefers-Color-Scheme header value
type SecCHPrefersColorScheme string

// SecCHPrefersReducedMotion represents the Sec-CH-Prefers-Reduced-Motion header value
type SecCHPrefersReducedMotion string
