package session

import (
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"github.com/dashenwo/go-library/session/config"
	"github.com/dashenwo/go-library/session/cookie"
	"log"
	"time"
)

type Session struct {
	id      string        // Session id.
	data    *hashmap.Map  // Session data.
	cookie  cookie.Cookie // cookie manage
	manager *Manager      // Parent manager.
	config *config.Config // Configuration information of cookie and session
	start   bool
	// idFunc is a callback function used for creating custom session id.
	// This is called if session id is empty ever when session starts.
	idFunc func(ttl time.Duration) (id string)
}

type SessionKey struct {}

func (s *Session) init()  {
	// If already initialized
	if s.start {
		return
	}
	// If there is no id
	if s.id!="" {
		var err error
		// Retrieve stored session data from storage.
		if s.manager.storage != nil {
			if s.data, err = s.manager.storage.GetSession(s.id, s.manager.ttl, s.data); err != nil {
				log.Printf("session restoring failed for id '%s': %v", s.id, err)
			}
		}
		// If it's an invalid or expired session id,
		// it should create a new session id.
		if s.data == nil {
			s.id = ""
		}
	}
	// Use custom session id creating function.
	if s.id == "" && s.idFunc != nil {
		s.id = s.idFunc(s.manager.ttl)
	}
	// Use default session id creating function.
	if s.id == "" {
		s.id = s.manager.identifiers.Generate()
	}
	if s.data == nil {
		s.data = hashmap.New(true)
	}
	s.start = true
}

// 获取session的id标识
func (s *Session) Id() string {
	return s.id
}