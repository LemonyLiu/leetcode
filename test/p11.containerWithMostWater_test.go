package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestMaxArea0001(t *testing.T){
	source := []int{1,8,6,2,5,4,8,3,7}
	result := 49

	s := leetcode.MaxArea(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestMaxArea0002(t *testing.T){
	source := []int{1,1}
	result := 1

	s := leetcode.MaxArea(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}
