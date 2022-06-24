package models

import (
	"students/skiplist"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("test new a skip list", t, func() {
		Convey("success", func() {
			type ifRun struct {
				run bool
			}
			type result struct {
				sl *skiplist.SkipList
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
					sl: &skiplist.SkipList{
						Level:       -1,
						HeadNodeArr: make([]*skiplist.Node, skiplist.MaxLevel),
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
