package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	list := [][]int{}
	count := 0

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		result := splitNumbers(scanner.Text())

		var element []int

		for i := 0; i < len(result); i++ {
			element = append(element, stringToInt(result[i]))
		}

		list = append(list, element)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for j := 0; j < len(list); j++ {
		var element []int = list[j]

		var isSafe = isSaveReportV2(element)

		if isSafe {
			count++
		}
	}

	fmt.Println(count)
}

func isSaveReport(list []int) bool {
	var isSafe = true
	var isDesc = isDesc(list)

	for i := 0; i < len(list)-1; i++ {
		var distance = (list[i] - list[i+1])
		if isDesc {
			if !(distance <= 3 && distance > 0) {
				isSafe = false
			}
		} else {
			if !(distance >= -3 && distance < 0) {
				isSafe = false
			}
		}
	}

	return isSafe
}

func isSaveReportV2(list []int) bool {

	if isSaveReport(list) {
		return true
	} else {
		for i := 0; i < len(list); i++ {
			permutationList := removeIndex(list, i)
			if isSaveReport(permutationList) {
				return true
			}
		}
		return false
	}
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isDesc(list []int) bool {
	return list[0] > list[1]
}

func splitNumbers(line string) []string {
	const DELIMITER = " "

	return strings.Split(line, DELIMITER)
}

func stringToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return number
}
