package session

import (
	"github.com/dashenwo/go-library/session/cookie"
	"net/http"
)

type Option func(s *Session)

// 设置session_Id
func Id(id string) Option {
	return func(s *Session) {
		s.id = id
	}
}

func Scheme(scheme string) Option {
	return func(s *Session) {
		s.cookie.Init(cookie.Scheme(scheme))
	}
}

func Request(r *http.Request) Option {
	return func(s *Session) {
		s.cookie.Init(cookie.Request(r))
	}
}

func Writer(w http.ResponseWriter) Option {
	return func(s *Session) {
		s.cookie.Init(cookie.Writer(w))
	}
}
