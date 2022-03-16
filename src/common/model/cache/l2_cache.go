package cache

import "time"

// 二级缓存实现，本地缓存未查询到的信息，将从远程redis缓存当中查询，仍然查询不到则返回ErrCacheKeyNotFound,查询到则进行回写
type l2Cache struct {
	l1CacheAdapter Cache
	l2CacheAdapter Cache
}

func newL2Cache(l1CacheAdapter, l2CacheAdapter Cache) *l2Cache {
	return &l2Cache{
		l1CacheAdapter: l1CacheAdapter,
		l2CacheAdapter: l2CacheAdapter,
	}
}

func (receiver *l2Cache) Set(key string, value interface{}, expiration time.Duration) error{
	return nil
}
func (receiver *l2Cache) SetWithDefaultExpiration(key string, value interface{}) error{
	return nil
	
}
func (receiver *l2Cache) Get(key string, value interface{}) error{
	return nil

}
func (receiver *l2Cache) SetWithJson(key string, value interface{}, expiration time.Duration) (bool, string){
	return false,""

}
func (receiver *l2Cache) SetWithJsonAndDefaultExpiration(key string, value interface{}) (bool, string){
	return false,""

}
func (receiver *l2Cache) GetWithJson(key string) (bool, string){
	return false,""

}
func (receiver *l2Cache) Keys(keyPrefix string) (bool, []string){
	return false,nil
}
func (receiver *l2Cache) Delete(key ...string) (bool, int64){
	return false,-1

}
