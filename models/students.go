package models

import (
	"fmt"
	"sort"
	"strconv"

	"students/skiplist"
)

var (
	skipList *skiplist.SkipList
)

func New() {
	skipList = skiplist.New()
}

func Add(id int, k, v string) {
	skipList.Add(id, k, v)
}

func Get(id int) map[string]string {
	node := skipList.Get(id)
	if node == nil {
		return make(map[string]string)
	}

	return node.Info
}

func UpdateByID(id int, args ...string) {
	skipList.Update(id, args...)
}

func UpdateByName(name string, args ...string) {
	studentList := skipList.List()
	var id string
	for _, student := range studentList {
		if student["name"] == name {
			id = student["id"]
		}
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("id[%s] atoi err\n", id)
		return
	}

	skipList.Update(idInt, args...)
}

func Delete(id int) {
	skipList.Delete(id)
}

func List(key, value string) []map[string]string {
	studentList := skipList.List()
	result := make([]map[string]string, 0, len(studentList))
	if len(key) > 0 {
		for _, student := range studentList {
			v, ok := student[key]
			if !ok {
				continue
			}

			if v == value {
				result = append(result, student)
			}
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i][key] < result[j][key]
		})

		return result
	}

	for _, student := range studentList {
		result = append(result, student)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i]["id"] < result[j]["id"]
	})

	return result
}
