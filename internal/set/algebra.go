package set

const setAlgebra = `
// Contains tests whether an item is already in the {{.TName}}Set.
func (set {{.TName}}Set) Contains(i {{.TName}}) bool {
	_, found := set[i]
	return found
}

// ContainsAll tests whether many items are all in the {{.TName}}Set.
func (set {{.TName}}Set) ContainsAll(i ...{{.TName}}) bool {
	for _, v := range i {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set {{.TName}}Set) actualSubset(other {{.TName}}Set) bool {
	for item := range set {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// Equals determines if two sets are equal to each other.
// They are considered equal if both are the same size and both have the same items.
func (set {{.TName}}Set) Equals(other {{.TName}}Set) bool {
	return set.Size() == other.Size() && set.actualSubset(other)
}

// IsSubset determines if every item in the other set is in this set.
func (set {{.TName}}Set) IsSubset(other {{.TName}}Set) bool {
	return set.Size() <= other.Size() && set.actualSubset(other)
}

// IsProperSubset determines if every item in the other set is in this set and this set is
// smaller than the other.
func (set {{.TName}}Set) IsProperSubset(other {{.TName}}Set) bool {
	return set.Size() < other.Size() && set.actualSubset(other)
}

// IsSuperset determines if every item of this set is in the other set.
func (set {{.TName}}Set) IsSuperset(other {{.TName}}Set) bool {
	return other.IsSubset(set)
}

// Union returns a new set with all items in both sets.
func (set {{.TName}}Set) Union(other {{.TName}}Set) {{.TName}}Set {
	union := New{{.TName}}Set()
	for item := range set {
		union[item] = struct{}{}
	}
	for item := range other {
		union[item] = struct{}{}
	}
	return union
}

// Intersection returns a new set with items that exist only in both sets.
func (set {{.TName}}Set) Intersection(other {{.TName}}Set) {{.TName}}Set {
	intersection := New{{.TName}}Set()
	// loop over the smaller set
	if set.Size() < other.Size() {
		for item := range set {
			if other.Contains(item) {
				intersection[item] = struct{}{}
			}
		}
	} else {
		for item := range other {
			if set.Contains(item) {
				intersection[item] = struct{}{}
			}
		}
	}
	return intersection
}

// Difference returns a new set with items in the current set but not in the other set
func (set {{.TName}}Set) Difference(other {{.TName}}Set) {{.TName}}Set {
	diffs := New{{.TName}}Set()
	for item := range set {
		if !other.Contains(item) {
			diffs[item] = struct{}{}
		}
	}
	return diffs
}

`
