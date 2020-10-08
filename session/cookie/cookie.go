package cookie

import (
	"net/http"
	"time"
)

type Cookie interface {
	// Options returns the cookie options.
	Options() Options
	// Map the cookie init.
	Init(opts ...Option)
	// Map returns the cookie items as map[string]string.
	Map() map[string]string
	// Contains checks if given key exists and not expired in cookie.
	Contains(key string) bool
	// Set sets cookie item with default domain, path and expiration age.
	Set(key, value string)
	// SetCookie sets cookie item given given domain, path and expiration age.
	// The optional parameter <httpOnly> specifies if the cookie item is only available in HTTP,
	// which is usually empty.
	SetCookie(key, value, domain, path string, maxAge time.Duration, httpOnly ...bool)
	// SetHttpCookie sets cookie with *http.Cookie.
	SetHttpCookie(cookie http.Cookie)
	// GetSessionId retrieves and returns the session id from cookie.
	GetSessionId() string
	// SetSessionId sets session id in the cookie.
	SetSessionId(id string)
}