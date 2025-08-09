package headers

import (
	"encoding/base64"
	"maps"
	"strconv"
)

// HeaderOpts contains options for building headers using custom types
type HeaderOpts struct {
	ContentType               ContentType
	Accept                    Accept
	AcceptLanguage            AcceptLanguage
	AcceptEncoding            AcceptEncoding
	Connection                Connection
	UserAgent                 UserAgent
	Referer                   Referer
	Origin                    Origin
	Authorization             Authorization
	CacheControl              CacheControl
	Pragma                    Pragma
	DNT                       DNT
	SecFetchDest              SecFetchDest
	SecFetchMode              SecFetchMode
	SecFetchSite              SecFetchSite
	SecFetchUser              SecFetchUser
	XRequestedWith            XRequestedWith
	XFrameOptions             XFrameOptions
	XContentTypeOptions       XContentTypeOptions
	XCSRFToken                XCSRFToken
	StrictTransportSecurity   StrictTransportSecurity
	ContentSecurityPolicy     ContentSecurityPolicy
	AccessControlAllowOrigin  AccessControlAllowOrigin
	AccessControlAllowMethods AccessControlAllowMethods
	AccessControlAllowHeaders AccessControlAllowHeaders
	Range                     Range
	IfModifiedSince           IfModifiedSince
	IfNoneMatch               IfNoneMatch
	ContentDisposition        ContentDisposition
	Custom                    map[string]string
	IncludeSec                bool // Include Sec-CH-* headers
}

// HeaderBuilder provides a fluent interface for building headers with reusable basic headers
type HeaderBuilder struct {
	basicHeaders map[string]string
	opts         HeaderOpts
}

// NewHeaderBuilder creates a new HeaderBuilder with custom basic headers
func NewHeaderBuilder(basicHeaders map[string]string) *HeaderBuilder {
	basics := make(map[string]string)
	for k, v := range basicHeaders {
		basics[k] = v
	}

	return &HeaderBuilder{
		basicHeaders: basics,
		opts: HeaderOpts{
			Custom: make(map[string]string),
		},
	}
}

// Reset clears all options but keeps the basic headers
func (hb *HeaderBuilder) Reset() *HeaderBuilder {
	hb.opts = HeaderOpts{
		Custom: make(map[string]string),
	}
	return hb
}

// SetBasicHeader sets or updates a basic header that will be included in every build
func (hb *HeaderBuilder) SetBasicHeader(key, value string) *HeaderBuilder {
	if hb.basicHeaders == nil {
		hb.basicHeaders = make(map[string]string)
	}
	hb.basicHeaders[key] = value
	return hb
}

// RemoveBasicHeader removes a basic header
func (hb *HeaderBuilder) RemoveBasicHeader(key string) *HeaderBuilder {
	delete(hb.basicHeaders, key)
	return hb
}

// WithContentType sets the Content-Type header using type-safe ContentType
func (hb *HeaderBuilder) WithContentType(contentType ContentType) *HeaderBuilder {
	hb.opts.ContentType = contentType
	return hb
}

// WithContentTypeString sets the Content-Type header using a string (for custom values)
func (hb *HeaderBuilder) WithContentTypeString(contentType string) *HeaderBuilder {
	hb.opts.ContentType = ContentType(contentType)
	return hb
}

// WithAccept sets the Accept header using type-safe Accept
func (hb *HeaderBuilder) WithAccept(accept Accept) *HeaderBuilder {
	hb.opts.Accept = accept
	return hb
}

// WithAcceptString sets the Accept header using a string (for custom values)
func (hb *HeaderBuilder) WithAcceptString(accept string) *HeaderBuilder {
	hb.opts.Accept = Accept(accept)
	return hb
}

// WithAcceptLanguage sets the Accept-Language header using type-safe AcceptLanguage
func (hb *HeaderBuilder) WithAcceptLanguage(lang AcceptLanguage) *HeaderBuilder {
	hb.opts.AcceptLanguage = lang
	return hb
}

// WithAcceptLanguageString sets the Accept-Language header using a string
func (hb *HeaderBuilder) WithAcceptLanguageString(lang string) *HeaderBuilder {
	hb.opts.AcceptLanguage = AcceptLanguage(lang)
	return hb
}

// WithAcceptEncoding sets the Accept-Encoding header using type-safe AcceptEncoding
func (hb *HeaderBuilder) WithAcceptEncoding(encoding AcceptEncoding) *HeaderBuilder {
	hb.opts.AcceptEncoding = encoding
	return hb
}

// WithAcceptEncodingString sets the Accept-Encoding header using a string
func (hb *HeaderBuilder) WithAcceptEncodingString(encoding string) *HeaderBuilder {
	hb.opts.AcceptEncoding = AcceptEncoding(encoding)
	return hb
}

// WithConnection sets the Connection header using type-safe Connection
func (hb *HeaderBuilder) WithConnection(connection Connection) *HeaderBuilder {
	hb.opts.Connection = connection
	return hb
}

// WithConnectionString sets the Connection header using a string (for custom values)
func (hb *HeaderBuilder) WithConnectionString(connection string) *HeaderBuilder {
	hb.opts.Connection = Connection(connection)
	return hb
}

// WithUserAgent sets the User-Agent header using type-safe UserAgent
func (hb *HeaderBuilder) WithUserAgent(userAgent UserAgent) *HeaderBuilder {
	hb.opts.UserAgent = userAgent
	return hb
}

// WithUserAgentString sets the User-Agent header using a string
func (hb *HeaderBuilder) WithUserAgentString(userAgent string) *HeaderBuilder {
	hb.opts.UserAgent = UserAgent(userAgent)
	return hb
}

// WithReferer sets the Referer header using type-safe Referer
func (hb *HeaderBuilder) WithReferer(referer Referer) *HeaderBuilder {
	hb.opts.Referer = referer
	return hb
}

// WithRefererString sets the Referer header using a string
func (hb *HeaderBuilder) WithRefererString(referer string) *HeaderBuilder {
	hb.opts.Referer = Referer(referer)
	return hb
}

// WithOrigin sets the Origin header using type-safe Origin
func (hb *HeaderBuilder) WithOrigin(origin Origin) *HeaderBuilder {
	hb.opts.Origin = origin
	return hb
}

// WithOriginString sets the Origin header using a string
func (hb *HeaderBuilder) WithOriginString(origin string) *HeaderBuilder {
	hb.opts.Origin = Origin(origin)
	return hb
}

// WithAuthorization sets the Authorization header using type-safe Authorization
func (hb *HeaderBuilder) WithAuthorization(auth Authorization) *HeaderBuilder {
	hb.opts.Authorization = auth
	return hb
}

// WithAuthorizationString sets the Authorization header using a string
func (hb *HeaderBuilder) WithAuthorizationString(auth string) *HeaderBuilder {
	hb.opts.Authorization = Authorization(auth)
	return hb
}

// WithBearerToken sets Authorization header with Bearer token
func (hb *HeaderBuilder) WithBearerToken(token string) *HeaderBuilder {
	hb.opts.Authorization = AuthorizationBearerPrefix + Authorization(token)
	return hb
}

// WithBasicAuth sets Authorization header with Basic auth (properly base64 encoded)
func (hb *HeaderBuilder) WithBasicAuth(username, password string) *HeaderBuilder {
	credentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	hb.opts.Authorization = AuthorizationBasicPrefix + Authorization(credentials)
	return hb
}

// WithAPIKey sets Authorization header with API key
func (hb *HeaderBuilder) WithAPIKey(apiKey string) *HeaderBuilder {
	hb.opts.Authorization = AuthorizationAPIKey + Authorization(apiKey)
	return hb
}

// WithCacheControl sets the Cache-Control header using type-safe CacheControl
func (hb *HeaderBuilder) WithCacheControl(cacheControl CacheControl) *HeaderBuilder {
	hb.opts.CacheControl = cacheControl
	return hb
}

// WithCacheControlString sets the Cache-Control header using a string (for custom values)
func (hb *HeaderBuilder) WithCacheControlString(cacheControl string) *HeaderBuilder {
	hb.opts.CacheControl = CacheControl(cacheControl)
	return hb
}

// WithPragma sets the Pragma header using type-safe Pragma
func (hb *HeaderBuilder) WithPragma(pragma Pragma) *HeaderBuilder {
	hb.opts.Pragma = pragma
	return hb
}

// WithPragmaString sets the Pragma header using a string
func (hb *HeaderBuilder) WithPragmaString(pragma string) *HeaderBuilder {
	hb.opts.Pragma = Pragma(pragma)
	return hb
}

// WithDNT sets the DNT (Do Not Track) header using type-safe DNT
func (hb *HeaderBuilder) WithDNT(dnt DNT) *HeaderBuilder {
	hb.opts.DNT = dnt
	return hb
}

// WithDNTString sets the DNT header using a string (for custom values)
func (hb *HeaderBuilder) WithDNTString(dnt string) *HeaderBuilder {
	hb.opts.DNT = DNT(dnt)
	return hb
}

// WithSecFetch sets Sec-Fetch-* headers using type-safe types
func (hb *HeaderBuilder) WithSecFetch(dest SecFetchDest, mode SecFetchMode, site SecFetchSite) *HeaderBuilder {
	hb.opts.SecFetchDest = dest
	hb.opts.SecFetchMode = mode
	hb.opts.SecFetchSite = site
	return hb
}

// WithSecFetchUser sets the Sec-Fetch-User header
func (hb *HeaderBuilder) WithSecFetchUser(user SecFetchUser) *HeaderBuilder {
	hb.opts.SecFetchUser = user
	return hb
}

// WithSecFetchString sets Sec-Fetch-* headers using strings (for custom values)
func (hb *HeaderBuilder) WithSecFetchString(dest, mode, site string) *HeaderBuilder {
	hb.opts.SecFetchDest = SecFetchDest(dest)
	hb.opts.SecFetchMode = SecFetchMode(mode)
	hb.opts.SecFetchSite = SecFetchSite(site)
	return hb
}

// WithXRequestedWith sets the X-Requested-With header using type-safe XRequestedWith
func (hb *HeaderBuilder) WithXRequestedWith(value XRequestedWith) *HeaderBuilder {
	hb.opts.XRequestedWith = value
	return hb
}

// WithXRequestedWithString sets the X-Requested-With header using a string
func (hb *HeaderBuilder) WithXRequestedWithString(value string) *HeaderBuilder {
	hb.opts.XRequestedWith = XRequestedWith(value)
	return hb
}

// WithXMLHttpRequest sets X-Requested-With to XMLHttpRequest for AJAX
func (hb *HeaderBuilder) WithXMLHttpRequest() *HeaderBuilder {
	hb.opts.XRequestedWith = XRequestedWithXMLHttpRequest
	return hb
}

// WithRange sets the Range header for partial content requests
func (hb *HeaderBuilder) WithRange(rangeValue Range) *HeaderBuilder {
	hb.opts.Range = rangeValue
	return hb
}

// WithRangeString sets the Range header using a string
func (hb *HeaderBuilder) WithRangeString(rangeValue string) *HeaderBuilder {
	hb.opts.Range = Range(rangeValue)
	return hb
}

// WithRangeBytes sets the Range header for byte ranges
func (hb *HeaderBuilder) WithRangeBytes(start, end int64) *HeaderBuilder {
	if end > 0 {
		hb.opts.Range = Range("bytes=" + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10))
	} else {
		hb.opts.Range = Range("bytes=" + strconv.FormatInt(start, 10) + "-")
	}
	return hb
}

// WithIfModifiedSince sets the If-Modified-Since header
func (hb *HeaderBuilder) WithIfModifiedSince(value IfModifiedSince) *HeaderBuilder {
	hb.opts.IfModifiedSince = value
	return hb
}

// WithIfModifiedSinceString sets the If-Modified-Since header using a string
func (hb *HeaderBuilder) WithIfModifiedSinceString(value string) *HeaderBuilder {
	hb.opts.IfModifiedSince = IfModifiedSince(value)
	return hb
}

// WithIfNoneMatch sets the If-None-Match header
func (hb *HeaderBuilder) WithIfNoneMatch(value IfNoneMatch) *HeaderBuilder {
	hb.opts.IfNoneMatch = value
	return hb
}

// WithIfNoneMatchString sets the If-None-Match header using a string
func (hb *HeaderBuilder) WithIfNoneMatchString(value string) *HeaderBuilder {
	hb.opts.IfNoneMatch = IfNoneMatch(value)
	return hb
}

// WithContentDisposition sets the Content-Disposition header
func (hb *HeaderBuilder) WithContentDisposition(value ContentDisposition) *HeaderBuilder {
	hb.opts.ContentDisposition = value
	return hb
}

// WithContentDispositionString sets the Content-Disposition header using a string
func (hb *HeaderBuilder) WithContentDispositionString(value string) *HeaderBuilder {
	hb.opts.ContentDisposition = ContentDisposition(value)
	return hb
}

// Security headers

// WithXFrameOptions sets the X-Frame-Options header
func (hb *HeaderBuilder) WithXFrameOptions(value XFrameOptions) *HeaderBuilder {
	hb.opts.XFrameOptions = value
	return hb
}

// WithXContentTypeOptions sets the X-Content-Type-Options header
func (hb *HeaderBuilder) WithXContentTypeOptions(value XContentTypeOptions) *HeaderBuilder {
	hb.opts.XContentTypeOptions = value
	return hb
}

// WithCSRFToken sets the X-CSRF-Token header
func (hb *HeaderBuilder) WithCSRFToken(token XCSRFToken) *HeaderBuilder {
	hb.opts.XCSRFToken = token
	return hb
}

// WithCSRFTokenString sets the X-CSRF-Token header using a string
func (hb *HeaderBuilder) WithCSRFTokenString(token string) *HeaderBuilder {
	hb.opts.XCSRFToken = XCSRFToken(token)
	return hb
}

// WithStrictTransportSecurity sets the Strict-Transport-Security header
func (hb *HeaderBuilder) WithStrictTransportSecurity(value StrictTransportSecurity) *HeaderBuilder {
	hb.opts.StrictTransportSecurity = value
	return hb
}

// WithContentSecurityPolicy sets the Content-Security-Policy header
func (hb *HeaderBuilder) WithContentSecurityPolicy(policy ContentSecurityPolicy) *HeaderBuilder {
	hb.opts.ContentSecurityPolicy = policy
	return hb
}

// WithContentSecurityPolicyString sets the Content-Security-Policy header using a string
func (hb *HeaderBuilder) WithContentSecurityPolicyString(policy string) *HeaderBuilder {
	hb.opts.ContentSecurityPolicy = ContentSecurityPolicy(policy)
	return hb
}

// CORS headers

// WithCORSAllowOrigin sets the Access-Control-Allow-Origin header
func (hb *HeaderBuilder) WithCORSAllowOrigin(origin AccessControlAllowOrigin) *HeaderBuilder {
	hb.opts.AccessControlAllowOrigin = origin
	return hb
}

// WithCORSAllowMethods sets the Access-Control-Allow-Methods header
func (hb *HeaderBuilder) WithCORSAllowMethods(methods AccessControlAllowMethods) *HeaderBuilder {
	hb.opts.AccessControlAllowMethods = methods
	return hb
}

// WithCORSAllowHeaders sets the Access-Control-Allow-Headers header
func (hb *HeaderBuilder) WithCORSAllowHeaders(headers AccessControlAllowHeaders) *HeaderBuilder {
	hb.opts.AccessControlAllowHeaders = headers
	return hb
}

// WithCustom adds a custom header
func (hb *HeaderBuilder) WithCustom(key, value string) *HeaderBuilder {
	if hb.opts.Custom == nil {
		hb.opts.Custom = make(map[string]string)
	}
	hb.opts.Custom[key] = value
	return hb
}

// WithSecHeaders includes Sec-CH-* headers
func (hb *HeaderBuilder) WithSecHeaders() *HeaderBuilder {
	hb.opts.IncludeSec = true
	return hb
}

// WithOpts sets multiple options at once using HeaderOpts struct
func (hb *HeaderBuilder) WithOpts(opts HeaderOpts) *HeaderBuilder {
	if opts.ContentType != "" {
		hb.opts.ContentType = opts.ContentType
	}
	if opts.Accept != "" {
		hb.opts.Accept = opts.Accept
	}
	if opts.AcceptLanguage != "" {
		hb.opts.AcceptLanguage = opts.AcceptLanguage
	}
	if opts.AcceptEncoding != "" {
		hb.opts.AcceptEncoding = opts.AcceptEncoding
	}
	if opts.Connection != "" {
		hb.opts.Connection = opts.Connection
	}
	if opts.UserAgent != "" {
		hb.opts.UserAgent = opts.UserAgent
	}
	if opts.Referer != "" {
		hb.opts.Referer = opts.Referer
	}
	if opts.Origin != "" {
		hb.opts.Origin = opts.Origin
	}
	if opts.Authorization != "" {
		hb.opts.Authorization = opts.Authorization
	}
	if opts.CacheControl != "" {
		hb.opts.CacheControl = opts.CacheControl
	}
	if opts.Pragma != "" {
		hb.opts.Pragma = opts.Pragma
	}
	if opts.DNT != "" {
		hb.opts.DNT = opts.DNT
	}
	if opts.SecFetchDest != "" {
		hb.opts.SecFetchDest = opts.SecFetchDest
	}
	if opts.SecFetchMode != "" {
		hb.opts.SecFetchMode = opts.SecFetchMode
	}
	if opts.SecFetchSite != "" {
		hb.opts.SecFetchSite = opts.SecFetchSite
	}
	if opts.SecFetchUser != "" {
		hb.opts.SecFetchUser = opts.SecFetchUser
	}
	if opts.XRequestedWith != "" {
		hb.opts.XRequestedWith = opts.XRequestedWith
	}
	if opts.Range != "" {
		hb.opts.Range = opts.Range
	}
	if opts.IfModifiedSince != "" {
		hb.opts.IfModifiedSince = opts.IfModifiedSince
	}
	if opts.IfNoneMatch != "" {
		hb.opts.IfNoneMatch = opts.IfNoneMatch
	}
	if opts.ContentDisposition != "" {
		hb.opts.ContentDisposition = opts.ContentDisposition
	}
	if opts.XFrameOptions != "" {
		hb.opts.XFrameOptions = opts.XFrameOptions
	}
	if opts.XContentTypeOptions != "" {
		hb.opts.XContentTypeOptions = opts.XContentTypeOptions
	}
	if opts.XCSRFToken != "" {
		hb.opts.XCSRFToken = opts.XCSRFToken
	}
	if opts.StrictTransportSecurity != "" {
		hb.opts.StrictTransportSecurity = opts.StrictTransportSecurity
	}
	if opts.ContentSecurityPolicy != "" {
		hb.opts.ContentSecurityPolicy = opts.ContentSecurityPolicy
	}
	if opts.AccessControlAllowOrigin != "" {
		hb.opts.AccessControlAllowOrigin = opts.AccessControlAllowOrigin
	}
	if opts.AccessControlAllowMethods != "" {
		hb.opts.AccessControlAllowMethods = opts.AccessControlAllowMethods
	}
	if opts.AccessControlAllowHeaders != "" {
		hb.opts.AccessControlAllowHeaders = opts.AccessControlAllowHeaders
	}
	if opts.Custom != nil {
		for k, v := range opts.Custom {
			hb.WithCustom(k, v)
		}
	}
	if opts.IncludeSec {
		hb.opts.IncludeSec = opts.IncludeSec
	}
	return hb
}

// Preset methods for common configurations

// JSON sets headers for JSON requests using type-safe constants
func (hb *HeaderBuilder) JSON() *HeaderBuilder {
	return hb.WithContentType(ContentTypeJSON).
		WithAccept(AcceptJSON)
}

// JSONUTF8 sets headers for JSON requests with UTF-8 charset
func (hb *HeaderBuilder) JSONUTF8() *HeaderBuilder {
	return hb.WithContentType(ContentTypeJSONUTF8).
		WithAccept(AcceptJSON)
}

// Form sets headers for form data requests using type-safe constants
func (hb *HeaderBuilder) Form() *HeaderBuilder {
	return hb.WithContentType(ContentTypeForm)
}

// FormUTF8 sets headers for form data requests with UTF-8 charset
func (hb *HeaderBuilder) FormUTF8() *HeaderBuilder {
	return hb.WithContentType(ContentTypeFormUTF8)
}

// Multipart sets headers for multipart form data using type-safe constants
func (hb *HeaderBuilder) Multipart() *HeaderBuilder {
	return hb.WithContentType(ContentTypeMultipart)
}

// XML sets headers for XML requests
func (hb *HeaderBuilder) XML() *HeaderBuilder {
	return hb.WithContentType(ContentTypeXML).
		WithAccept(AcceptXML)
}

// XMLUTF8 sets headers for XML requests with UTF-8 charset
func (hb *HeaderBuilder) XMLUTF8() *HeaderBuilder {
	return hb.WithContentType(ContentTypeXMLUTF8).
		WithAccept(AcceptXML)
}

// AJAX sets headers commonly used for AJAX requests using type-safe constants
func (hb *HeaderBuilder) AJAX() *HeaderBuilder {
	return hb.WithXMLHttpRequest().
		WithSecFetch(SecFetchDestEmpty, SecFetchModeCORS, SecFetchSiteSameOrigin)
}

// API sets headers commonly used for API requests
func (hb *HeaderBuilder) API() *HeaderBuilder {
	return hb.JSON().
		WithConnection(ConnectionKeepAlive).
		WithCacheControl(CacheControlNoCache)
}

// Browser sets headers to mimic a browser request
func (hb *HeaderBuilder) Browser() *HeaderBuilder {
	return hb.WithUserAgent(UserAgentChrome).
		WithAccept(AcceptHTMLXHTML).
		WithAcceptLanguage(AcceptLanguageEnglish).
		WithAcceptEncoding(AcceptEncodingAll).
		WithDNT(DNTDisable).
		WithSecHeaders().
		WithSecFetch(SecFetchDestDocument, SecFetchModeNavigate, SecFetchSiteNone)
}

// Mobile sets headers to mimic a mobile browser request
func (hb *HeaderBuilder) Mobile() *HeaderBuilder {
	return hb.WithUserAgent(UserAgentMobile).
		WithAccept(AcceptHTMLXHTML).
		WithAcceptLanguage(AcceptLanguageEnglish).
		WithAcceptEncoding(AcceptEncodingAll)
}

// Bot sets headers to mimic a bot/crawler request
func (hb *HeaderBuilder) Bot() *HeaderBuilder {
	return hb.WithUserAgent(UserAgentBot).
		WithAccept(AcceptHTML).
		WithConnection(ConnectionClose)
}

// NoCache sets headers to prevent caching
func (hb *HeaderBuilder) NoCache() *HeaderBuilder {
	return hb.WithCacheControl(CacheControlNoCache).
		WithPragma(PragmaNoCache)
}

// Build constructs the final headers map, merging basic headers with options
func (hb *HeaderBuilder) Build() map[string]string {
	headers := make(map[string]string)

	// Start with basic headers
	for key, value := range hb.basicHeaders {
		headers[key] = value
	}

	// Override with specific options (only if they're set)
	if hb.opts.Accept != "" {
		headers["Accept"] = string(hb.opts.Accept)
	}
	if hb.opts.AcceptLanguage != "" {
		headers["Accept-Language"] = string(hb.opts.AcceptLanguage)
	}
	if hb.opts.AcceptEncoding != "" {
		headers["Accept-Encoding"] = string(hb.opts.AcceptEncoding)
	}
	if hb.opts.Connection != "" {
		headers["Connection"] = string(hb.opts.Connection)
	}
	if hb.opts.UserAgent != "" {
		headers["User-Agent"] = string(hb.opts.UserAgent)
	}

	// Add optional headers
	if hb.opts.ContentType != "" {
		headers["Content-Type"] = string(hb.opts.ContentType)
	}
	if hb.opts.Referer != "" {
		headers["Referer"] = string(hb.opts.Referer)
	}
	if hb.opts.Origin != "" {
		headers["Origin"] = string(hb.opts.Origin)
	}
	if hb.opts.Authorization != "" {
		headers["Authorization"] = string(hb.opts.Authorization)
	}
	if hb.opts.CacheControl != "" {
		headers["Cache-Control"] = string(hb.opts.CacheControl)
	}
	if hb.opts.Pragma != "" {
		headers["Pragma"] = string(hb.opts.Pragma)
	}
	if hb.opts.DNT != "" {
		headers["DNT"] = string(hb.opts.DNT)
	}
	if hb.opts.SecFetchDest != "" {
		headers["Sec-Fetch-Dest"] = string(hb.opts.SecFetchDest)
	}
	if hb.opts.SecFetchMode != "" {
		headers["Sec-Fetch-Mode"] = string(hb.opts.SecFetchMode)
	}
	if hb.opts.SecFetchSite != "" {
		headers["Sec-Fetch-Site"] = string(hb.opts.SecFetchSite)
	}
	if hb.opts.SecFetchUser != "" {
		headers["Sec-Fetch-User"] = string(hb.opts.SecFetchUser)
	}
	if hb.opts.XRequestedWith != "" {
		headers["X-Requested-With"] = string(hb.opts.XRequestedWith)
	}
	if hb.opts.Range != "" {
		headers["Range"] = string(hb.opts.Range)
	}
	if hb.opts.IfModifiedSince != "" {
		headers["If-Modified-Since"] = string(hb.opts.IfModifiedSince)
	}
	if hb.opts.IfNoneMatch != "" {
		headers["If-None-Match"] = string(hb.opts.IfNoneMatch)
	}
	if hb.opts.ContentDisposition != "" {
		headers["Content-Disposition"] = string(hb.opts.ContentDisposition)
	}

	// Security headers
	if hb.opts.XFrameOptions != "" {
		headers["X-Frame-Options"] = string(hb.opts.XFrameOptions)
	}
	if hb.opts.XContentTypeOptions != "" {
		headers["X-Content-Type-Options"] = string(hb.opts.XContentTypeOptions)
	}
	if hb.opts.XCSRFToken != "" {
		headers["X-CSRF-Token"] = string(hb.opts.XCSRFToken)
	}
	if hb.opts.StrictTransportSecurity != "" {
		headers["Strict-Transport-Security"] = string(hb.opts.StrictTransportSecurity)
	}
	if hb.opts.ContentSecurityPolicy != "" {
		headers["Content-Security-Policy"] = string(hb.opts.ContentSecurityPolicy)
	}

	// CORS headers
	if hb.opts.AccessControlAllowOrigin != "" {
		headers["Access-Control-Allow-Origin"] = string(hb.opts.AccessControlAllowOrigin)
	}
	if hb.opts.AccessControlAllowMethods != "" {
		headers["Access-Control-Allow-Methods"] = string(hb.opts.AccessControlAllowMethods)
	}
	if hb.opts.AccessControlAllowHeaders != "" {
		headers["Access-Control-Allow-Headers"] = string(hb.opts.AccessControlAllowHeaders)
	}

	// Add Sec-CH headers if requested
	if hb.opts.IncludeSec {
		headers["Sec-CH-UA"] = SecCHUserAgentDefault
		headers["Sec-CH-UA-Full-Version-List"] = SecCHFullVersionDefault
		headers["Sec-CH-UA-Platform"] = SecCHPlatformDefault
		headers["Sec-CH-UA-Platform-Version"] = SecCHPlatformVersionDefault
		headers["Sec-CH-UA-Mobile"] = SecCHMobileDefault
		if SecCHModelDefault != "" {
			headers["Sec-CH-UA-Model"] = SecCHModelDefault
		}
		headers["Sec-CH-Prefers-Color-Scheme"] = SecCHPrefersColorSchemeDefault
	}

	// Add custom headers (these can override anything)
	maps.Copy(headers, hb.opts.Custom)

	return headers
}
