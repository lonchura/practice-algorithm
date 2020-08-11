package libs

import (
	. "beego-hw/libs"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return true when a ＝= nil  && b ＝= nil", t, func() {
		So(StringSliceEqual(nil, nil), ShouldBeTrue)
	})

	Convey("TestStringSliceEqual should return false when a ＝= nil  && b != nil", t, func() {
		a := []string(nil)
		b := []string{}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

	Convey("TestStringSliceEqual should return false when a != nil  && b != nil", t, func() {
		a := []string{"hello", "world"}
		b := []string{"hello", "goconvey"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}

func Test2StringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		Convey("should return true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return true when a ＝= nil  && b ＝= nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a ＝= nil  && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a != nil  && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}