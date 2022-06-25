package models

import (
	"fmt"
	"strconv"
	"students/lib"
)

func UpdateByID(id int, args ...string) error {
	return skipList.Update(id, args...)
}

func UpdateByName(name string, args ...string) error {
	studentList := skipList.List()
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

	return skipList.Update(idInt, args...)
}
