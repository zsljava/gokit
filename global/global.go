package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/zsljava/gokit/config"
	"github.com/zsljava/gokit/util/log"
	"gorm.io/gorm"
	"sync"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VIPER  *viper.Viper
	LOG    *log.Logger
	lock   sync.RWMutex
)

func NewGlobal(db *gorm.DB, rdb *redis.Client, logger *log.Logger) {
	DB = db
	REDIS = rdb
	LOG = logger
}
