package storage

import (
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"time"
)

type noopStore struct {
	opts Options
}

func (r *noopStore) Set(id string, key string, value interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

func (r *noopStore) GetSession(id string, ttl time.Duration, data *hashmap.Map) (*hashmap.Map, error) {

	panic("implement me")
}

func (r *noopStore) UpdateTTL(id string, ttl time.Duration) error {
	panic("implement me")
}

func (r *noopStore) Init(option ...Option) error {
	panic("implement me")
}

func (r *noopStore) Key(id string) string {
	panic("implement me")
}

func (r *noopStore) Lock(key string) error {
	panic("implement me")
}

func (r *noopStore) UnLock(key string) error {
	panic("implement me")
}

func (r *noopStore) Options() Options {
	panic("implement me")
}

func (r *noopStore) SetSession(id string, data *hashmap.Map, ttl time.Duration) error {
	panic("implement me")
}

func NewDefaultStorage(opts ...Option) Storage {
	// 设置存储驱动
	s := new(noopStore)
	options := Options{
		Prefix:       "session",
		MaxLockWait:  30,
		SpinLockWait: 150,
	}
	for _, o := range opts {
		o(&options)
	}
	s.opts = options
	return s
}
