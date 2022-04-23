package cache

import (
	"time"

	"go.uber.org/zap"
)

// 二级缓存实现，本地缓存未查询到的信息，将从远程redis缓存当中查询，仍然查询不到则返回ErrCacheKeyNotFound,查询到则进行回写
type l2Cache struct {
	l1CacheAdapter Cache
	l2CacheAdapter Cache
}

func newL2Cache(l1CacheAdapter, l2CacheAdapter Cache) *l2Cache {
	l2 := &l2Cache{
		l1CacheAdapter: l1CacheAdapter,
		l2CacheAdapter: l2CacheAdapter,
	}

	go l2.handleL2CacheKeyExpiredEvent()

	go l2.handleL2CacheKeySetEvent()

	go l2.handleL2CacheDelKeyEvent()

	return l2

}

func (receiver *l2Cache) Set(key string, value interface{}, expiration time.Duration) (err error) {
	err = receiver.l2CacheAdapter.Set(key, value, expiration)
	err = receiver.l1CacheAdapter.Set(key, value, expiration)
	return
}
func (receiver *l2Cache) SetWithDefaultExpiration(key string, value interface{}) (err error) {
	err = receiver.l2CacheAdapter.SetWithDefaultExpiration(key, value)
	err = receiver.l1CacheAdapter.SetWithDefaultExpiration(key, value)
	return

}
func (receiver *l2Cache) Get(key string, value interface{}) (err error) {
	err = receiver.l1CacheAdapter.Get(key, value)
	if err == nil {
		return
	}
	zap.L().Warn("本地缓存未获取到数据，将从远程获取")
	err = receiver.l2CacheAdapter.Get(key, value)
	if err == nil {
		// TODO 这里应该先获取远端缓存key的过期时间，再设值
		err = receiver.l1CacheAdapter.SetWithDefaultExpiration(key, value)
	}
	return

}
func (receiver *l2Cache) SetWithJson(key string, value interface{}, expiration time.Duration) (flag bool, str string) {
	flag, str = receiver.l2CacheAdapter.SetWithJson(key, value, expiration)

	flag, str = receiver.l1CacheAdapter.SetWithJson(key, value, expiration)

	return

}
func (receiver *l2Cache) SetWithJsonAndDefaultExpiration(key string, value interface{}) (flag bool, str string) {
	flag, str = receiver.l2CacheAdapter.SetWithJsonAndDefaultExpiration(key, value)

	flag, str = receiver.l1CacheAdapter.SetWithJsonAndDefaultExpiration(key, value)
	return

}
func (receiver *l2Cache) GetWithJson(key string) (flag bool, jsonStr string) {
	flag, jsonStr = receiver.l1CacheAdapter.GetWithJson(key)
	if flag {
		return
	}

	flag, jsonStr = receiver.l2CacheAdapter.GetWithJson(key)
	if flag {
		flag, _ = receiver.l1CacheAdapter.SetWithJsonAndDefaultExpiration(key, jsonStr)
	}

	return
}

// TODO 需要添加二级缓存key同步逻辑，对于本地有，远程没有的进行补写；对于远程有，本地没有的进行回写
func (receiver *l2Cache) Keys(keyPrefix string) (flag bool, strs []string) {
	flag, strs = receiver.l1CacheAdapter.Keys(keyPrefix)
	if flag {
		return
	}
	flag, strs = receiver.l2CacheAdapter.Keys(keyPrefix)
	return
}
func (receiver *l2Cache) Delete(key ...string) (flag bool, count int64) {

	flag, count = receiver.l2CacheAdapter.Delete(key...)
	if flag {
		flag, count = receiver.l1CacheAdapter.Delete(key...)
	}
	return
}

func (receiver *l2Cache) handleL2CacheKeyExpiredEvent() {

	rd := receiver.l2CacheAdapter.(*redisCache)

	pubSub := rd.GetKeyExpiredEventPubSub()

	for {
		zap.L().Info("wait...")
		msg := <-pubSub.Channel()
		zap.L().Info("收到expire消息\n", zap.Any("Pattern", msg.Pattern), zap.Any("payload", msg.Payload), zap.Any("Payload slice", msg.PayloadSlice))
		receiver.l1CacheAdapter.Delete(msg.Payload)
	}

}

func (receiver *l2Cache) handleL2CacheKeySetEvent() {

	rd := receiver.l2CacheAdapter.(*redisCache)

	pubSub := rd.GetKeySetEventPubSub()

	for {
		zap.L().Info("wait...")
		msg := <-pubSub.Channel()
		zap.L().Info("收到set消息\n", zap.Any("Pattern", msg.Pattern), zap.Any("payload", msg.Payload), zap.Any("Payload slice", msg.PayloadSlice))
		receiver.l1CacheAdapter.Delete(msg.Payload)

	}

}

func (receiver *l2Cache) handleL2CacheDelKeyEvent() {

	rd := receiver.l2CacheAdapter.(*redisCache)

	pubSub := rd.GetKeyDelEventPubSub()

	for {
		zap.L().Info("wait...")
		msg := <-pubSub.Channel()
		zap.L().Info("收到del消息\n", zap.Any("Pattern", msg.Pattern), zap.Any("payload", msg.Payload), zap.Any("Payload slice", msg.PayloadSlice))
		receiver.l1CacheAdapter.Delete(msg.Payload)

	}

}
