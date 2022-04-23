package cache

import (
	"encoding"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/common/model/cache/util"

	"github.com/patrickmn/go-cache"
)

const localStorageFile = "cache.dump"

// 本地内存缓存
type localCache struct {
	cacheAdapter      *cache.Cache
	defaultExpiration time.Duration
}

// 创建一个指定过期时间和清除间隔时间的本地缓存
func newLocalCache(defaultExpiration time.Duration, cleanupInterval time.Duration) *localCache {
	cacheAdapter := cache.New(defaultExpiration, cleanupInterval)

	cacheAdapter.LoadFile(localStorageFile)

	return &localCache{
		cacheAdapter:      cacheAdapter,
		defaultExpiration: defaultExpiration,
	}
}

func (lc *localCache) Set(key string, value interface{}, expiration time.Duration) error {
	ok, errstr := lc.SetWithJson(key, value, expiration)
	if ok {
		return nil
	} else {
		return errors.New(errstr)
	}
}
func (lc *localCache) SetWithDefaultExpiration(key string, value interface{}) error {
	zap.L().Info("local cache  set key: " + key)
	return lc.Set(key, value, lc.defaultExpiration)

}

func (lc *localCache) SetWithJson(key string, value interface{}, expiration time.Duration) (bool, string) {

	if key == "" {
		return false, errors.ErrCacheKeyCannotBeEmpty.Error()
	}

	jsonByte, err := json.Marshal(value)
	if err != nil {
		return false, err.Error()
	}

	lc.cacheAdapter.Set(key, string(jsonByte), expiration)

	defer lc.cacheAdapter.SaveFile(localStorageFile)

	return true, ""
}
func (lc *localCache) SetWithJsonAndDefaultExpiration(key string, value interface{}) (bool, string) {

	return lc.SetWithJson(key, value, lc.defaultExpiration)

}

func (lc *localCache) Get(key string, value interface{}) error {

	ok, str := lc.GetWithJson(key)
	if !ok || str == "null" {
		return errors.New(str)
	} else {
		zap.L().Info("从本地缓存拿到数据", zap.Any(key, str))
	}

	return lc.scan([]byte(str), value)

}

func (lc *localCache) GetWithJson(key string) (bool, string) {
	val, ok := lc.cacheAdapter.Get(key)

	if ok {
		s := val.(string)

		return ok, s
	}

	return false, ""

}
func (lc *localCache) Keys(keyPrefix string) (bool, []string) {
	keys := make([]string, 0)
	for k := range lc.cacheAdapter.Items() {
		if strings.HasPrefix(k, keyPrefix) {
			keys = append(keys, k)
		}
	}

	if len(keys) == 0 {
		return false, nil
	}

	return true, keys

}
func (lc *localCache) Delete(key ...string) (bool, int64) {
	defer lc.cacheAdapter.SaveFile(localStorageFile)
	count := 0
	for _, k := range key {
		lc.cacheAdapter.Delete(k)
		count++
	}
	return true, int64(count)
}

func (lc *localCache) scan(b []byte, v interface{}) error {
	switch v := v.(type) {
	case nil:
		return fmt.Errorf("local-cache: Scan(nil)")
	case *string:
		*v = util.BytesToString(b)
		return nil
	case *[]byte:
		*v = b
		return nil
	case *int:
		var err error
		*v, err = util.Atoi(b)
		return err
	case *int8:
		n, err := util.ParseInt(b, 10, 8)
		if err != nil {
			return err
		}
		*v = int8(n)
		return nil
	case *int16:
		n, err := util.ParseInt(b, 10, 16)
		if err != nil {
			return err
		}
		*v = int16(n)
		return nil
	case *int32:
		n, err := util.ParseInt(b, 10, 32)
		if err != nil {
			return err
		}
		*v = int32(n)
		return nil
	case *int64:
		n, err := util.ParseInt(b, 10, 64)
		if err != nil {
			return err
		}
		*v = n
		return nil
	case *uint:
		n, err := util.ParseUint(b, 10, 64)
		if err != nil {
			return err
		}
		*v = uint(n)
		return nil
	case *uint8:
		n, err := util.ParseUint(b, 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(n)
		return nil
	case *uint16:
		n, err := util.ParseUint(b, 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(n)
		return nil
	case *uint32:
		n, err := util.ParseUint(b, 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(n)
		return nil
	case *uint64:
		n, err := util.ParseUint(b, 10, 64)
		if err != nil {
			return err
		}
		*v = n
		return nil
	case *float32:
		n, err := util.ParseFloat(b, 32)
		if err != nil {
			return err
		}
		*v = float32(n)
		return err
	case *float64:
		var err error
		*v, err = util.ParseFloat(b, 64)
		return err
	case *bool:
		*v = len(b) == 1 && b[0] == '1'
		return nil
	case *time.Time:
		var err error
		*v, err = time.Parse(time.RFC3339Nano, util.BytesToString(b))
		return err
	case encoding.BinaryUnmarshaler:
		return v.UnmarshalBinary(b)
	default:
		return fmt.Errorf(
			"redis: can't unmarshal %T (consider implementing BinaryUnmarshaler)", v)
	}
}
