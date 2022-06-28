package skiplist

var skipList *SkipList

func Init() {
	skipList = New()
}

// Add adds a new node to skip list
func Add(id int, k, v string) error {
	return skipList.Add(id, k, v)
}

// Update updates node from skip list
func Update(id int, args ...string) error {
	return skipList.Update(id, args...)
}

// Get finds a node from skip list by id
func Get(id int) (*Node, error) {
	return skipList.Get(id)
}

// List returns all nodes from skip list
func List() []map[string]string {
	return skipList.List()
}

// Delete removes a node from skip list by id
func Delete(id int) error {
	return skipList.Delete(id)
}
