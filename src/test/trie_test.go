package test

import (
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/bingfenglai/gt/config"
	v2 "github.com/bingfenglai/gt/model/permission/v2"
)

func TestTrie(t *testing.T) {

	mt := make(v2.MethodTrie)

	root := mt.GetRoot(http.MethodGet)

	tn := root.Root()

	paths := strings.Split("gt/model/permission/*", "/")
	tn.InsertChild(paths)

	paths1 := strings.Split("gt/model/permission/:v2", "/")
	paths3 := strings.Split("gt/model/permission/:v2/wee", "/")

	//flag := tn.Search(paths)
	flag1 := tn.Search(paths1)
	flag3 := tn.Search(paths3)

	log.Default().Println(flag1, flag3)
}

func TestTrieInsertChildBatch(t *testing.T) {

	mt := make(v2.MethodTrie)

	root := mt.GetRoot(http.MethodGet)

	tn := root.Root()

	tn.InsertChildBatch(config.Conf.Auth.AnonymousUrls)

	flag := tn.Search(strings.Split("v1/test/name", "/"))

	log.Default().Println(flag)

}
