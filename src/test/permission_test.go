package test

import (
	"log"
	"strings"
	"testing"

	"github.com/bingfenglai/gt/model/permission"
)

func TestBuildNode(t *testing.T) {

	pts := strings.Split("v1/test/:name/:age/print", "/")
	
	log.Default().Println(pts)

	node, p := permission.BuildNodeLink(pts, nil)
	
	log.Default().Println(node,"\n",p)

	paths := make([]string, 0)
	_,paths = permission.Iterate(node, paths)


	log.Default().Println(paths)

	flag := permission.IsAuthoried(node,"v1/test/:name/:age/print")

	

	log.Default().Println(flag)

}
