package models

import "students/skiplist"

// Delete removes a node from skip list
func Delete(id int) error {
	return skiplist.Delete(id)
}
