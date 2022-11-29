package cache

import (
	"cloudwego/cache-action/biz/model"
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
	"time"
)

func TestSetCache(t *testing.T) {
	var err error
	info := &model.Info{
		Id:   "测试Info",
		Name: "测试Name",
	}

	ttl := 60 * 5

	ctx := context.Background()
	err = New(Redis).SetWithExpired(ctx, "test_key", info, time.Duration(ttl)*time.Second)

	if err != nil {
		fmt.Printf("redis set key error: %s", err)
	}
}

func TestGetCache(t *testing.T) {

	var info *model.Info
	//
	//json.Unmarshal([]byt
	//
	//e("test"), &info)
	//
	//info := new(model.Info)
	//
	//
	//info3 := &model.Info{}

	sonic.Unmarshal([]byte("test"), info)

	ctx := context.Background()
	hit, err := New(Redis).Get(ctx, "test_key", info)
	if !hit {
		fmt.Printf("redis this key is not hit")
	}
	if err != nil {
		fmt.Printf("redis get key error: %s", err)
	}

	fmt.Printf("info: %s", info)
}
