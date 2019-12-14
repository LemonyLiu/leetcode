package test

import (
	"aiTest/leetcode"
	"math"
	"testing"
)

func TestReverseInteger0001(test *testing.T){
	s:=leetcode.Reverse(123)
	if s!=321{
		test.Error("Expected 321, got ", s)
	}
}

func TestReverseInteger0002(test *testing.T){
	s:=leetcode.Reverse(1230)
	if s!=321{
		test.Error("Expected 321, got ", s)
	}
}

func TestReverseInteger0003(test *testing.T){
	source := 1024
	result := 4201
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

func TestReverseInteger0004(test *testing.T){
	source := math.MaxInt32
	result := 0
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

func TestReverseInteger0005(test *testing.T){
	source := -452
	result := -254
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

func TestReverseInteger0006(test *testing.T){
	source := -1
	result := -1
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

func TestReverseInteger0007(test *testing.T){
	source := 9
	result := 9
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

func TestReverseInteger0008(test *testing.T){
	source := 0
	result := 0
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

func TestReverseInteger0009(test *testing.T){
	source := 10
	result := 1
	s:=leetcode.Reverse(source)
	if s!=result{
		test.Error("Expected ", result, ", got ", s)
	}
}

