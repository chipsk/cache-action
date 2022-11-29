package cache

import "time"

const defaultExpiredDuration = 60 * time.Second

type Storage int

const (
	Memory Storage = 1
	Redis  Storage = 2
)
