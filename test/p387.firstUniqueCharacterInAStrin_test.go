package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestFirstUniqChar0001(t *testing.T) {
	var source = "test"
	var result = 1

	s := leetcode.FirstUniqChar2(source)

	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}


func TestFirstUniqChar0002(t *testing.T) {
	var source = "loveleetcode"
	var result = 2

	s := leetcode.FirstUniqChar2(source)

	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}