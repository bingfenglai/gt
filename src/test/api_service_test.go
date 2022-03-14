package test

import (
	"log"
	"testing"

	"github.com/bingfenglai/gt/storage"
)

func TestGetApisByRoleIds(t *testing.T){


	apis,err:= storage.ApiStorage.GetApisByRoleIds(make([]int, 0))
	log.Default().Println(apis,"\n",err)
}