package test

import (
	"aiTest/leetcode"
	"fmt"
	"testing"
)

func TestMoveZeroes0001(t *testing.T) {
	var source = []int{1,2,0,1}
	var result = []int{1,2,1,0}

	leetcode.MoveZeroesSolutionTwo(source)

	sStr := fmt.Sprint(source)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", source)
	}
}

func TestMoveZeroes0002(t *testing.T) {
	var source = []int{1,0,0,1}
	var result = []int{1,1,0,0}

	leetcode.MoveZeroesSolutionTwo(source)

	sStr := fmt.Sprint(source)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", source)
	}
}

func TestMoveZeroes0003(t *testing.T) {
	var source = []int{0,0,0,1}
	var result = []int{1,0,0,0}

	leetcode.MoveZeroesSolutionTwo(source)

	sStr := fmt.Sprint(source)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", source)
	}
}

func TestMoveZeroes0004(t *testing.T) {
	var source = []int{0,0,0,0}
	var result = []int{0,0,0,0}

	leetcode.MoveZeroesSolutionTwo(source)

	sStr := fmt.Sprint(source)
	resultStr := fmt.Sprint(result)

	if sStr != resultStr {
		t.Error("Expected ", result, ", got ", source)
	}
}
