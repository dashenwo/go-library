package cookie

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"time"
)

var (
	DefaultCookie = NewCookie()
)

type cookie struct {
	opts Options
}

func (c *cookie) GetSessionSource() string {
	return c.Get(c.opts.Config.SessionSourceName)
}

func (c *cookie) SetSessionSource(source string) {
	c.Set(c.opts.Config.SessionSourceName, source)
}

func (c *cookie) Init(opts ...Option) {
	for _, o := range opts {
		o(&c.opts)
	}
}

func (c *cookie) IsOk() bool {
	if c.opts.Request != nil && c.opts.Writer != nil && c.opts.Config != nil {
		return true
	}
	return false
}

func (c *cookie) Options() Options {
	return c.opts
}

func (c *cookie) GetSessionId() string {
	return c.Get(c.opts.Config.SessionIdName)
}

func (c *cookie) SetSessionId(id string) {
	c.Set(c.opts.Config.SessionIdName, id)
}

func (c *cookie) Map() map[string]string {
	c.init()
	m := make(map[string]string)
	for k, v := range c.opts.Data {
		m[k] = v.Value
	}
	return m
}

func (c *cookie) Contains(key string) bool {
	c.init()
	if r, ok := c.opts.Data[key]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return true
		}
	}
	return false
}

func (c *cookie) Set(key, value string) {
	c.SetCookie(key, value, c.opts.Domain, c.opts.Path, c.opts.MaxAge)
}

func (c *cookie) SetCookie(key, value, domain, path string, maxAge time.Duration, httpOnly ...bool) {
	c.init()
	isHttpOnly := false
	if len(httpOnly) > 0 {
		isHttpOnly = httpOnly[0]
	}
	c.opts.Data[key] = &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     path,
		Domain:   domain,
		Expires:  time.Now().Add(maxAge),
		HttpOnly: isHttpOnly,
	}
}

func (c *cookie) SetHttpCookie(cookie *http.Cookie) {
	c.init()
	if cookie.Expires.IsZero() {
		cookie.Expires = time.Now().Add(c.opts.MaxAge)
	}
	c.opts.Data[cookie.Name] = cookie
}

func (c *cookie) Get(key string, def ...string) string {
	c.init()
	if r, ok := c.opts.Data[key]; ok {
		if r.Expires.IsZero() || r.Expires.After(time.Now()) {
			return r.Value
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// Flush outputs the cookie items to client.
func (c *cookie) Flush() {
	if len(c.opts.Data) == 0 {
		return
	}
	header := metadata.New(map[string]string{})
	for _, v := range c.opts.Data {
		// If cookie item is v.Expires.IsZero() means it is set in this request,
		// which should be outputted to client.
		if v.Expires.IsZero() {
			continue
		}
		if c.opts.Scheme == "http" {
			http.SetCookie(c.opts.Writer, v)
		} else {
			header.Append("Set-Cookie", v.String())
		}
	}
	if header.Len() > 0 {
		err := grpc.SendHeader(c.opts.Request.Context(), header)
		log.Print(err)
	}
}

func NewCookie(opts ...Option) Cookie {
	// 新建一个自己
	options := NewOptions(opts...)
	// 返回cookie实例
	return &cookie{
		opts: options,
	}
}

func (c *cookie) init() {
	if c.opts.Data != nil {
		return
	}
	c.opts.Data = make(map[string]*http.Cookie)
	c.opts.Path = c.opts.Config.CookiePath
	c.opts.Domain = c.opts.Config.CookieDomain
	c.opts.MaxAge = c.opts.Config.CookieMaxAge
	for _, v := range c.opts.Request.Cookies() {
		c.opts.Data[v.Name] = v
	}
}
