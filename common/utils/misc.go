package utils

import (
	"github.com/bytedance/sonic"
)

// Misc Miscellaneous
type Misc struct {
}

// Encode 序列化要保存的值
func Encode(val interface{}) (interface{}, error) {
	var value interface{}
	switch v := val.(type) {
	case string, int, uint, int8, int16, int32, int64, float32, float64, bool:
		value = v
	default:
		b, _ := sonic.Marshal(v)
		value = string(b)
	}
	return value, nil
}

func Decode(reply interface{}, err error, val interface{}) error {
	str, err := String(reply, err)
	if err != nil {
		return err
	}
	return sonic.Unmarshal([]byte(str), val)
}
