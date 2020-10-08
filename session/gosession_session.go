package session

import (
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"github.com/dashenwo/go-library/session/config"
	"github.com/dashenwo/go-library/session/cookie"
	"github.com/dashenwo/go-library/session/identifiers"
	"github.com/dashenwo/go-library/session/storage"
	"log"
)

type Session struct {
	id          string                  // Session id.
	data        *hashmap.Map            // Session data.
	cookie      cookie.Cookie           // cookie manage
	manager     *Manager                // Parent manager.
	config      *config.Config          // Configuration information of cookie and session
	start       bool                    // Used to mark session is modified.
	dirty       bool                    // Used to mark session is modified.
	source      string                  // source of session
	identifiers identifiers.Identifiers // Session ID generator
}

// init does the lazy initialization for session.
// It here initializes real session if necessary.
func (s *Session) init() {
	// If already initialized
	if s.start {
		return
	}
	// If there is no source
	if s.source == "" {
		s.source = "web"
		s.cookie.SetSessionSource(s.source)
	}
	// If there is no id
	if s.id != "" {
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
	// Use default session id creating function.
	if s.id == "" {
		s.id = s.identifiers.Generate()
		s.cookie.SetSessionId(s.id)
	}
	if s.data == nil {
		s.data = hashmap.New(true)
	}
	s.start = true
}

// Close closes current session and updates its ttl in the session manager.
// If this session is dirty, it also exports it to storage.
//
// NOTE that this function must be called ever after a session request done.
func (s *Session) Save() {
	if s.start && s.id != "" {
		size := s.data.Size()
		if s.manager.storage != nil {
			if s.dirty {
				if err := s.manager.storage.SetSession(s.id, s.data, s.manager.ttl); err != nil {
					panic(err)
				}
			} else if size > 0 {
				if err := s.manager.storage.UpdateTTL(s.id, s.manager.ttl); err != nil {
					panic(err)
				}
			}
		}
		if s.cookie.IsOk() {
			s.cookie.Flush()
		}
	}
}

// Set sets key-value pair to this session.
func (s *Session) Set(key string, value interface{}) error {
	s.init()
	if err := s.manager.storage.Set(s.id, key, value, s.manager.ttl); err != nil {
		if err == storage.ErrorDisabled {
			s.data.Put(key, value)
		} else {
			return err
		}
	}
	s.dirty = true
	return nil
}

// 获取session的id标识
func (s *Session) Id() string {
	s.init()
	return s.id
}

// 设置来源
func (s *Session) SetId(id string) {
	s.id = id
}

// 设置来源
func (s *Session) SetSource(source string) {
	s.source = source
}
