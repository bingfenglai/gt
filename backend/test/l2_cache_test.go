package test

import (
	"github.com/bingfenglai/gt/service"
	"testing"
)

func TestL2CacheSet(t *testing.T) {

	if err := service.CacheService.SetWithDefaultExpiration("1", "Hello World");err!=nil{
		t.Error(err)
	}


}

func TestL2CacheGet(t *testing.T) {
	s1 := ""
	service.CacheService.Get("1",s1)

}
