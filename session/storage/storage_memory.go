package storage

import (
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"time"
)

type noopStore struct{}

func (r *noopStore) Init(opts ...Option) error {
	return nil
}

func (r *noopStore) GetSession(id string, ttl time.Duration, data *hashmap.Map) (*hashmap.Map, error) {
	panic("implement me")
}

func (r *noopStore) Key(id string) string {
	return ""
}

func (r *noopStore) Lock(key string) error {
	return nil
}

func (r *noopStore) UnLock(key string) error {
	return nil
}

func (r *noopStore) Get(key string) ([]byte, error) {
	return nil, nil
}

func (r *noopStore) Set(key string, value interface{}) error {
	return nil
}

func (r *noopStore) Options() Options {
	return Options{}
}
