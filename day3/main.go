package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := readFile()

	fmt.Println(getSum(input))
}

func readFile() string {
	var sb strings.Builder

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		result := scanner.Text()

		sb.WriteString(result)

	}

	return sb.String()
}

func getSum(value string) int {
	sum := 0
	const MIN_MULT_LENGTH = 8
	enableMultiplying := true

	for i := 0; i < len(value)-MIN_MULT_LENGTH; i++ {

		isDo := matchDo(value[i:])

		if isDo {
			enableMultiplying = true
		}

		isDont := matchDont(value[i:])

		if isDont {
			enableMultiplying = false
		}

		if enableMultiplying {
			isMul, res := matchMul(value[i:])

			if isMul {
				// fmt.Println(isMul, res)
				sum += res.x * res.y
			}
		}

	}
	return sum
}

type Result struct {
	x int
	y int
}

func matchMul(value string) (bool, Result) {
	const PREFIX_LENGTH = 4
	const POSTFIX_LENGTH = 1

	start := string(value[:PREFIX_LENGTH])
	if start != "mul(" {
		return false, Result{0, 0}
	}

	firstNumber, stopIndex := parseNumber(value[PREFIX_LENGTH:])

	if string(value[PREFIX_LENGTH+stopIndex]) != "," || len(firstNumber) == 0 {
		return false, Result{0, 0}
	}

	stopIndex += PREFIX_LENGTH

	secondNumber, sIndex := parseNumber(value[stopIndex+POSTFIX_LENGTH:])

	stopIndex += sIndex + POSTFIX_LENGTH

	if string(value[stopIndex]) != ")" || len(secondNumber) == 0 {
		return false, Result{0, 0}
	}

	return true, Result{stringToInt(firstNumber), stringToInt(secondNumber)}
}

func matchDont(value string) bool {
	const LENGTH = 7
	start := string(value[:LENGTH])

	return start == "don't()"
}

func matchDo(value string) bool {
	const LENGTH = 4
	start := string(value[:LENGTH])

	return start == "do()"
}

func parseNumber(value string) (number string, index int) {
	number = ""
	stopIndex := 0

	for index, v := range value {
		if isNumber(string(v)) {
			number += string(v)
		} else {
			stopIndex = index

			break
		}
	}

	return number, stopIndex
}

func isNumber(value string) bool {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	return slices.Contains(numbers, value)
}

func stringToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return number
}
