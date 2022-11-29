package cache

import (
	"context"
	"time"
)

type Interface interface {
	Set(ctx context.Context, key string, value interface{}) error

	SetWithExpired(ctx context.Context, key string, value interface{}, expiredInterval time.Duration) error

	Get(ctx context.Context, key string, outRefer interface{}) (hit bool, err error)
}

func New(s Storage) Interface {

	switch s {
	case Memory:
		return &memImpl{}
	case Redis:
		return &redisImpl{}
	default:
		return &redisImpl{}
	}

}
