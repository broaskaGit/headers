package headers

import (
	"maps"
	"sync"
)

// HeaderOpts contains options for building headers using custom types
type HeaderOpts struct {
	ContentType               ContentType
	Accept                    Accept
	AcceptLanguage            AcceptLanguage
	AcceptEncoding            AcceptEncoding
	Connection                Connection
	UserAgent                 UserAgent
	Referer                   string
	Origin                    string
	Host                      string
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
	XCSRFToken                string
	StrictTransportSecurity   StrictTransportSecurity
	ContentSecurityPolicy     string
	AccessControlAllowOrigin  AccessControlAllowOrigin
	AccessControlAllowMethods AccessControlAllowMethods
	AccessControlAllowHeaders AccessControlAllowHeaders
	Range                     Range
	IfModifiedSince           string
	IfNoneMatch               string
	ContentDisposition        ContentDisposition
	Custom                    map[string]string
	IncludeSecUserAgent       bool // Include Sec-CH-* headers
}

// Builder provides a reusable builder for basic headers
type Builder struct {
	mu           sync.RWMutex
	basicHeaders map[string]string
}

// OptionsBuilder provides a fluent interface for building HeaderOpts
type OptionsBuilder struct {
	opts HeaderOpts
}

// NewBuilder creates a new Builder with custom basic headers
func NewBuilder(basicHeaders map[string]string) *Builder {
	basics := make(map[string]string)
	if basicHeaders != nil {
		maps.Copy(basics, basicHeaders)
	}

	return &Builder{
		basicHeaders: basics,
	}
}

// SetBasicHeader sets or updates a basic header that will be included in every build
func (hb *Builder) SetBasicHeader(key, value string) *Builder {
	hb.mu.Lock()
	defer hb.mu.Unlock()

	if hb.basicHeaders == nil {
		hb.basicHeaders = make(map[string]string)
	}
	hb.basicHeaders[key] = value
	return hb
}

// RemoveBasicHeader removes a basic header
func (hb *Builder) RemoveBasicHeader(key string) *Builder {
	hb.mu.Lock()
	defer hb.mu.Unlock()

	delete(hb.basicHeaders, key)
	return hb
}

// Build constructs the final headers map, merging basic headers with HeaderOpts
func (hb *Builder) Build(opt HeaderOpts) map[string]string {
	headers := make(map[string]string)

	// Start with basic headers
	hb.mu.RLock()
	maps.Copy(headers, hb.basicHeaders)
	hb.mu.RUnlock()

	// Using provided options

	// Override with specific options (only if they're set)
	if opt.Accept != "" {
		headers["Accept"] = string(opt.Accept)
	}
	if opt.AcceptLanguage != "" {
		headers["Accept-Language"] = string(opt.AcceptLanguage)
	}
	if opt.AcceptEncoding != "" {
		headers["Accept-Encoding"] = string(opt.AcceptEncoding)
	}
	if opt.Connection != "" {
		headers["Connection"] = string(opt.Connection)
	}
	if opt.UserAgent != "" {
		headers["User-Agent"] = string(opt.UserAgent)
	}

	// Add optional headers
	if opt.ContentType != "" {
		headers["Content-Type"] = string(opt.ContentType)
	}
	if opt.Referer != "" {
		headers["Referer"] = opt.Referer
	}
	if opt.Origin != "" {
		headers["Origin"] = opt.Origin
	}
	if opt.Host != "" {
		headers["Host"] = opt.Host
	}
	if opt.Authorization != "" {
		headers["Authorization"] = string(opt.Authorization)
	}
	if opt.CacheControl != "" {
		headers["Cache-Control"] = string(opt.CacheControl)
	}
	if opt.Pragma != "" {
		headers["Pragma"] = string(opt.Pragma)
	}
	if opt.DNT != "" {
		headers["DNT"] = string(opt.DNT)
	}
	if opt.SecFetchDest != "" {
		headers["Sec-Fetch-Dest"] = string(opt.SecFetchDest)
	}
	if opt.SecFetchMode != "" {
		headers["Sec-Fetch-Mode"] = string(opt.SecFetchMode)
	}
	if opt.SecFetchSite != "" {
		headers["Sec-Fetch-Site"] = string(opt.SecFetchSite)
	}
	if opt.SecFetchUser != "" {
		headers["Sec-Fetch-User"] = string(opt.SecFetchUser)
	}
	if opt.XRequestedWith != "" {
		headers["X-Requested-With"] = string(opt.XRequestedWith)
	}
	if opt.Range != "" {
		headers["Range"] = string(opt.Range)
	}
	if opt.IfModifiedSince != "" {
		headers["If-Modified-Since"] = opt.IfModifiedSince
	}
	if opt.IfNoneMatch != "" {
		headers["If-None-Match"] = opt.IfNoneMatch
	}
	if opt.ContentDisposition != "" {
		headers["Content-Disposition"] = string(opt.ContentDisposition)
	}

	// Security headers
	if opt.XFrameOptions != "" {
		headers["X-Frame-Options"] = string(opt.XFrameOptions)
	}
	if opt.XContentTypeOptions != "" {
		headers["X-Content-Type-Options"] = string(opt.XContentTypeOptions)
	}
	if opt.XCSRFToken != "" {
		headers["X-CSRF-Token"] = opt.XCSRFToken
	}
	if opt.StrictTransportSecurity != "" {
		headers["Strict-Transport-Security"] = string(opt.StrictTransportSecurity)
	}
	if opt.ContentSecurityPolicy != "" {
		headers["Content-Security-Policy"] = opt.ContentSecurityPolicy
	}

	// CORS headers
	if opt.AccessControlAllowOrigin != "" {
		headers["Access-Control-Allow-Origin"] = string(opt.AccessControlAllowOrigin)
	}
	if opt.AccessControlAllowMethods != "" {
		headers["Access-Control-Allow-Methods"] = string(opt.AccessControlAllowMethods)
	}
	if opt.AccessControlAllowHeaders != "" {
		headers["Access-Control-Allow-Headers"] = string(opt.AccessControlAllowHeaders)
	}

	// Add Sec-CH headers if requested
	if opt.IncludeSecUserAgent {
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
	if opt.Custom != nil {
		maps.Copy(headers, opt.Custom)
	}

	return headers
}

func Build(opts HeaderOpts) map[string]string {
	headers := NewBuilder(nil).Build(opts)
	return headers
}
