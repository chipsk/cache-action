package cache

import (
	"cloudwego/cache-action/common/redis"
	"cloudwego/cache-action/common/utils"
	"context"
	"github.com/bytedance/sonic"
	"time"
)

type redisImpl struct {
}

func (r *redisImpl) Set(ctx context.Context, key string, value interface{}) error {
	v, _ := sonic.Marshal(value)
	err := redis.Set(key, v)
	return err
}

func (r *redisImpl) SetWithExpired(ctx context.Context, key string, value interface{}, expiredInterval time.Duration) error {
	v, err := utils.Encode(value)
	err = redis.SetEx(key, int(expiredInterval.Seconds()), v)
	return err
}

func (r *redisImpl) Get(ctx context.Context, key string, outRefer interface{}) (hit bool, err error) {
	reply, err := redis.Get(key)
	if err != nil {
		return false, err
	}
	//err = sonic.Unmarshal(reply, &outRefer)
	//err = utils.Decode(reply, err, outRefer)
	//str, err := utils.String(reply, err)
	//if err != nil {
	//	return false, err
	//}
	err = sonic.Unmarshal(reply, &outRefer)
	return true, err
}
