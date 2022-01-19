package test

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/bingfenglai/gt/models/shortcodegen"
	_ "github.com/bingfenglai/gt/routers"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGenShortCode(t *testing.T) {

	generator := shortcodegen.NewMd5ShortCodeGenerator()
	method := generator.GetGenMethod()
	log.Default().Println(method)

	gen, err := shortcodegen.GetShortCodeGeneratorByMethod(shortcodegen.CryptoRoundGen)

	if err != nil {
		logs.Error(err.Error())
	} else {
		code, _ := gen.GenShortCode("https://www.baidu.com")
		logs.Info(code)
	}
}
