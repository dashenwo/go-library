package session

import (
	"github.com/dashenwo/go-library/session/cookie"
	"github.com/dashenwo/go-library/session/identifiers"
	"net/http"
)

type Option func(s *Session)

// 设置session_Id
func Id(id string) Option {
	return func(s *Session) {
		s.id = id
	}
}

// 设置id生成方法
func Identifiers(identifiers identifiers.Identifiers) Option {
	return func(s *Session) {
		s.identifiers = identifiers
	}
}

func Scheme(scheme string) Option {
	return func(s *Session) {
		s.cookie.Init(cookie.Scheme(scheme))
	}
}

func Source(source string) Option {
	return func(s *Session) {
		s.source = source
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
