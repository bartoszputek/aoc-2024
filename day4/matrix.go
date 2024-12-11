package main

type Matrix struct {
	matrix []string
}

const WORD_TO_FIND = "XMAS"

func (m Matrix) GetHorizontalRight(x int, y int) string {

	row := m.matrix[y]

	if x+len(WORD_TO_FIND) > len(row) {
		return row[x:]
	}

	return row[x : x+len(WORD_TO_FIND)]
}

func (m Matrix) GetHorizontalLeft(x int, y int) string {

	row := m.matrix[y]

	if x < len(WORD_TO_FIND) {
		return Reverse(row[:x])
	}

	return Reverse(row[x-len(WORD_TO_FIND) : x])
}

func (m Matrix) GetVerticalUp(x int, y int) string {

	length := Min(y, len(WORD_TO_FIND))

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(m.matrix[y-i-1][x])
	}
	return string(runes)
}

func (m Matrix) GetVerticalDown(x int, y int) string {

	length := Min(y, len(WORD_TO_FIND))

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(m.matrix[y-i-1][x])
	}
	return string(runes)
}

func (m Matrix) ColumnLength() int {
	return len(m.matrix)
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
