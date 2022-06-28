package models

import "students/skiplist"

// Add adds a new node to skip list
func Add(id int, k, v string) error {
	return skiplist.Add(id, k, v)
}
