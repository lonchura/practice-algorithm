package libs

import (
	. "beego-hw/libs"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFuncDemo(t *testing.T) {
	Convey("TestFuncDemo", t, func() {
		var num int
		Convey("for succ", func() {
			stubs := Stub(&num, 150)
			defer stubs.Reset()

			stubs.StubFunc(&Exec,"success", nil)
			/*
			var liLei = `{"name":"LiLei", "age":"21"}`
			stubs.StubFunc(&adapter.Marshal, []byte(liLei), nil)
			*/
			//several So assert
			actual, err := Exec("input")
			So(actual == "success", ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
}
