package sparsebitset

import "sort"

type By func(p1, p2 *KeyValueGroup) bool

type KeyValueGroupSorter struct {
	keyValues []KeyValueGroup
	by        func(p1, p2 *KeyValueGroup) bool // Closure used in the Less method.
}

func (s *KeyValueGroupSorter) Len() int {
	return len(s.keyValues)
}

func (s *KeyValueGroupSorter) Less(i, j int) bool {
	return s.by(&s.keyValues[i], &s.keyValues[j])
}

func (s *KeyValueGroupSorter) Swap(i, j int) {
	s.keyValues[i], s.keyValues[j] = s.keyValues[j], s.keyValues[i]
}

func (by By) Sort(keyValues []KeyValueGroup) {
	ps := &KeyValueGroupSorter{
		keyValues: keyValues,
		by:        by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}
