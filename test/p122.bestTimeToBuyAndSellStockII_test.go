package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestMaxProfit0001(t *testing.T) {
	var source = []int{1,2}
	var result = 1

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}


func TestMaxProfit0002(t *testing.T) {
	var source = []int{7,1,5,3,6,4}
	var result = 7

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestMaxProfit0003(t *testing.T) {
	var source = []int{1,2,3,4,5}
	var result = 4

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestMaxProfit0004(t *testing.T) {
	var source = []int{7,6,4,3,1}
	var result = 0

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestMaxProfit0005(t *testing.T) {
	var source = []int{1,2,2,3}
	var result = 2

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestMaxProfit0006(t *testing.T) {
	var source = []int{1,2,2,1}
	var result = 1

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestMaxProfit0007(t *testing.T) {
	var source = []int{2,2,2,2}
	var result = 0

	s := leetcode.MaxProfit(source)
	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}