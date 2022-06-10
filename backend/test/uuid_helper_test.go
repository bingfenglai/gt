package test

import (
	"github.com/bingfenglai/gt/common/helper"
	"log"
	"testing"
)

func TestUuidGen(t *testing.T) {
	str := helper.GenUUIDStr()
	log.Default().Println(str)
}
