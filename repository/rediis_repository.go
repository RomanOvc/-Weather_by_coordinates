package repository

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type RedisRepos struct {
	RedissRepos *redis.Client
}
type RedisToken struct {
	RToken string `json:"rtoken"`
}

func NewRedisRepository(Redis *redis.Client) *RedisRepos {
	return &RedisRepos{RedissRepos: Redis}
}

type RedisReposInterface interface {
	AddDatToRedis(name, token string) (string, error)
	FindByKey(key string) (string, error)
}

//TODO снова об ошибках
func (rr *RedisRepos) AddTokenToRedis(name, token string) (string, error) {
	json, err := json.Marshal(RedisToken{RToken: token})

	if err != nil {
		errors.Wrapf(err, "такая запись существует")
	}

	err = rr.RedissRepos.Set(name, json, time.Hour*1).Err()
	if err != nil {
		errors.Wrapf(err, "такая запись существует")
	}
	return "ok", nil
}

func (rr *RedisRepos) ExistKey(key string) (int64, error) {
	val, err := rr.RedissRepos.Exists(key).Result()
	if err != nil {
		errors.Wrapf(err, "что-то есть")
	}
	return val, err
}

func (rr *RedisRepos) UserRedisToken(key string) (string, error) {

	redisUsertoekn, err := rr.RedissRepos.Get(key).Result()
	if err != nil {
		errors.Wrap(err, "есть редис")
	}

	return redisUsertoekn, err
}
