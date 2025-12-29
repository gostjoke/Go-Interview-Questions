package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// ================================
// Redis Client
// ================================
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// ================================
// Spin Lock
// ================================
type RedisSpinLock struct {
	rdb   *redis.Client
	key   string
	value string
	ttl   time.Duration
}

// 建立鎖
func NewRedisSpinLock(rdb *redis.Client, key string, ttl time.Duration) *RedisSpinLock {
	return &RedisSpinLock{
		rdb:   rdb,
		key:   key,
		value: uuid.NewString(),
		ttl:   ttl,
	}
}

// 嘗試取得鎖
func (l *RedisSpinLock) tryLock(ctx context.Context) (bool, error) {
	return l.rdb.SetNX(ctx, l.key, l.value, l.ttl).Result()
}

// 自旋取得鎖
func (l *RedisSpinLock) Lock(ctx context.Context, spinTimeout time.Duration) error {
	deadline := time.Now().Add(spinTimeout)

	for {
		ok, err := l.tryLock(ctx)
		if err != nil {
			return err
		}

		if ok {
			fmt.Println("🔐 Lock acquired")
			return nil
		}

		if time.Now().After(deadline) {
			return errors.New("spin lock timeout")
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(50 * time.Millisecond):
			// sleep then retry
		}
	}
}

// Lua 安全解鎖
var unlockScript = redis.NewScript(`
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`)

// 解鎖
func (l *RedisSpinLock) Unlock(ctx context.Context) error {
	_, err := unlockScript.Run(ctx, l.rdb, []string{l.key}, l.value).Result()
	if err == nil {
		fmt.Println("🔓 Lock released")
	}
	return err
}

// ================================
// Demo 使用
// ================================
func main() {
	rdb := NewRedisClient()
	defer rdb.Close()

	lock := NewRedisSpinLock(
		rdb,
		"demo:spinlock:order:123",
		5*time.Second, // 鎖 TTL
	)

	err := lock.Lock(ctx, 3*time.Second) // 自旋最多 3 秒
	if err != nil {
		panic(err)
	}

	defer lock.Unlock(ctx)

	// ====== Critical Section ======
	fmt.Println("🚀 Doing critical work...")
	time.Sleep(2 * time.Second)
	fmt.Println("✅ Done")
}
