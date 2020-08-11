package libs

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func ShouldSummerBeComming(actual interface{}, expected ...interface{}) string {
	if actual == "summer" && expected[0] == "comming" {
		return ""
	} else {
		return "summer is not comming!"
	}
}

func TestSummer(t *testing.T) {
	Convey("TestSummer", t, func() {
		So("summer", ShouldSummerBeComming, "comming")
		SkipSo("autumn", ShouldSummerBeComming, "comming")
		So("winter", ShouldSummerBeComming, "comming")
	})
}