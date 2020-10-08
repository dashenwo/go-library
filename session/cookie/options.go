package cookie

import (
	"github.com/dashenwo/go-library/session/config"
	"net/http"
	"time"
)

type Options struct {
	Scheme  string                  //The default cookie is grpc Scheme.
	Data    map[string]*http.Cookie // Underlying cookie items.
	Path    string                  // The default cookie path.
	Domain  string                  // The default cookie domain
	MaxAge  time.Duration           // The default cookie max age.
	Config  *config.Config          // The default cookie session config.
	Request *http.Request           // Belonged HTTP request.
	Writer  http.ResponseWriter     // Belonged HTTP response
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	opt := Options{
		Scheme: "http",
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func Scheme(scheme string) Option {
	return func(o *Options) {
		o.Scheme = scheme
	}
}

func Path(path string) Option {
	return func(o *Options) {
		o.Path = path
	}
}

func Domain(domain string) Option {
	return func(o *Options) {
		o.Domain = domain
	}
}

func MaxAge(max time.Duration) Option {
	return func(o *Options) {
		o.MaxAge = max
	}
}

func Config(config *config.Config) Option {
	return func(o *Options) {
		o.Config = config
	}
}

func Request(request *http.Request) Option {
	return func(o *Options) {
		o.Request = request
	}
}

func Writer(response http.ResponseWriter) Option {
	return func(o *Options) {
		o.Writer = response
	}
}
