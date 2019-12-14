package test

import (
	"aiTest/leetcode"
	"testing"
)

func TestIsValidSudoku0001(t *testing.T) {
	var source = [][]byte{
		{'8','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	var result = false

	s := leetcode.IsValidSudoku(source)


	if s != result {
		t.Error("Expected ", result, ", got ", s)
	}
}