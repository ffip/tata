package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	Session *redis.Client
	ctx     context.Context
}

// 初始化连接
func (r *Redis) Connect(uri string) (err error) {
	info, err := redis.ParseURL(uri)
	if err != nil {
		return err
	}

	r.Session = redis.NewClient(info)

	r.ctx = context.Background()

	_, err = r.Session.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 设置项
func (r *Redis) Set(key string, value interface{}, expire time.Duration) (err error) {
	err = r.Session.Set(key, value, expire).Err()
	return
}

// 删除项
func (r *Redis) Del(key string) (err error) {
	err = r.Session.Del(key).Err()
	return
}

// 获取项
func (r *Redis) Keys(key string) (val []string, err error) {
	val, err = r.Session.Keys(key).Result()
	return
}

// 获取项
func (r *Redis) Get(key string) (val string, err error) {
	val, err = r.Session.Get(key).Result()
	return
}

// 清空表
func (r *Redis) Flush() (err error) {
	err = r.Session.FlushDB().Err()
	return
}
