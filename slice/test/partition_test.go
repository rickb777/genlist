package main

import (
	"testing"
	"reflect"
)

func TestPartition(t *testing.T) {
	things := ThingSlice{
		{"First", 0},
		{"Second", 0},
		{"Third", 0},
		{"Second", 10},
	}

	m1, o1 := things.Partition(func(x Thing) bool {
		return x.Name == "Second"
	})

	expected1 := ThingSlice{
		{"Second", 0},
		{"Second", 10},
	}

	others1 := ThingSlice{
		{"First", 0},
		{"Third", 0},
	}

	if !reflect.DeepEqual(m1, expected1) {
		t.Errorf("Partition should result in %v, got %v", expected1, m1)
	}

	if !reflect.DeepEqual(o1, others1) {
		t.Errorf("Partition should result in %v, got %v", others1, o1)
	}

	m2, o2 := things.Partition(func(x Thing) bool {
		return x.Name == "Dummy"
	})

	if len(m2) != 0 {
		t.Errorf("Partition should result in empty slice, got %v", m2)
	}

	if !reflect.DeepEqual(o2, things) {
		t.Errorf("Partition should result in %v, got %v", things, o2)
	}

	m3, o3 := ThingSlice{}.Partition(func(x Thing) bool {
		return true
	})

	if len(m3) != 0 {
		t.Errorf("Partition should result in empty slice, got %v", m3)
	}

	if len(o3) != 0 {
		t.Errorf("Partition should result in empty slice, got %v", o3)
	}
}
