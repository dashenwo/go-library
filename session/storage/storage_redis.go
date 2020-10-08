package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type redisStore struct {
	opts   Options
	client redis.UniversalClient
}

func (r *redisStore) GetSession(id string, ttl time.Duration, data *hashmap.Map) (*hashmap.Map, error) {
	log.Printf("StorageRedis.GetSession: %s, %v", id, ttl)
	res, err := r.client.Get(r.client.Context(), r.Key(id)).Result()
	log.Print(res)
	if err != nil {
		return nil, err
	}
	content := []byte(res)
	if len(content) == 0 {
		return nil, nil
	}
	var m map[interface{}]interface{}
	if err = json.Unmarshal(content, &m); err != nil {
		return nil, err
	}
	if m == nil {
		return nil, nil
	}
	if data == nil {
		return hashmap.NewFrom(m, true), nil
	} else {
		data.Replace(m)
	}
	return data, nil
}

func NewRedisStorage(client redis.UniversalClient, opts ...Option) Storage {
	// 设置存储驱动
	s := new(redisStore)
	options := Options{
		Prefix:       "session",
		MaxLockWait:  30,
		SpinLockWait: 150,
	}
	s.client = client
	for _, o := range opts {
		o(&options)
	}
	s.opts = options
	return s
}

// 初始化
func (r *redisStore) Init(opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}
	// 创建存储驱动
	if r.client == nil {
		return r.configure()
	}
	return nil
}

// 配置信息
func (r *redisStore) configure() error {

	return nil
}

func (r *redisStore) Key(id string) string {
	return r.opts.Prefix + id
}

func (r *redisStore) Lock(key string) error {
	lockKey := key + ".lock"
	lockTtl := r.opts.MaxLockWait + 1
	attempts := (1000 / r.opts.SpinLockWait) * r.opts.MaxLockWait
	for i := 0; i < attempts; i++ {
		err := r.client.SetNX(context.Background(), lockKey, 1, time.Second*time.Duration(lockTtl)).Err()
		if err == nil {
			return nil
		}
		time.Sleep(time.Duration(r.opts.SpinLockWait) * time.Millisecond)
	}
	return errors.New("unable to acquire a session lock")
}

func (r *redisStore) UnLock(key string) error {
	lockKey := key + ".lock"
	return r.client.Del(context.Background(), lockKey).Err()
}

func (r *redisStore) Get(key string) ([]byte, error) {
	// 循环拿到锁
	err := r.Lock(key)
	if err != nil {
		return nil, err
	}
	b, e := r.client.Get(context.Background(), key).Bytes()
	// 操作完毕解锁
	r.UnLock(key)
	if e != nil {
		return nil, e
	}
	return b, e
}

func (r *redisStore) Set(key string, value interface{}) error {

	return nil
}

func (r *redisStore) Options() Options {
	return Options{}
}
