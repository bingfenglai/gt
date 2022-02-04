package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	//"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// mysql
var DB *gorm.DB

var RedisClient *redis.Client

var Log *zap.Logger
