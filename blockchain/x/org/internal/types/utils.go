package types

import (
	"sort"
	"strings"
)

type Findable interface {
	ElAtIndex(index int) string
	Len() int
}

type Skills []string

func (s Skills) Find(id string) (skill string, found bool) {
	index := s.find(id)
	if index == -1 {
		return skill, false
	}
	return s[index], true
}

func (s Skills) find(id string) int {
	return FindUtil(s, id)
}

func FindUtil(group Findable, el string) int {
	if group.Len() == 0 {
		return -1
	}
	low := 0
	high := group.Len() - 1
	median := 0
	for low <= high {
		median = (low + high) / 2
		switch compare := strings.Compare(group.ElAtIndex(median), el); {
		case compare == 0:
			return median
		case compare == -1:
			low = median + 1
		default:
			high = median - 1
		}
	}
	return -1
}

func (s Skills) ElAtIndex(index int) string { return s[index] }

func (s Skills) String() string {
	var a string
	if len(s) == 0 {
		return ""
	}
	for _, skill := range s {
		a = a + "\n" + skill
	}

	return a[:len(a)-1]
}

func (s Skills) Len() int           { return len(s) }
func (s Skills) Less(i, j int) bool { return strings.Compare(s[i], s[j]) == -1 }
func (s Skills) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var _ sort.Interface = Skills{}

func (Skills Skills) Sort() Skills {
	sort.Sort(Skills)
	return Skills
}
