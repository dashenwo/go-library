package cookie

import (
	"net/http"
	"time"
)

var (
	DefaultCookie = NewCookie()
)

type cookie struct {
	opts Options
}

func (c *cookie) Init(opts ...Option) {
	for _, o := range opts {
		o(&c.opts)
	}
}

func (c *cookie) Options() Options {
	return c.opts
}

func (c *cookie) GetSessionId() string {
	panic("implement me")
}

func (c *cookie) SetSessionId(id string) {
	panic("implement me")
}

func (c *cookie) Map() map[string]string {
	panic("implement me")
}

func (c *cookie) Contains(key string) bool {
	panic("implement me")
}

func (c *cookie) Set(key, value string) {
	panic("implement me")
}

func (c *cookie) SetCookie(key, value, domain, path string, maxAge time.Duration, httpOnly ...bool) {
	panic("implement me")
}

func (c *cookie) SetHttpCookie(cookie http.Cookie) {
	panic("implement me")
}

func NewCookie(opts ...Option) Cookie {
	// 新建一个自己
	options := NewOptions(opts...)
	// 返回cookie实例
	return &cookie{
		opts: options,
	}
}

func (c *cookie) init()  {
	if c.opts.Data != nil {
		return
	}
	for _, v := range c.opts.Request.Cookies() {
		c.opts.Data[v.Name] = v
	}
}
