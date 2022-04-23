package context

import (
	"github.com/gin-gonic/gin"
	"time"
)

type GtContext struct {
}

func NewGtContext(ctx *gin.Context) {

}

func (*GtContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*GtContext) Done() <-chan struct{} {
	return nil
}

func (*GtContext) Err() error {
	return nil
}

func (*GtContext) Value(key interface{}) interface{} {
	return nil
}
