package main

import (
	"testing"
	"reflect"
)

func TestPartitionNum(t *testing.T) {
	things := num1Collection(60, -20, -10, 100)

	m1, o1 := things.Partition(func(x Num1) bool {
		return x > 0
	})

	expected1 := num1Collection(60, 100)

	others1 := num1Collection(-20, -10)

	if !reflect.DeepEqual(m1, expected1) {
		t.Errorf("Partition should result in %#v, got %#v", expected1, m1)
	}

	if !reflect.DeepEqual(o1, others1) {
		t.Errorf("Partition should result in %#v, got %#v", others1, o1)
	}

	m2, o2 := things.Partition(func(x Num1) bool {
		return x == 1
	})

	if m2.Size() != 0 {
		t.Errorf("Partition should result in empty list, got %#v", m2)
	}

	if !reflect.DeepEqual(o2, things) {
		t.Errorf("Partition should result in %#v, got %#v", things, o2)
	}

	m3, o3 := num1Collection().Partition(func(x Num1) bool {
		return true
	})

	if m3.Size() != 0 {
		t.Errorf("Partition should result in empty list, got %#v", m3)
	}

	if o3.Size() != 0 {
		t.Errorf("Partition should result in empty list, got %#v", o3)
	}
}

func TestPartitionThing(t *testing.T) {
	things := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)

	m1, o1 := things.Partition(func(x Thing) bool {
		return x.Name[0] == 'F'
	})

	expected1 := thingCollection(
		Thing{"Fee", 1},
		Thing{"Fie", 2},
		Thing{"Foe", 3},
	)


	others1 := thingCollection(
		Thing{"Boo", 5},
		Thing{"Boo", 8},
		Thing{"Bam", 13},
		Thing{"Bam", 21},
	)


	if !reflect.DeepEqual(m1, expected1) {
		t.Errorf("Partition should result in %#v, got %#v", expected1, m1)
	}

	if !reflect.DeepEqual(o1, others1) {
		t.Errorf("Partition should result in %#v, got %#v", others1, o1)
	}

	m2, o2 := things.Partition(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if m2.Size() != 0 {
		t.Errorf("Partition should result in empty list, got %#v", m2)
	}

	if !reflect.DeepEqual(o2, things) {
		t.Errorf("Partition should result in %#v, got %#v", things, o2)
	}

	m3, o3 := ThingList{}.Partition(func(x Thing) bool {
		return true
	})

	if m3.Size() != 0 {
		t.Errorf("Partition should result in empty list, got %#v", m3)
	}

	if o3.Size() != 0 {
		t.Errorf("Partition should result in empty list, got %#v", o3)
	}
}
