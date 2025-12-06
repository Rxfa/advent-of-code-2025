package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeatedTwice(number int) bool {
	s := strconv.Itoa(number)
	n := len(s)
	if n%2 != 0 { // odd number of digits -> can never have sequence of digits repeated twice
		return false
	}
	return s[0:n/2] == s[n/2:n]
}

func getInvalidIds(start int, end int) []int {
	var invalidIds []int
	for i := start; i < end; i++ {
		if isRepeatedTwice(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func giftShop(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	ranges := strings.Split(line, ",")
	var invalidIds []int
	for _, numRange := range ranges {
		splitRange := strings.Split(numRange, "-")
		start, err := strconv.Atoi(splitRange[0])
		if err != nil {
			return 0, nil
		}
		end, err := strconv.Atoi(splitRange[1])
		if err != nil {
			return 0, nil
		}
		invalidIds = append(invalidIds, getInvalidIds(start, end)...)
	}
	idSum := 0
	for _, i := range invalidIds {
		idSum += i
	}
	return idSum, nil
}

func Day2() {
	filepath := "day2/input.txt"
	solution, err := giftShop(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[day 2] The solution is: %d\n", solution)
}
