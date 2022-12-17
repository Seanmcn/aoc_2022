package main

import (
	"testing"
)

func TestItemPriority(t *testing.T) {
	want := 1
	value := "a"
	priority := getItemPriority(value)

	if want != priority {
		t.Fatalf(`TestItemPriority('%v') = %v, want %v`, value, priority, want)
	}

	want = 27
	value = "A"
	priority = getItemPriority(value)

	if want != priority {
		t.Fatalf(`TestItemPriority('%v') = %v, want %v`, value, priority, want)
	}
}
