package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/go-redis/redis/v8"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/samber/lo"
	"net"
	"time"
)

// RedisStore redis存储驱动
type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

type redisItem struct {
	Value interface{}
}

func serializer(value interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	storeValue := redisItem{
		Value: value,
	}
	err := enc.Encode(storeValue)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func deserializer(value []byte) (interface{}, error) {
	var res redisItem
	buffer := bytes.NewReader(value)
	dec := gob.NewDecoder(buffer)
	err := dec.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

// NewRedisStore 创建新的redis存储
func NewRedisStore(size int, network, address, password string, database int) *RedisStore {
	return &RedisStore{
		client: redis.NewClient(&redis.Options{
			Network:  network,
			Addr:     address,
			Password: password,
			DB:       database,
			Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
				c, err := net.Dial(network, addr)
				if err != nil {
					log.Log().Warnf("无法创建 Redis 连接，%s", err)
					return nil, err
				}

				return c, nil
			},
			IdleTimeout:  time.Second * 240,
			MinIdleConns: size,
			MaxRetries:   3,
		}),
		ctx: context.Background(),
	}
}

// Set 存储值
func (store *RedisStore) Set(key string, value interface{}, ttl int) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)

	serialized, err := serializer(value)
	if err != nil {
		return err
	}

	if ttl > 0 {
		err = rc.SetEX(store.ctx, key, serialized, time.Second*time.Duration(ttl)).Err()
	} else {
		err = rc.Set(store.ctx, key, serialized, 0).Err()
	}

	if err != nil {
		return err
	}
	return nil

}

// Get 取值
func (store *RedisStore) Get(key string) (interface{}, bool) {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)

	v, err := rc.Get(store.ctx, key).Bytes()
	if err != nil || v == nil {
		return nil, false
	}

	finalValue, err := deserializer(v)
	if err != nil {
		return nil, false
	}

	return finalValue, true

}

// Gets 批量取值
func (store *RedisStore) Gets(keys []string, prefix string) (map[string]interface{}, []string) {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)

	v, err := rc.MGet(store.ctx, lo.Map(keys, func(key string, index int) string {
		return prefix + key
	})...).Result()
	if err != nil {
		return nil, keys
	}

	res := make(map[string]interface{})
	missed := make([]string, 0, len(keys))

	for key, value := range v {
		decoded, err := deserializer([]byte(value.(string)))
		if err != nil || decoded == nil {
			missed = append(missed, keys[key])
		} else {
			res[keys[key]] = decoded
		}
	}
	return res, missed
}

// Sets 批量设置值
func (store *RedisStore) Sets(values map[string]interface{}, prefix string) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)
	var setValues = make(map[string]interface{})

	// 编码待设置值
	for key, value := range values {
		serialized, err := serializer(value)
		if err != nil {
			return err
		}
		setValues[prefix+key] = serialized
	}

	_, err := rc.MSet(store.ctx, setValues).Result()
	if err != nil {
		return err
	}
	return nil

}

// Delete 删除给定的键
func (store *RedisStore) Delete(key string) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)

	_, err := rc.Del(store.ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

// Deletes 批量删除给定的键
func (store *RedisStore) Deletes(keys []string, prefix string) error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)

	// 处理前缀
	keys = lo.Map[string, string](keys, func(key string, index int) string {
		return prefix + key
	})

	_, err := rc.Del(store.ctx, keys...).Result()
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll 批量所有键
func (store *RedisStore) DeleteAll() error {
	rc := store.client.Conn(store.ctx)
	defer func(rc *redis.Conn) {
		err := rc.Close()
		if err != nil {
			log.Log().Errorf("Redis 关闭连接错误，%s", err)
		}
	}(rc)

	_, err := rc.FlushDB(store.ctx).Result()

	return err
}
