package session

import (
	"github.com/dashenwo/go-library/session/config"
	"github.com/dashenwo/go-library/session/encoders"
	"github.com/dashenwo/go-library/session/storage"
)

type ManageOption func(m *Manager)

// 初始化参数，在新建的时候可以传入
func Storage(storage storage.Storage) ManageOption {
	return func(m *Manager) {
		m.storage = storage
	}
}

// 设置序列化
func Encoders(encoders encoders.Encoders) ManageOption {
	return func(m *Manager) {
		m.storage.Init(
			storage.Encoders(encoders),
		)
	}
}

// 设置配置信息
func Config(config *config.Config) ManageOption {
	return func(m *Manager) {
		m.config = config
	}
}
