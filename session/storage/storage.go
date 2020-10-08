package storage

import (
	"errors"
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"time"
)

var (
	// DefaultStore is the memory store.
	DefaultStore  = NewDefaultStorage()
	ErrorDisabled = errors.New("this feature is disabled in this storage")
)

type Storage interface {
	Set(id string, key string, value interface{}, ttl time.Duration) error
	// 获取session数据
	GetSession(id string, ttl time.Duration, data *hashmap.Map) (*hashmap.Map, error)
	// 保存session数据
	SetSession(id string, data *hashmap.Map, ttl time.Duration) error
	// 更改redis的超时时间
	UpdateTTL(id string, ttl time.Duration) error
	// 初始化操作
	Init(...Option) error
	// 获取存储的key
	Key(id string) string
	// 写入操作的时候上锁
	Lock(key string) error
	// 在写入操作的时候上锁
	UnLock(key string) error
	// Options allows you to view the current options.
	Options() Options
}
