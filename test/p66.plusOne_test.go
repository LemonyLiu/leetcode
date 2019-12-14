package test

import (
	"aiTest/leetcode"
	"fmt"
	"testing"
)

func TestPlusOne0001(t *testing.T) {
	var source = []int{1,2,2,1}
	var result = []int{1,2,2,2}

	s := leetcode.PlusOne(source)

	sStr := fmt.Sprint(s)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestPlusOne0002(t *testing.T) {
	var source = []int{1,2,2,9}
	var result = []int{1,2,3,0}

	s := leetcode.PlusOne(source)

	sStr := fmt.Sprint(s)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestPlusOne0003(t *testing.T) {
	var source = []int{0}
	var result = []int{1}

	s := leetcode.PlusOne(source)

	sStr := fmt.Sprint(s)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestPlusOne0004(t *testing.T) {
	var source = []int{4,9,9}
	var result = []int{5,0,0}

	s := leetcode.PlusOne(source)

	sStr := fmt.Sprint(s)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", s)
	}
}