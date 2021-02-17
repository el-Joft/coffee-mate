package redis

import (
	"coffee-mate/src/utils/apputils"
	"coffee-mate/src/utils/redis"
	"coffee-mate/src/utils/security"
	"coffee-mate/src/validations/schemas"
	"context"
	"fmt"
	"time"
)

// Service -> the propose of redis service is handling business logic application
type Service struct {
	RedisRepository redis.Database
	LoginTokenStore string
	AppConfigStore  string
}

var ctx = context.Background()

// RService -> user service instance
func RService() Service {

	return Service{
		RedisRepository: *redis.Repository(),
		LoginTokenStore: apputils.GetEnv("COFFEE-MATE-APP-AUTH-TOKEN"),
		AppConfigStore:  apputils.GetEnv("REDIS_APP_CONFIG_KEY"),
	}
}

var redisService = RService()

// Load -> method to return all store redis key and value
func (r *Service) Load() map[string]string {
	var keys map[string]string
	var err error
	keys, err = r.RedisHmGetAll(redisService.AppConfigStore)

	if err != nil {
		return keys
	}
	return keys
}

// RedisHSetNX -> method to check if user already exist in database by email or id
func (r *Service) RedisHSetNX(rData schemas.RedisData) error {
	if err := r.RedisRepository.Client.HSetNX(ctx, security.CacheHashKey(rData.ID), security.CacheHashField(), &rData).Err(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	// r.RedisRepository.Client.Expire(r.ctx, security.CacheHashKey(rData.ID), time.Minute)

	return nil
}

//RedisHSet ->
func (r *Service) RedisHSet(rData schemas.RedisData) error {
	if _, err := r.RedisRepository.Client.HSet(ctx, rData.ID, rData.Key, rData.Value).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

//RedisSet ->
func (r *Service) RedisSet(rData schemas.RedisData) error {
	var err error
	_, err = r.RedisRepository.Client.Set(ctx, "my:redis:key", "value", 1*time.Hour).Result()
	if err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

//RedisHmGetAll ->
func (r *Service) RedisHmGetAll(redisKey string) (map[string]string, error) {
	// var err error
	data, err := r.RedisRepository.Client.HGetAll(ctx, redisKey).Result()
	if err != nil {
		return nil, fmt.Errorf("create: redis error: %w", err)
	}
	return data, nil
}

//RedisHVal ->
func (r *Service) RedisHVal(redisKey string) interface{} {
	// var err error
	data, err := r.RedisRepository.Client.HVals(ctx, redisKey).Result()
	if err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return data
}

//RedisHGet ->
func (r *Service) RedisHGet(key, field string) (string, error) {
	// var err error
	data, err := r.RedisRepository.Client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", fmt.Errorf("create: redis error: %w", err)
	}
	return data, nil
}
