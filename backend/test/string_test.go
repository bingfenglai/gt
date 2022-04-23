package test

import (
	"fmt"
	"log"
	"testing"
)

func TestRepalce(t *testing.T) {

	s := "hello %s Welcome to golang"

	s = fmt.Sprintf(s, "韩立")

	log.Default().Println(s)
}
