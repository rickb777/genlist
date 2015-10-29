package main

import (
	"reflect"
	"testing"
)

func TestSortBy(t *testing.T) {
	first := Thing{"First", 60}
	second := Thing{"Second", 40}
	third := Thing{"Third", 100}
	anotherThird := Thing{"Third", 100}
	fourth := Thing{"Fourth", 40}
	fifth := Thing{"Fifth", 70}
	sixth := Thing{"Sixth", 10}
	seventh := Thing{"Seventh", 50}
	eighth := Thing{"Eighth", 110}

	things := thingCollection(
		first,
		second,
		third,
		anotherThird,
		fourth,
	)

	lotsOfThings := thingCollection(
		first,
		second,
		third,
		fourth,
		fifth,
		sixth,
		seventh,
		eighth,
	)

	name := func(a, b Thing) bool {
		return a.Name < b.Name
	}

	sort1 := things.SortWith(name)

	sorted1 := ThingList{first, fourth, second, third, anotherThird}

	if !reflect.DeepEqual(sort1, ThingList{first, fourth, second, third, anotherThird}) {
		t.Errorf("SortWith name should be %#v, got %#v", sorted1, sort1)
	}

	if !sort1.IsSortedWith(name) {
		t.Errorf("IsSortedWith name should be true")
	}

	if things.IsSortedWith(name) {
		t.Errorf("things should not be sorted by name")
	}

	sort2 := things.SortWithDesc(name)

	sorted2 := ThingList{anotherThird, third, second, fourth, first}

	if !reflect.DeepEqual(sort2, sorted2) {
		t.Errorf("SortWithDesc name should be %#v, got %#v", sorted2, sort2)
	}

	if !sort2.IsSortedWithDesc(name) {
		t.Errorf("IsSortedWithDesc name should be true %#v", sort2)
	}

	if things.IsSortedWithDesc(name) {
		t.Errorf("things should not be sorted desc by name")
	}

	// intended to hit threshold to invoke quicksort (7)
	sort3 := lotsOfThings.SortWith(name)

	sorted3 := ThingList{eighth, fifth, first, fourth, second, seventh, sixth, third}

	if !reflect.DeepEqual(sort3, sorted3) {
		t.Errorf("Sort name should be %#v, got %#v", sorted3, sort3)
	}

	// intended to hit threshold to invoke medianOfThree (40)
	var evenMore ThingList
	evenMore = append(evenMore, lotsOfThings...)
	evenMore = append(evenMore, lotsOfThings...)
	evenMore = append(evenMore, lotsOfThings...)
	evenMore = append(evenMore, lotsOfThings...)
	evenMore = append(evenMore, lotsOfThings...)
	evenMore = append(evenMore, lotsOfThings...)

	sort4 := evenMore.SortWith(name)

	sorted4 := ThingList{eighth, eighth, eighth, eighth, eighth, eighth}
	sorted4 = append(sorted4, appendMany(fifth, 6)...)
	sorted4 = append(sorted4, appendMany(first, 6)...)
	sorted4 = append(sorted4, appendMany(fourth, 6)...)
	sorted4 = append(sorted4, appendMany(second, 6)...)
	sorted4 = append(sorted4, appendMany(seventh, 6)...)
	sorted4 = append(sorted4, appendMany(sixth, 6)...)
	sorted4 = append(sorted4, appendMany(third, 6)...)

	if !reflect.DeepEqual(sort4, sorted4) {
		t.Errorf("Sort name should be %#v, got %#v", sorted3, sort3)
	}
}

func appendMany(x Thing, n int) (result ThingList) {
	for i := 0; i < n; i++ {
		result = append(result, x)
	}
	return
}
