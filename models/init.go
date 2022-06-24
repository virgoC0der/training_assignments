package models

import (
	"students/skiplist"
)

var (
	skipList *skiplist.SkipList
)

func New() {
	skipList = skiplist.New()
}
