package paramspeed

import (
	"testing"

	"github.com/codemodus/parth"
	"github.com/julienschmidt/httprouter"
)

var (
	gstr string
)

func BenchmarkSimpleParthSubSeg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v, err := parth.SubSegToString("/1/class/12", "class")
		if err != nil {
			b.Fatal(err)
		}
		gstr = v
	}
}

func BenchmarkComplexParth(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pp := parth.New("/1/class/12/string/object/34")

		v := pp.SubSegToString("class")
		gstr = v

		v = pp.SegmentToString(3)
		gstr = v

		v = pp.SegmentToString(-1)
		gstr = v

		if pp.Err() != nil {
			b.Fatal(pp.Err())
		}
	}
}

func BenchmarkSimpleParthFromFront(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v, err := parth.SegmentToString("/1/class/12", 2)
		if err != nil {
			b.Fatal(err)
		}
		gstr = v
	}
}

func BenchmarkSimpleParthFromEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v, err := parth.SegmentToString("/1/class/12", -1)
		if err != nil {
			b.Fatal(err)
		}
		gstr = v
	}
}

func BenchmarkSimpleHTTPRouterParams(b *testing.B) {
	ps := make(httprouter.Params, 1)
	ps[0] = httprouter.Param{Key: "classID", Value: "12"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		v := ps.ByName("classID")
		gstr = v
	}
}

func BenchmarkComplexHTTPRouterParams(b *testing.B) {
	ps := make(httprouter.Params, 3)
	ps[0] = httprouter.Param{Key: "classID", Value: "12"}
	ps[1] = httprouter.Param{Key: "alt", Value: "string"}
	ps[2] = httprouter.Param{Key: "objectID", Value: "34"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		v := ps.ByName("classID")
		gstr = v

		v = ps.ByName("alt")
		gstr = v

		v = ps.ByName("objectID")
		gstr = v
	}
}
