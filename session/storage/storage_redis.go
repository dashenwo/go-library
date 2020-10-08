package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dashenwo/go-library/container/maps/hashmap"
	"github.com/dashenwo/go-library/session/encoders"
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
	// lock
	if err := r.Lock(r.Key(id)); err != nil {
		return nil, err
	}
	defer r.UnLock(r.Key(id))
	res, err := r.client.Get(r.client.Context(), r.Key(id)).Result()
	if err != nil {
		return nil, err
	}
	content, err := r.opts.Encoders.Decode(res)
	if err != nil {
		return nil, err
	}
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

func (r *redisStore) SetSession(id string, data *hashmap.Map, ttl time.Duration) error {
	log.Printf("StorageRedis.SetSession: %s, %v, %v", id, data, ttl)
	if err := r.Lock(r.Key(id)); err != nil {
		return err
	}
	defer r.UnLock(r.Key(id))
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if r.opts.Encoders != nil && content != nil {
		content = []byte(r.opts.Encoders.Encode(content))
	}
	_, err = r.client.Set(r.client.Context(), r.Key(id), string(content), ttl).Result()
	return err
}

func NewRedisStorage(client redis.UniversalClient, opts ...Option) Storage {
	// 设置存储驱动
	s := new(redisStore)
	options := Options{
		Prefix:       "session",
		MaxLockWait:  30,
		SpinLockWait: 150,
		Encoders:     encoders.DefaultEncoders,
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
	return nil
}

func (r *redisStore) Key(id string) string {
	return r.opts.Prefix + ":" + id
}

// Set sets key-value session pair to the storage.
// The parameter <ttl> specifies the TTL for the session id (not for the key-value pair).
func (r *redisStore) Set(id string, key string, value interface{}, ttl time.Duration) error {
	return ErrorDisabled
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

func (r *redisStore) UpdateTTL(id string, ttl time.Duration) error {
	log.Printf("StorageRedis.UpdateTTL: %s, %v", id, ttl)
	if err := r.Lock(r.Key(id)); err != nil {
		return err
	}
	defer r.UnLock(r.Key(id))
	_, err := r.client.Expire(r.client.Context(), r.Key(id), ttl).Result()
	return err
}

func (r *redisStore) Options() Options {
	return Options{}
}
