package main

import (
	"testing"
)

func TestAscReport(t *testing.T) {
	list := []int{46, 47, 49, 52, 54, 56, 58, 59}

	got := isSaveReport(list)
	want := true

	if got != want {
		t.Error(got, want)
	}
}

func TestDescRaport(t *testing.T) {
	list := []int{87, 86, 84, 83, 81, 80}

	got := isSaveReport(list)
	want := true

	if got != want {
		t.Error(got, want)
	}
}

func TestIncorrectAscRaport(t *testing.T) {
	list := []int{46, 47, 49, 57, 54, 56, 58, 59}

	got := isSaveReport(list)
	want := false

	if got != want {
		t.Error(got, want)
	}
}

func TestIncorrectDescRaport(t *testing.T) {
	list := []int{87, 86, 84, 83, 71, 80}

	got := isSaveReport(list)
	want := false

	if got != want {
		t.Error(got, want)
	}
}

func TestAscReportV2(t *testing.T) {
	list := [][]int{
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{7, 6, 4, 2, 1},
	}

	for i := 0; i < len(list); i++ {
		got := isSaveReportV2(list[i])
		want := true

		if got != want {
			t.Error(got, want)
		}
	}
}

func TestDescReportV2(t *testing.T) {
	list := [][]int{
		{48, 46, 47, 49, 51, 54, 56},
		{1, 1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 5},
		{5, 1, 2, 3, 4, 5},
		{1, 4, 3, 2, 1},
		{1, 6, 7, 8, 9},
		{1, 2, 3, 4, 3},
		{9, 8, 7, 6, 7},
		{7, 10, 8, 10, 11},
		{29, 28, 27, 25, 26, 25, 22, 20},
	}

	for i := 0; i < len(list); i++ {
		got := isSaveReportV2(list[i])
		want := true

		if got != want {
			t.Error(list[i], got, want)
		}
	}
}
