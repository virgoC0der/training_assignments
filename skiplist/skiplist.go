package skiplist

import (
	"math/rand"
	"time"
)

const MaxLevel = 7

type Node struct {
	Value int
	Prev  *Node //
	Next  *Node
	Down  *Node
}

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
	if randNum == 0 {
		return true
	}

	return false
}

func (l *SkipList) Insert(value int, headNodeInsertPosition []*Node) {
	node := &Node{
		Value: value,
	}
	// 调表为空插入最底层
	if l.Level < 0 {
		l.Level = 0
		l.HeadNodeArr[0] = &Node{}
		l.HeadNodeArr[0].Next = node
		node.Prev = l.HeadNodeArr[0]
	} else {
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
				root = &Node{}
				l.HeadNodeArr[currentLevel] = root
			} else {
				root = headNodeInsertPosition[currentLevel]
			}

			next = root.Next
			upNode := &Node{}
			upNode.Value = value
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

func (l *SkipList) Add(value int) {
	if l.Exist(value) != nil {
		return
	}

	headNodeInsertPosition := make([]*Node, 0, MaxLevel)
	if l.Level >= 0 {
		level := l.Level
		node := l.HeadNodeArr[level].Next
		for node != nil && level >= 0 {
			if node.Value > value {
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

			if node.Value < value {
				// if node's value is smaller than value and next node is nil,
				// enter next level
				if node.Next == nil {
					headNodeInsertPosition[level] = node
				}
			} else {
				node = node.Next
			}
		}
	}

	l.Insert(value, headNodeInsertPosition)
}

func (l *SkipList) Exist(value int) *Node {
	// level < 0 represents no data
	if l.Level < 0 {
		return nil
	}

	level := l.Level
	node := l.HeadNodeArr[level].Next
	for node != nil {
		if node.Value == value {
			return node
		}

		if node.Value > value {
			// if node's value is bigger than value, should return last node and enter next level
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
		} else if node.Value < value {
			node = node.Next
			// if node's value is smaller than value and next node is nil,
			// the level has already been searched, enter next level
			if node == nil {
				level--
				if level >= 0 {
					node = l.HeadNodeArr[level].Next
				}
			}
		}
	}

	return nil
}

// Delete deletes the node which matches the value
func (l *SkipList) Delete(value int) {
	node := l.Exist(value)
	if node == nil {
		return
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
}
