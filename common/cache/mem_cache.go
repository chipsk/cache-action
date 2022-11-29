package cache

import (
	"context"
	"time"
)

type memImpl struct {
}

func (m *memImpl) Set(ctx context.Context, key string, value interface{}) error {
	panic("implement mem")
}

func (m *memImpl) SetWithExpired(ctx context.Context, key string, value interface{}, expiredInterval time.Duration) error {
	panic("implement mem")
}

func (m *memImpl) Get(ctx context.Context, key string, outRefer interface{}) (hitin bool, err error) {
	panic("implement mem")
}
