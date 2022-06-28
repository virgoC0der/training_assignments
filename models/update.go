package models

import (
	"fmt"
	"strconv"
	"students/lib"
	"students/skiplist"
)

// UpdateByID updates a node by id
func UpdateByID(id int, args ...string) error {
	return skiplist.Update(id, args...)
}

// UpdateByName updates a node by name
func UpdateByName(name string, args ...string) error {
	studentList := skiplist.List()
	var id string
	for _, student := range studentList {
		if student["name"] == name {
			id = student["id"]
		}
	}

	if len(id) == 0 {
		return lib.ErrNameNotFound
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("id[%s] atoi err\n", id)
		return err
	}

	return skiplist.Update(idInt, args...)
}
