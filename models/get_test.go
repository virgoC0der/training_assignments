package models

import (
	"students/common"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	New()
	Add(1, "name", "abc")
	Add(2, "name", "abcd")
	Convey("test get a member", t, func() {
		Convey("success", func() {
			type info struct {
				id int
			}

			type except struct {
				result map[string]string
				err    error
			}

			cases := []info{
				{
					id: 1,
				},
				{
					id: 2,
				},
				{
					id: 3,
				},
			}
			excepts := []except{
				{
					result: map[string]string{"name": "abc"},
					err:    nil,
				},
				{
					result: map[string]string{"name": "abcd"},
					err:    nil,
				},
				{
					result: map[string]string{},
					err:    common.ErrIDNotFound,
				},
			}
			for i, c := range cases {
				node, err := Get(c.id)
				So(node, ShouldResemble, excepts[i].result)
				So(err, ShouldEqual, excepts[i].err)
			}
		})
	})
}

func TestList(t *testing.T) {
	Convey("test list members", t, func() {
		Convey("success", func() {
			type info struct {
				id   int
				name string
			}

			type listInput struct {
				list    []info
				key     string
				value   string
				sortKey string
			}

			type except struct {
				length int
			}

			cases := []listInput{
				{
					list: []info{
						{
							id:   1,
							name: "abc",
						},
						{
							id:   2,
							name: "abcd",
						},
					},
					key:     "",
					value:   "",
					sortKey: "",
				},
				{
					list: []info{
						{
							id:   1,
							name: "abc",
						},
						{
							id:   2,
							name: "abcd",
						},
					},
					key:     "name",
					value:   "abc",
					sortKey: "id",
				},
				{
					list:    []info{},
					key:     "name",
					value:   "abc",
					sortKey: "id",
				},
			}
			excepts := []except{
				{length: 2},
				{length: 1},
				{length: 0},
			}
			for i, c := range cases {
				New()
				for _, v := range c.list {
					Add(v.id, "name", v.name)
				}
				result := List(c.key, c.value, c.sortKey)
				So(len(result), ShouldEqual, excepts[i].length)
			}
		})
	})
}
