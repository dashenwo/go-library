package config

import "time"

var (
	DefaultConfig = New()
)

type Config struct {
	// ==================================
	// Cookie.
	// ==================================

	// CookieMaxAge specifies the max TTL for cookie items.
	CookieMaxAge time.Duration

	// CookiePath specifies cookie path.
	// It also affects the default storage for session id.
	CookiePath string

	// CookieDomain specifies cookie domain.
	// It also affects the default storage for session id.
	CookieDomain string

	// ==================================
	// Session.
	// ==================================

	// SessionMaxAge specifies max TTL for session items.
	SessionMaxAge time.Duration

	// SessionIdName specifies the session id name.
	SessionIdName string
}

func New(opts ...Option) *Config {
	config := Config{
		CookieMaxAge:      time.Hour * 24 * 365,
		CookiePath:        "/",
		CookieDomain:      "",
		SessionMaxAge:     time.Hour * 24,
		SessionIdName:     "gosessionid",
	}
	for _, o := range opts {
		o(&config)
	}
	return &config
}