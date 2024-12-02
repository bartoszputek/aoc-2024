package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	var firstList []int
	var secondList []int

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := splitNumbers(scanner.Text())

		firstList = append(firstList, stringToInt(result[0]))
		secondList = append(secondList, stringToInt(result[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	fmt.Println(calculateDistance(firstList, secondList))
	fmt.Println(calculateSimilarityScore(firstList, secondList))

}

func splitNumbers(line string) []string {
	const DELIMITER = "   "

	return strings.Split(line, DELIMITER)
}

func stringToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return number
}

func calculateDistance(firstList []int, secondList []int) int {
	var distance int = 0

	for i := 0; i < len(firstList); i++ {
		distance += int(math.Abs(float64(firstList[i] - secondList[i])))
	}

	return distance
}

func calculateSimilarityScore(firstList []int, secondList []int) int {
	var similarityScore int = 0

	for i := 0; i < len(firstList); i++ {
		count := 0
		currentValue := firstList[i]

		for j := 0; j < len(secondList); j++ {
			candidate := secondList[j]
			if currentValue == candidate {
				count += 1
			}
		}

		similarityScore += currentValue * count
	}

	return similarityScore
}
