package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestContainsDuplicate0001(t *testing.T) {
	var source = []int{1,2}
	var result = false

	s := leetcode.ContainsDuplicate(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestContainsDuplicate0002(t *testing.T) {
	var source = []int{1,2,2}
	var result = true

	s := leetcode.ContainsDuplicate(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}
