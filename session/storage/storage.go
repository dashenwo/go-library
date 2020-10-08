package storage

import (
	"errors"
	"github.com/dashenwo/goutils/maps/hashmap"
	"time"
)

var (
	// ErrNotFound is returned when a key doesn't exist
	ErrNotFound = errors.New("not found")
	// DefaultStore is the memory store.
	DefaultStore  = new(noopStore)
)

type Storage interface {
	// 获取session数据
	GetSession(id string, ttl time.Duration, data *hashmap.Map) (*hashmap.Map, error)
	// 初始化操作
	Init(...Option) error
	// 获取存储的key
	Key(id string) string
	// 写入操作的时候上锁
	Lock(key string) error
	// 在写入操作的时候上锁
	UnLock(key string) error
	// 获取值
	Get(key string) ([]byte,error)
	// 设置值
	Set(key string,value interface{}) error
	// Options allows you to view the current options.
	Options() Options
}