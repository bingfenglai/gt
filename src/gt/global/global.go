package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

// mysql
var DB *gorm.DB

var RedisClient *redis.Client

