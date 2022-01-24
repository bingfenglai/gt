package cache

import "time"

	


type Cache interface {
	Set(key string, value interface{},expiration time.Duration) bool
	SetWithDefaultExpiration(key string, value interface{}) bool

	Get(key string) (bool,interface{})
	Keys(keyPrefix string) (bool,[]string)

	Delete(key string)bool
}

