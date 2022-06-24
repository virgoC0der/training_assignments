package models

import (
	"students/common"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDelete(t *testing.T) {
	New()
	Add(1, "name", "abc")
	Convey("delete student", t, func() {
		Convey("success", func() {
			type info struct {
				id int
			}

			type except struct {
				err error
			}

			cases := []info{
				{
					id: 1,
				},
				{
					id: 2,
				},
			}
			excepts := []except{
				{
					err: nil,
				},
				{
					err: common.ErrIDNotFound,
				},
			}

			for i, c := range cases {
				err := Delete(c.id)
				So(err, ShouldEqual, excepts[i].err)
			}
		})
	})
}
