package strs

import "sort"

// SetList is an ordered, distinct set of strings.
type SetList struct {
	size  int
	items []string
	pos   map[string]int
}

// init sets up a new SetList.
func (sl *SetList) init() {
	if sl.pos == nil {
		sl.pos = make(map[string]int)
	}
}

// Equal returns true if both SetLists contain the same items, in the same
// order. See also: Match.
func (sl *SetList) Equal(other SetList) bool {
	return Equal(sl.items, other.items)
}

// Match returns true if both SetLists contain the same items, disregarding
// order. See also: Equal.
func (sl *SetList) Match(other SetList) bool {

	if sl.size != other.size {
		return false
	}

	for _, val := range sl.items {
		if !other.Contains(val) {
			return false
		}
	}

	return true
}

// Append will append the given values to the SetList. If the SetList already
// contains a value, it is not added.
func (sl *SetList) Append(values ...string) {
	sl.init()

	for _, val := range values {
		_, ok := sl.pos[val]
		if ok {
			continue
		}

		sl.items = append(sl.items, val)
		sl.pos[val] = sl.size
		sl.size++
	}
}

// Pull is the inverse of Append, removing the last item (successfully) appended
// and returning it. If there are no items, then the second return value will be
// false.
func (sl *SetList) Pull() (string, bool) {

	if sl.size == 0 {
		return "", false
	}

	sl.size--
	s := sl.items[sl.size]
	sl.items = sl.items[:sl.size]
	delete(sl.pos, s)

	return s, true
}

// Remove will remove the given value from the SetList, if present. Returns true
// if the item was found and removed.
func (sl *SetList) Remove(value string) bool {

	p, ok := sl.pos[value]
	if !ok {
		return false
	}

	sl.items = append(sl.items[:p], sl.items[p+1:]...)
	sl.size--
	delete(sl.pos, value)

	return true
}

// Len returns the number of items in the SetList.
func (sl *SetList) Len() int {
	return sl.size
}

// Get returns the string at index i. Returns "" (empty string) if i is out of
// range.
func (sl *SetList) Get(i int) string {
	if i >= sl.size {
		return ""
	}

	return sl.items[i]
}

// Contains returns true/false to indicate if the value is in the SetList.
func (sl *SetList) Contains(value string) bool {

	_, ok := sl.pos[value]
	return ok
}

// Items returns the items contained in the SetList as a []string.
func (sl *SetList) Items() []string {
	return append([]string(nil), sl.items...)
}

// Sorted returns a new SetList with the items in order.
func (sl *SetList) Sorted() *SetList {

	result := &SetList{}
	result.init()

	result.items = append(result.items, sl.items...)
	sort.Strings(result.items)

	for i, s := range result.items {
		result.pos[s] = i
	}

	return result
}
