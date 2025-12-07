package day3

import (
	"bufio"
	"fmt"
	"os"
)

func maxDigit(s string) (byte, int) {
	maxNum := byte('0')
	maxNumIdx := -1
	for i := 0; i < len(s); i++ {
		if s[i] > maxNum {
			maxNum = s[i]
			maxNumIdx = i
		}
	}
	return maxNum, maxNumIdx
}

func getMaxJoltage(bank string) int {
	a, aIdx := maxDigit(bank)
	var b byte

	if aIdx == len(bank)-1 {
		b, _ = maxDigit(bank[:aIdx])
	} else {
		b, _ = maxDigit(bank[aIdx+1:])
	}
	return int(a-'0')*10 + int(b-'0')
}

func lobby(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += getMaxJoltage(scanner.Text())
	}
	return total, nil
}

func Day3() {
	filepath := "day3/input.txt"
	solution, err := lobby(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[day 3 - lobby] The solution is: %d\n", solution)
}
