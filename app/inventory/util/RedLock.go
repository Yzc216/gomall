package util

import (
	redis "github.com/Yzc216/gomall/app/inventory/biz/dal/redis"
	"github.com/go-redsync/redsync/v4"
	"strconv"
)

const lockPrefix = "inventory_lock:"

func GetLock(skuID uint64) *redsync.Mutex {
	mutexName := lockPrefix + strconv.FormatUint(skuID, 10)
	return redis.RS.NewMutex(mutexName)
}
