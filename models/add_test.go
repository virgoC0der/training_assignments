package models

import (
	"students/skiplist"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"students/lib"
)

func TestAdd(t *testing.T) {
	skiplist.Init()
	type info struct {
		id   int
		name string
	}

	type except struct {
		err error
	}

	Convey("add a new student", t, func() {
		Convey("success", func() {
			cases := []info{
				{
					id:   1,
					name: "abc",
				},
				{
					id:   2,
					name: "abc",
				},
				{
					id:   1,
					name: "aaa",
				},
			}
			excepts := []except{
				{
					err: nil,
				},
				{
					err: nil,
				},
				{
					err: lib.ErrIDExist,
				},
			}

			for i, c := range cases {
				err := Add(c.id, "name", c.name)
				So(err, ShouldEqual, excepts[i].err)
			}
		})
	})
}
