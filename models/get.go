package models

import "sort"

func Get(id int) (map[string]string, error) {
	node, err := skipList.Get(id)
	if err != nil {
		return make(map[string]string), err
	}

	return node.Info, nil
}

func List(key, value, sortKey string) []map[string]string {
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
