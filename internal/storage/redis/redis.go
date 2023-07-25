package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/h3ll0kitt1/tg-bot-indecisive-helper/internal/config"
)

type Redis struct {
	client *redis.Client
}

func New(cfg config.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.DbHost + ":" + cfg.DbPort, // Addr: "localhost:6379"
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("Redis ping: %w", err)
	}
	return &Redis{client}, nil
}

func (r Redis) Save(chatId int64, book string) (bool, error) {
	id := strconv.FormatInt(chatId, 10)
	ok, err := r.client.SAdd(id, book).Result()
	if err != nil {
		return false, fmt.Errorf("Redis SAdd: %w", err)
	}

	if ok == 0 {
		return false, nil
	}
	return true, nil
}

func (r Redis) Delete(chatId int64, book string) (bool, error) {
	id := strconv.FormatInt(chatId, 10)
	ok, err := r.client.SRem(id, book).Result()
	if err != nil {
		return false, fmt.Errorf("Redis SRem: %w", err)
	}
	if ok == 0 {
		return false, nil
	}
	return true, nil
}

func (r Redis) Exists(chatId int64, book string) (bool, error) {
	id := strconv.FormatInt(chatId, 10)
	ok, err := r.client.SIsMember(id, book).Result()
	if err != nil {
		return false, fmt.Errorf("Redis sIsMember: %w", err)
	}
	return ok, nil
}

func (r Redis) LenNotZero(chatId int64) (bool, error) {
	id := strconv.FormatInt(chatId, 10)
	len, err := r.client.SCard(id).Result()
	if err != nil {
		return false, fmt.Errorf("Redis sCard: %w", err)
	}

	if len == 0 {
		return false, nil
	}
	return true, nil
}

func (r Redis) List(chatId int64) ([]string, error) {
	id := strconv.FormatInt(chatId, 10)

	list, err := r.client.SMembers(id).Result()
	if err != nil {
		return nil, fmt.Errorf("Redis SMembers: %w", err)
	}
	return list, nil
}

func (r Redis) Rand(chatId int64) (string, error) {
	id := strconv.FormatInt(chatId, 10)
	random, err := r.client.SRandMember(id).Result()
	if err != nil {
		return "", fmt.Errorf("Redis SRandMember: %w", err)
	}
	return random, nil
}
