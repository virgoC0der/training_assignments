package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"students/lib"
)

func TestUpdateByID(t *testing.T) {
	New()
	Add(1, "name", "abc")
	Convey("update by id", t, func() {
		Convey("success", func() {

			type update struct {
				id  int
				kvs []string
			}

			type except struct {
				err error
			}

			cases := []update{
				{
					id:  1,
					kvs: []string{"name", "billy"},
				},
				{
					id:  1,
					kvs: []string{"name", "billy1", "department", "xcentral"},
				},
				{
					id:  3,
					kvs: []string{"name", "billy"},
				},
				{
					id:  1,
					kvs: []string{},
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
					err: lib.ErrIDNotFound,
				},
				{
					err: nil,
				},
			}
			for i, c := range cases {
				err := UpdateByID(c.id, c.kvs...)
				So(err, ShouldEqual, excepts[i].err)
			}
		})
	})
}

func TestUpdateByName(t *testing.T) {
	New()
	Add(1, "name", "abc")
	Convey("update member by name", t, func() {
		Convey("success", func() {
			type update struct {
				name string
				kvs  []string
			}

			type except struct {
				err error
			}

			cases := []update{
				{
					name: "abc",
					kvs:  []string{"department", "xcentral", "date", "2022-01-01"},
				},
				{
					name: "abc",
					kvs:  []string{},
				},
				{
					name: "asdf",
					kvs:  []string{"department", "xcentral", "date", "2022-01-01"},
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
					err: lib.ErrNameNotFound,
				},
			}

			for i, c := range cases {
				err := UpdateByName(c.name, c.kvs...)
				So(err, ShouldEqual, excepts[i].err)
			}
		})
	})
}
