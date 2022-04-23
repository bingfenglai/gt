package test

import (
	"log"
	"testing"

	"github.com/bingfenglai/gt/common/helper"
)

func TestGetoStr(t *testing.T) {
	s1 := "/v1/test/*/print/*"
	s2 := "/v1/test/hello/print"

	s1, s2 = helper.Match(s1, s2)
	log.Default().Println("s1", s1)
	log.Default().Println("s2", s2)

	if s1 != s2 {
		t.Fail()
	}

}
