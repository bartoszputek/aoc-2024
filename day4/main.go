package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile()

	fmt.Println(input)
}

func readFile() []string {
	var matrix []string

	file, err := os.Open("./input.txt")

	var k string = "test"

	println(k[4])

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {

		result := scanner.Text()

		matrix = append(matrix, result)
	}

	return matrix
}

func matchHorizontal(x int, y int, matrix []string) int {

	if y+3 < len(matrix[x]) {
		if rune(matrix[x][y]) == 'X' {
			if rune(matrix[x][y+1]) == 'M' {
				if rune(matrix[x][y+2]) == 'A' {
					if rune(matrix[x][y+3]) == 'S' {
						return 1
					}
				}
			}
		}
	}
	return 0
}
