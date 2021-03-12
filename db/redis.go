package db

import (
	"context"
	"os"
	"strconv"
	"time"

	"bytes"
	"encoding/gob"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Redis *redis.Client

func ConnectRedisDB() {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
		DB:   db,
	})
}

func SetCache(key string, e interface{}, s time.Duration) {
	var buffer bytes.Buffer

	if err := gob.NewEncoder(&buffer).Encode(e); err == nil {
		Redis.Set(Ctx, key, buffer.Bytes(), s)
	}
}
