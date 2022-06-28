package skiplist

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInit(t *testing.T) {
	Convey("test new a skip list", t, func() {
		Convey("success", func() {
			type ifRun struct {
				run bool
			}
			type result struct {
				sl *SkipList
			}
			cases := []ifRun{
				{
					run: false,
				},
				{
					run: true,
				},
			}
			results := []result{
				{
					sl: nil,
				},
				{
					sl: &SkipList{
						Level:       -1,
						HeadNodeArr: make([]*Node, MaxLevel),
					},
				},
			}

			for i, c := range cases {
				skipList = nil
				if c.run {
					New()
				}
				So(skipList, ShouldResemble, results[i].sl)
			}
		})
	})
}
