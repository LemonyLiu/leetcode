package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestRomanToInt0001(t *testing.T) {
	source := "V"
	result := 5

	s := leetcode.RomanToInt(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}


func TestRomanToInt0002(t *testing.T) {
	source := "III"
	result := 3

	s := leetcode.RomanToInt(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRomanToInt0003(t *testing.T) {
	source := "IV"
	result := 4

	s := leetcode.RomanToInt(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRomanToInt0004(t *testing.T) {
	source := "IX"
	result := 9

	s := leetcode.RomanToInt(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRomanToInt0005(t *testing.T) {
	source := "LVIII"
	result := 58

	s := leetcode.RomanToInt(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}

func TestRomanToInt0006(t *testing.T) {
	source := "MCMXCIV"
	result := 1994

	s := leetcode.RomanToInt(source)

	if s != result{
		t.Error("Expected ", result, ", got ", s)
	}
}