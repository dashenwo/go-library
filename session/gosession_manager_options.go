package session

import (
	"github.com/dashenwo/go-library/session/config"
	"github.com/dashenwo/go-library/session/encoders"
	"github.com/dashenwo/go-library/session/identifiers"
	"github.com/dashenwo/go-library/session/storage"
)

type ManageOption func(m *Manager)

// 初始化参数，在新建的时候可以传入
func Storage(storage storage.Storage) ManageOption {
	return func(m *Manager) {
		m.storage = storage
	}
}
// 设置id生成方法
func Identifiers(identifiers identifiers.Identifiers) ManageOption {
	return func(m *Manager) {
		m.identifiers = identifiers
	}
}
// 设置序列化
func Encoders(encoders encoders.Encoders) ManageOption  {
	return func(m *Manager) {
		m.encoders = encoders
	}
}
// 设置配置信息
func Config(config *config.Config) ManageOption {
	return func(m *Manager) {
		m.config = config
	}
}