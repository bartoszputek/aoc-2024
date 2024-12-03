package main

import (
	"testing"
)

type TestCaseMatchMul struct {
	Value   string
	IsValid bool
	Want    Result
}

func TestMatchMul(t *testing.T) {
	var testCases = []TestCaseMatchMul{
		{"mul(2,4)", true, Result{2, 4}},
		{"mul(34,1234)", true, Result{34, 1234}},
		{"mudl(34,1234)", false, Result{0, 0}},
		{"mul(,1234)", false, Result{0, 0}},
		{"mul(1234,)", false, Result{0, 0}},
		{"mul(,)", false, Result{0, 0}},
		{"mul(  2,4)", false, Result{0, 0}},
		{"mul(2  ,4)", false, Result{0, 0}},
	}

	for _, testCase := range testCases {
		result, got := matchMul(testCase.Value)
		want := testCase.Want

		if got != want || result != testCase.IsValid {
			t.Error(testCase, got, want)
		}
	}
}

type TestCaseGetSum struct {
	Value string
	Want  int
}

func TestGetSum(t *testing.T) {
	var testCases = []TestCaseGetSum{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
	}

	for _, testCase := range testCases {
		got := getSum(testCase.Value)
		want := testCase.Want

		if got != want {
			t.Error(testCase, got, want)
		}
	}
}
