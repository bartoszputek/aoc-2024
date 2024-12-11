package main

import "testing"

func TestMatrixGetHorizontalRight(t *testing.T) {
	type TestCase struct {
		matrix Matrix
		x      int
		y      int
		Want   string
	}

	var testCases = []TestCase{
		{Matrix{matrix: []string{"AAAXMAS"}}, 3, 0, "XMAS"},
		{Matrix{matrix: []string{"", "AAAXMAS"}}, 4, 1, "MAS"},
	}

	for _, testCase := range testCases {
		got := testCase.matrix.GetHorizontalRight(testCase.x, testCase.y)
		want := testCase.Want

		if got != want {
			t.Error(testCase, got, want)
		}
	}
}

func TestMatrixGetHorizontalLeft(t *testing.T) {
	type TestCase struct {
		matrix Matrix
		x      int
		y      int
		Want   string
	}

	var testCases = []TestCase{
		{Matrix{matrix: []string{"SAMXXXX"}}, 4, 0, "XMAS"},
		{Matrix{matrix: []string{"", "SAM"}}, 3, 1, "MAS"},
		{Matrix{matrix: []string{"", ""}}, 0, 1, ""},
	}

	for _, testCase := range testCases {
		got := testCase.matrix.GetHorizontalLeft(testCase.x, testCase.y)
		want := testCase.Want

		if got != want {
			t.Error(testCase, got, want)
		}
	}
}

func TestMatrixGetVerticalUp(t *testing.T) {
	type TestCase struct {
		matrix Matrix
		x      int
		y      int
		Want   string
	}

	var testCases = []TestCase{
		{Matrix{matrix: []string{"SXX", "AXX", "MXX", "XXX"}}, 0, 4, "XMAS"},
		{Matrix{matrix: []string{"SXX", "AXX", "MXX", "XXX"}}, 0, 3, "MAS"},
		{Matrix{matrix: []string{"SXX"}}, 1, 0, ""},
	}

	for _, testCase := range testCases {
		got := testCase.matrix.GetVerticalUp(testCase.x, testCase.y)
		want := testCase.Want

		if got != want {
			t.Error(testCase, got, want)
		}
	}
}
