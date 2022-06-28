package models

import (
	"sort"

	"students/skiplist"
)

// Get finds a node from skip list
func Get(id int) (map[string]string, error) {
	node, err := skiplist.Get(id)
	if err != nil {
		return make(map[string]string), err
	}

	return node.Info, nil
}

// List lists all nodes from skip list
func List(key, value, sortKey string) []map[string]string {
	studentList := skiplist.List()
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
			return result[i][sortKey] < result[j][sortKey]
		})

		return result
	}

	result = append(result, studentList...)
	sort.Slice(result, func(i, j int) bool {
		return result[i]["id"] < result[j]["id"]
	})

	return result
}
