package session

import (
	"github.com/dashenwo/go-library/session/config"
	"github.com/dashenwo/go-library/session/cookie"
	"github.com/dashenwo/go-library/session/encoders"
	"github.com/dashenwo/go-library/session/identifiers"
	"github.com/dashenwo/go-library/session/storage"
	"time"
)

type Manager struct {
	// 存活时间
	ttl time.Duration
	// 底层存储
	storage storage.Storage
	// id生成器
	identifiers identifiers.Identifiers
	// 序列化
	encoders encoders.Encoders
	// 配置信息
	config *config.Config
}

func NewSessionManage(opts ...ManageOption) *Manager {
	// 创建当前对象
	manager := Manager{
		ttl: time.Duration(3600),
		storage: storage.DefaultStore,
		config: config.DefaultConfig,
		identifiers: identifiers.DefaultIdentifiers,
		encoders: encoders.DefaultEncoders,
	}
	for _, o := range opts {
		o(&manager)
	}
	return &manager
}
// 设置存储源
func (m *Manager) New(opts ...Option) *Session {
	session := &Session{
		manager: m,
		config: m.config,
		cookie: cookie.DefaultCookie,
	}
	for _, o := range opts {
		o(session)
	}
	// 设置完成后,初始化
	session.cookie.Init(
		cookie.Config(m.config),
	)
	return session
}
// SetStorage sets the session storage for manager.
func (m *Manager) SetStorage(storage storage.Storage) {
	m.storage = storage
}
// SetName sets the session name for manager.
func (m *Manager) SetConfig(config *config.Config) {
	m.config = config
}
