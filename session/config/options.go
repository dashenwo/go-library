package config

import "time"

type Option func(o *Config)

func CookieMaxAge(t time.Duration) Option {
	return func(o *Config) {
		o.CookieMaxAge = t
	}
}

func CookiePath(path string) Option {
	return func(o *Config) {
		o.CookiePath = path
	}
}

func CookieDomain(domain string) Option {
	return func(o *Config) {
		o.CookieDomain =domain
	}
}

func SessionMaxAge(t time.Duration) Option {
	return func(o *Config) {
		o.SessionMaxAge = t
	}
}

func SessionIdName(name string) Option {
	return func(o *Config) {
		o.SessionIdName = name
	}
}
