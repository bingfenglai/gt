package errors

import "errors"

var (
	ErrCacheKeyCannotBeEmpty = errors.New("缓存健不能为空")
	ErrCacheValueNotFound = errors.New("缓存值不存在")
)