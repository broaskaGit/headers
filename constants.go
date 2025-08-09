package headers

// Browser and Platform Constants
const (
	// Browser information
	BrowserName             = "Chrome"
	ChromeVersion           = "139"
	ChromeVersionFull       = ChromeVersion + ".0.7258.66"
	UserAgentDefault        = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + ChromeVersion + ".0.0.0 Safari/537.36"
	SecCHUserAgentDefault   = `"Chromium";v="` + ChromeVersion + `", "Google Chrome";v="` + ChromeVersion + `", "Not-A.Brand";v="99"`
	SecCHFullVersionDefault = `"Chromium";v="` + ChromeVersionFull + `", "Google Chrome";v="` + ChromeVersionFull + `", "Not-A.Brand";v="99.0.0.0"`

	// Platform information
	OSName                         = "Linux"
	OSVersion                      = "6.8.0"
	SecCHPlatformDefault           = `"` + OSName + `"`
	SecCHPlatformVersionDefault    = `"` + OSVersion + `"`
	SecCHMobileDefault             = "?0"
	SecCHModelDefault              = ""
	SecCHPrefersColorSchemeDefault = "light"

	// System identifier
	UDID = OSName + "/" + BrowserName
)

// Common Header Values
const (
	AcceptDefault         = "*/*"
	AcceptLanguageDefault = "en-US,en;q=0.9"
	AcceptEncodingDefault = "gzip, deflate, br, zstd"
	ConnectionDefault     = "keep-alive"
)

// ContentType Constants
const (
	// JSON content types
	ContentTypeJSON     ContentType = "application/json"
	ContentTypeJSONUTF8 ContentType = "application/json; charset=utf-8"

	// XML content types
	ContentTypeXML     ContentType = "application/xml"
	ContentTypeXMLUTF8 ContentType = "application/xml; charset=utf-8"

	// Form content types
	ContentTypeForm            ContentType = "application/x-www-form-urlencoded"
	ContentTypeFormUTF8        ContentType = "application/x-www-form-urlencoded; charset=utf-8"
	ContentTypeMultipart       ContentType = "multipart/form-data"
	ContentTypeMultipartPrefix ContentType = "multipart/form-data; boundary="

	// Text content types
	ContentTypeTextPlain      ContentType = "text/plain"
	ContentTypeTextPlainUTF8  ContentType = "text/plain; charset=utf-8"
	ContentTypeTextHTML       ContentType = "text/html"
	ContentTypeTextHTMLUTF8   ContentType = "text/html; charset=utf-8"
	ContentTypeTextCSS        ContentType = "text/css"
	ContentTypeTextJavaScript ContentType = "text/javascript"
	ContentTypeTextXML        ContentType = "text/xml"
	ContentTypeTextCSV        ContentType = "text/csv"

	// Binary content types
	ContentTypeApplicationOctet ContentType = "application/octet-stream"
	ContentTypeApplicationPDF   ContentType = "application/pdf"
	ContentTypeApplicationZip   ContentType = "application/zip"
	ContentTypeApplicationGzip  ContentType = "application/gzip"

	// Image content types
	ContentTypeImageJPEG ContentType = "image/jpeg"
	ContentTypeImagePNG  ContentType = "image/png"
	ContentTypeImageGIF  ContentType = "image/gif"
	ContentTypeImageWebP ContentType = "image/webp"
	ContentTypeImageSVG  ContentType = "image/svg+xml"
	ContentTypeImageICO  ContentType = "image/x-icon"
	ContentTypeImageBMP  ContentType = "image/bmp"
	ContentTypeImageTIFF ContentType = "image/tiff"

	// Audio content types
	ContentTypeAudioMP3  ContentType = "audio/mpeg"
	ContentTypeAudioWAV  ContentType = "audio/wav"
	ContentTypeAudioOGG  ContentType = "audio/ogg"
	ContentTypeAudioAAC  ContentType = "audio/aac"
	ContentTypeAudioFLAC ContentType = "audio/flac"

	// Video content types
	ContentTypeVideoMP4  ContentType = "video/mp4"
	ContentTypeVideoAVI  ContentType = "video/x-msvideo"
	ContentTypeVideoMOV  ContentType = "video/quicktime"
	ContentTypeVideoWEBM ContentType = "video/webm"
	ContentTypeVideoOGV  ContentType = "video/ogg"
)

// Connection Constants
const (
	ConnectionKeepAlive Connection = "keep-alive"
	ConnectionClose     Connection = "close"
	ConnectionUpgrade   Connection = "upgrade"
)

// Accept Constants
const (
	AcceptAll       Accept = "*/*"
	AcceptJSON      Accept = "application/json"
	AcceptXML       Accept = "application/xml"
	AcceptHTML      Accept = "text/html"
	AcceptText      Accept = "text/plain"
	AcceptImage     Accept = "image/*"
	AcceptAudio     Accept = "audio/*"
	AcceptVideo     Accept = "video/*"
	AcceptJSONXML   Accept = "application/json, application/xml"
	AcceptHTMLXHTML Accept = "text/html, application/xhtml+xml"
	AcceptImageWebP Accept = "image/webp,image/apng,image/*,*/*;q=0.8"
)

// AcceptLanguage Constants
const (
	AcceptLanguageEnglish   AcceptLanguage = "en-US,en;q=0.9"
	AcceptLanguageSpanish   AcceptLanguage = "es-ES,es;q=0.9"
	AcceptLanguageFrench    AcceptLanguage = "fr-FR,fr;q=0.9"
	AcceptLanguageGerman    AcceptLanguage = "de-DE,de;q=0.9"
	AcceptLanguageChinese   AcceptLanguage = "zh-CN,zh;q=0.9"
	AcceptLanguageJapanese  AcceptLanguage = "ja-JP,ja;q=0.9"
	AcceptLanguageUniversal AcceptLanguage = "en-US,en;q=0.9,*;q=0.5"
)

// AcceptEncoding Constants
const (
	AcceptEncodingAll      AcceptEncoding = "gzip, deflate, br, zstd"
	AcceptEncodingGzip     AcceptEncoding = "gzip"
	AcceptEncodingDeflate  AcceptEncoding = "deflate"
	AcceptEncodingBrotli   AcceptEncoding = "br"
	AcceptEncodingZstd     AcceptEncoding = "zstd"
	AcceptEncodingIdentity AcceptEncoding = "identity"
)

// CacheControl Constants
const (
	CacheControlNoCache              CacheControl = "no-cache"
	CacheControlNoStore              CacheControl = "no-store"
	CacheControlMaxAge0              CacheControl = "max-age=0"
	CacheControlMaxAge3600           CacheControl = "max-age=3600"
	CacheControlMaxAge86400          CacheControl = "max-age=86400"
	CacheControlMustRevalidate       CacheControl = "must-revalidate"
	CacheControlPublic               CacheControl = "public"
	CacheControlPrivate              CacheControl = "private"
	CacheControlImmutable            CacheControl = "immutable"
	CacheControlStaleWhileRevalidate CacheControl = "stale-while-revalidate=86400"
)

// Authorization Constants
const (
	AuthorizationBearerPrefix Authorization = "Bearer "
	AuthorizationBasicPrefix  Authorization = "Basic "
	AuthorizationDigestPrefix Authorization = "Digest "
	AuthorizationOAuth        Authorization = "OAuth "
	AuthorizationAPIKey       Authorization = "ApiKey "
)

// Pragma Constants
const (
	PragmaNoCache Pragma = "no-cache"
	PragmaPublic  Pragma = "public"
)

// SecFetchDest Constants
const (
	SecFetchDestEmpty         SecFetchDest = "empty"
	SecFetchDestDocument      SecFetchDest = "document"
	SecFetchDestImage         SecFetchDest = "image"
	SecFetchDestScript        SecFetchDest = "script"
	SecFetchDestStyle         SecFetchDest = "style"
	SecFetchDestAudio         SecFetchDest = "audio"
	SecFetchDestVideo         SecFetchDest = "video"
	SecFetchDestFont          SecFetchDest = "font"
	SecFetchDestFrame         SecFetchDest = "frame"
	SecFetchDestIFrame        SecFetchDest = "iframe"
	SecFetchDestManifest      SecFetchDest = "manifest"
	SecFetchDestObject        SecFetchDest = "object"
	SecFetchDestServiceWorker SecFetchDest = "serviceworker"
	SecFetchDestSharedWorker  SecFetchDest = "sharedworker"
	SecFetchDestWorker        SecFetchDest = "worker"
)

// SecFetchMode Constants
const (
	SecFetchModeCORS       SecFetchMode = "cors"
	SecFetchModeNavigate   SecFetchMode = "navigate"
	SecFetchModeNoCORS     SecFetchMode = "no-cors"
	SecFetchModeSameOrigin SecFetchMode = "same-origin"
	SecFetchModeWebSocket  SecFetchMode = "websocket"
)

// SecFetchSite Constants
const (
	SecFetchSiteCrossOrigin SecFetchSite = "cross-origin"
	SecFetchSiteSameOrigin  SecFetchSite = "same-origin"
	SecFetchSiteSameSite    SecFetchSite = "same-site"
	SecFetchSiteNone        SecFetchSite = "none"
)

// SecFetchUser Constants
const (
	SecFetchUserTrue  SecFetchUser = "?1"
	SecFetchUserFalse SecFetchUser = "?0"
)

// DNT Constants
const (
	DNTEnable  DNT = "1"
	DNTDisable DNT = "0"
)

// XRequestedWith Constants
const (
	XRequestedWithXMLHttpRequest XRequestedWith = "XMLHttpRequest"
	XRequestedWithFlash          XRequestedWith = "ShockwaveFlash"
)

// XFrameOptions Constants
const (
	XFrameOptionsDeny       XFrameOptions = "DENY"
	XFrameOptionsSameOrigin XFrameOptions = "SAMEORIGIN"
	XFrameOptionsAllowFrom  XFrameOptions = "ALLOW-FROM"
)

// XContentTypeOptions Constants
const (
	XContentTypeOptionsNoSniff XContentTypeOptions = "nosniff"
)

// CORS Constants
const (
	AccessControlAllowOriginAll        AccessControlAllowOrigin      = "*"
	AccessControlAllowMethodsAll       AccessControlAllowMethods     = "*"
	AccessControlAllowMethodsCommon    AccessControlAllowMethods     = "GET, POST, PUT, DELETE, OPTIONS"
	AccessControlAllowHeadersAll       AccessControlAllowHeaders     = "*"
	AccessControlAllowHeadersCommon    AccessControlAllowHeaders     = "Content-Type, Authorization, X-Requested-With"
	AccessControlAllowCredentialsTrue  AccessControlAllowCredentials = "true"
	AccessControlAllowCredentialsFalse AccessControlAllowCredentials = "false"
	AccessControlMaxAgeOneHour         AccessControlMaxAge           = "3600"
	AccessControlMaxAgeOneDay          AccessControlMaxAge           = "86400"
)

// HTTP Method Constants
const (
	MethodGET     Method = "GET"
	MethodPOST    Method = "POST"
	MethodPUT     Method = "PUT"
	MethodPATCH   Method = "PATCH"
	MethodDELETE  Method = "DELETE"
	MethodHEAD    Method = "HEAD"
	MethodOPTIONS Method = "OPTIONS"
	MethodTRACE   Method = "TRACE"
	MethodCONNECT Method = "CONNECT"
)

// Upgrade Constants
const (
	UpgradeWebSocket Upgrade = "websocket"
	UpgradeHTTP2     Upgrade = "h2c"
)

// Range Constants
const (
	RangeBytes Range = "bytes="
)

// Transfer Encoding Constants
const (
	TransferEncodingChunked  TransferEncoding = "chunked"
	TransferEncodingCompress TransferEncoding = "compress"
	TransferEncodingDeflate  TransferEncoding = "deflate"
	TransferEncodingGzip     TransferEncoding = "gzip"
	TransferEncodingIdentity TransferEncoding = "identity"
)

// TE Constants
const (
	TEChunked  TE = "chunked"
	TECompress TE = "compress"
	TEDeflate  TE = "deflate"
	TEGzip     TE = "gzip"
	TETrailers TE = "trailers"
)

// Content Disposition Constants
const (
	ContentDispositionInline     ContentDisposition = "inline"
	ContentDispositionAttachment ContentDisposition = "attachment"
	ContentDispositionFormData   ContentDisposition = "form-data"
)

// Common User Agents
const (
	UserAgentChrome  UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"
	UserAgentFirefox UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:133.0) Gecko/20100101 Firefox/133.0"
	UserAgentSafari  UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.2 Safari/605.1.15"
	UserAgentEdge    UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 Edg/139.0.0.0"
	UserAgentMobile  UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 18_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.2 Mobile/15E148 Safari/604.1"
	UserAgentBot     UserAgent = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
)

// Security Header Constants
const (
	StrictTransportSecurityMaxAge     StrictTransportSecurity = "max-age=31536000"
	StrictTransportSecurityIncludeSub StrictTransportSecurity = "max-age=31536000; includeSubDomains"
	StrictTransportSecurityPreload    StrictTransportSecurity = "max-age=31536000; includeSubDomains; preload"
)

// Client Hints Constants
const (
	SecCHUAMobileDesktop SecCHUAMobile = "?0"
	SecCHUAMobileMobile  SecCHUAMobile = "?1"

	SecCHUAPlatformWindows SecCHUAPlatform = `"Windows"`
	SecCHUAPlatformMacOS   SecCHUAPlatform = `"macOS"`
	SecCHUAPlatformLinux   SecCHUAPlatform = `"Linux"`
	SecCHUAPlatformAndroid SecCHUAPlatform = `"Android"`
	SecCHUAPlatformiOS     SecCHUAPlatform = `"iOS"`

	SecCHPrefersColorSchemeLight SecCHPrefersColorScheme = "light"
	SecCHPrefersColorSchemeDark  SecCHPrefersColorScheme = "dark"

	SecCHPrefersReducedMotionNoPreference SecCHPrefersReducedMotion = "no-preference"
	SecCHPrefersReducedMotionReduce       SecCHPrefersReducedMotion = "reduce"
)

// Custom API Header Names (commonly used)
const (
	HeaderAPIKey             = "X-API-Key"
	HeaderAPIVersion         = "X-API-Version"
	HeaderRequestID          = "X-Request-ID"
	HeaderTimestamp          = "X-Timestamp"
	HeaderSignature          = "X-Signature"
	HeaderNonce              = "X-Nonce"
	HeaderRateLimit          = "X-RateLimit-Limit"
	HeaderRateLimitRemaining = "X-RateLimit-Remaining"
	HeaderRateLimitReset     = "X-RateLimit-Reset"
)
