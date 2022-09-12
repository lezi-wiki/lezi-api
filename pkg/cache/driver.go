package cache

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/samber/lo"
	"strings"
)

// store 缓存存储器
var store Driver

// Init 初始化缓存
func Init() {
	if conf.RedisConfig.Server != "" && gin.Mode() != gin.TestMode {
		log.Log().Infof("Redis has been enabled, server: %s", conf.RedisConfig.Server)

		store = NewRedisStore(
			10,
			conf.RedisConfig.Network,
			conf.RedisConfig.Server,
			conf.RedisConfig.Password,
			conf.RedisConfig.DB,
		)
	} else {
		store = NewMemoStore()
	}
}

// Driver 键值缓存存储容器
type Driver interface {
	// Set 设置值，ttl为过期时间，单位为秒
	Set(key string, value interface{}, ttl int) error

	// Get 取值，并返回是否成功
	Get(key string) (interface{}, bool)

	// Gets 批量取值，返回成功取值的map即不存在的值
	Gets(keys []string, prefix string) (map[string]interface{}, []string)

	// Sets 批量设置值，所有的key都会加上prefix前缀
	Sets(values map[string]interface{}, prefix string) error

	// Delete 删除值
	Delete(key string) error

	// Deletes 批量删除值，所有的key都会加上prefix前缀
	Deletes(keys []string, prefix string) error
}

// Set 设置缓存值
func Set(key string, value interface{}, ttl int) error {
	log.Log().Debugf("设置缓存：%s=%v，TTL：%d", key, value, ttl)
	return store.Set(key, value, ttl)
}

// Get 获取缓存值
func Get(key string) (interface{}, bool) {
	log.Log().Debugf("获取缓存：%s", key)
	return store.Get(key)
}

// Delete 删除缓存值
func Delete(key string) error {
	log.Log().Debugf("删除缓存：%s", key)
	return store.Delete(key)
}

// Deletes 删除值
func Deletes(keys []string, prefix string) error {
	go log.Log().Debugf("删除缓存：%s", strings.Join(lo.Map(keys, func(key string, i int) string {
		return prefix + key
	}), ","))
	return store.Deletes(keys, prefix)
}

// GetSettings 根据名称批量获取设置项缓存
func GetSettings(keys []string, prefix string) (map[string]string, []string) {
	raw, miss := store.Gets(keys, prefix)

	res := make(map[string]string, len(raw))
	for k, v := range raw {
		res[k] = v.(string)
	}

	return res, miss
}

// SetSettings 批量设置站点设置缓存
func SetSettings(values map[string]string, prefix string) error {
	var toBeSet = make(map[string]interface{}, len(values))
	for key, value := range values {
		toBeSet[key] = interface{}(value)
	}
	return store.Sets(toBeSet, prefix)
}
