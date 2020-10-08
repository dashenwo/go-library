package session

import (
	"github.com/dashenwo/go-library/session/config"
	"github.com/dashenwo/go-library/session/cookie"
	"github.com/dashenwo/go-library/session/identifiers"
	"github.com/dashenwo/go-library/session/storage"
	"time"
)

type Manager struct {
	ttl     time.Duration   // session life time
	storage storage.Storage // Underlying storage driver
	config  *config.Config  // Session and cookie configuration information
}

func NewSessionManage(opts ...ManageOption) *Manager {
	// 创建当前对象
	manager := Manager{
		ttl:     time.Duration(3600),
		storage: storage.DefaultStore,
		config:  config.DefaultConfig,
	}
	for _, o := range opts {
		o(&manager)
	}
	return &manager
}

// 设置存储源
func (m *Manager) New(opts ...Option) *Session {
	s := &Session{
		manager:     m,
		config:      m.config,
		cookie:      cookie.DefaultCookie,
		identifiers: identifiers.DefaultIdentifiers,
	}
	for _, o := range opts {
		o(s)
	}
	// 设置完成后,初始化
	s.cookie.Init(
		cookie.Config(m.config),
	)
	// 检测是否添加了cookie的request和responseWriter
	if !s.cookie.IsOk() {
		return nil
	}
	s.SetId(s.cookie.GetSessionId())
	s.SetSource(s.cookie.GetSessionSource())
	return s
}

// SetStorage sets the session storage for manager.
func (m *Manager) SetStorage(storage storage.Storage) {
	m.storage = storage
}

// SetName sets the session name for manager.
func (m *Manager) SetConfig(config *config.Config) {
	m.config = config
}
