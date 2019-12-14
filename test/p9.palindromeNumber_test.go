package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestIsPalindrome0001(t *testing.T) {
	source := 1
	result := true

	s := leetcode.IsPalindrome(source)
	if s != result{
		t.Error("Expected true, got ", s)
	}
}

func TestIsPalindrome0002(t *testing.T) {
	source := 10
	result := false

	s := leetcode.IsPalindrome(source)
	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestIsPalindrome0003(t *testing.T) {
	source := 101
	result := true

	s := leetcode.IsPalindrome(source)
	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestIsPalindrome0004(t *testing.T) {
	source := -1
	result := false

	s := leetcode.IsPalindrome(source)
	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}
