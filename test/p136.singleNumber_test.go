package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestSingleNumber0001(t *testing.T) {
	var source = []int{1,2,1}
	var result = 2

	s := leetcode.SingleNumber(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}


func TestSingleNumber0002(t *testing.T) {
	var source = []int{4,1,2,1,2}
	var result = 4

	s := leetcode.SingleNumber(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}
