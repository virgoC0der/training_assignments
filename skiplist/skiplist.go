package skiplist

import (
	"fmt"
	"math/rand"
	"strconv"
	"students/lib"
	"time"
)

// MaxLevel skip list's max level
const MaxLevel = 7

// Node skip list's node
type Node struct {
	ID   int
	Name string
	Info map[string]string
	Prev *Node
	Next *Node
	Down *Node
}

// SkipList skip list data
type SkipList struct {
	Level       int
	HeadNodeArr []*Node
}

// New inits a new skip list
func New() *SkipList {
	list := &SkipList{
		Level:       -1,
		HeadNodeArr: make([]*Node, MaxLevel),
	}

	rand.Seed(time.Now().UnixNano())
	return list
}

// randLevel decides whether to join the next level by flipping a coin
func randLevel() bool {
	randNum := rand.Intn(2)
	return randNum == 0
}

// Insert inserts new id into skip_list
func (l *SkipList) Insert(id int, headNodeInsertPosition []*Node, k, v string) {
	node := &Node{
		ID:   id,
		Info: make(map[string]string),
	}
	node.Info[k] = v
	// insert into the bottom level when the skip_list is empty
	if l.Level < 0 {
		l.Level = 0
		l.HeadNodeArr[0] = &Node{Info: make(map[string]string)}
		l.HeadNodeArr[0].Next = node
		node.Prev = l.HeadNodeArr[0]
	} else {
		// if not empty, insert into every level
		// insert into the bottom level
		root := headNodeInsertPosition[0]
		next := root.Next

		root.Next = node
		node.Prev = root
		node.Next = next
		if next != nil {
			next.Prev = node
		}

		currentLevel := 1
		for randLevel() && currentLevel <= l.Level+1 && currentLevel < MaxLevel {
			if headNodeInsertPosition[currentLevel] == nil {
				root = &Node{Info: make(map[string]string)}
				l.HeadNodeArr[currentLevel] = root
			} else {
				root = headNodeInsertPosition[currentLevel]
			}

			next = root.Next
			upNode := &Node{Info: node.Info}
			upNode.ID = id
			upNode.Down = node
			upNode.Prev = root
			upNode.Next = next

			root.Next = upNode
			if next != nil {
				next.Prev = node
			}

			node = upNode
			currentLevel++
		}

		l.Level = currentLevel - 1
	}
}

// Add adds new node to skip_list
func (l *SkipList) Add(id int, k, v string) error {
	node, err := l.Get(id)
	if err == nil && node != nil {
		fmt.Println("id already exists!")
		return lib.ErrIDExist
	}

	headNodeInsertPosition := make([]*Node, MaxLevel)
	if l.Level >= 0 {
		level := l.Level
		node := l.HeadNodeArr[level].Next
		for node != nil && level >= 0 {
			if node.ID > id {
				headNodeInsertPosition[level] = node.Prev
				if node.Prev.Down == nil {
					if level-1 >= 0 {
						node = l.HeadNodeArr[level-1].Next
					} else {
						node = nil
					}
				} else {
					node = node.Prev.Down
				}

				level--
				continue
			}

			if node.ID < id {
				// if node's value is smaller than value and next node is nil,
				// enter next level
				if node.Next == nil {
					headNodeInsertPosition[level] = node
					level--
					if level >= 0 {
						node = l.HeadNodeArr[level].Next
					}
				} else {
					node = node.Next
				}
			}
		}
	}

	l.Insert(id, headNodeInsertPosition, k, v)
	return nil
}

func (l *SkipList) Update(id int, args ...string) error {
	node, err := l.Get(id)
	if err != nil {
		fmt.Println("id does not exist")
		return err
	}

	for i := 0; i < len(args); i += 2 {
		node.Info[args[i]] = args[i+1]
	}

	return nil
}

// Get judges whether the id exists
func (l *SkipList) Get(id int) (*Node, error) {
	// level < 0 represents no data
	if l.Level < 0 {
		return nil, lib.ErrIDNotFound
	}

	level := l.Level
	node := l.HeadNodeArr[level].Next
	for node != nil {
		if node.ID == id {
			return node, nil
		}

		if node.ID > id {
			// if node's id is bigger than id, should return last node and enter next level
			if node.Prev.Down == nil {
				if level-1 >= 0 {
					node = l.HeadNodeArr[level-1].Next
				} else {
					node = nil
				}
			} else {
				node = node.Prev.Down
			}

			level--
		} else if node.ID < id {
			node = node.Next
			// if node's id is smaller than id and next node is nil,
			// the level has already been searched, enter next level
			if node == nil {
				level--
				if level >= 0 {
					node = l.HeadNodeArr[level].Next
				}
			}
		}
	}

	return nil, lib.ErrIDNotFound
}

// Delete deletes the node which matches the id
func (l *SkipList) Delete(id int) error {
	node, err := l.Get(id)
	if err != nil {
		return err
	}

	for node != nil {
		prev := node.Prev
		next := node.Next
		prev.Next = next
		if next != nil {
			next.Prev = prev
		}
		node = node.Down
	}

	return nil
}

// List return all nodes' info
func (l *SkipList) List() []map[string]string {
	if l.HeadNodeArr[0] == nil {
		return []map[string]string{}
	}

	node := l.HeadNodeArr[0].Next

	result := make([]map[string]string, 0)
	for node != nil {
		node.Info["id"] = strconv.Itoa(node.ID)
		result = append(result, node.Info)
		node = node.Next
	}

	return result
}
