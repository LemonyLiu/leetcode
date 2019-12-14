package test

import (
	"aiTest/leetcode"
	"fmt"
	"testing"
)

func TestIntersect0001(t *testing.T) {
	var source1 = []int{1,2,2,1}
	var source2 = []int{2,2}
	var result = []int{2,2}

	s := leetcode.Intersect(source1, source2)

	sStr := fmt.Sprint(s)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", s)
	}
}


func TestIntersect0002(t *testing.T) {
	var source1 = []int{4,9,5}
	var source2 = []int{9,4,9,8,4}
	var result = []int{9,4}

	s := leetcode.Intersect(source1, source2)

	sStr := fmt.Sprint(s)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", s)
	}
}