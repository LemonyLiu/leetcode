package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestRemoveDuplicates0001(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 5}
	result := 5

	s := leetcode.RemoveDuplicates(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}


func TestRemoveDuplicates0002(t *testing.T) {
	source := []int{1}
	result := 1

	s := leetcode.RemoveDuplicates(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRemoveDuplicates0003(t *testing.T) {
	source := []int{1, 1, 2}
	result := 2

	s := leetcode.RemoveDuplicates(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRemoveDuplicates0004(t *testing.T) {
	source := []int{1, 2, 3}
	result := 3

	s := leetcode.RemoveDuplicates(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRemoveDuplicates0005(t *testing.T) {
	source := []int{1, 2, 3, 3}
	result := 3

	s := leetcode.RemoveDuplicates(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

