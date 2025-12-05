package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNextRotation(str string) (int, error) {
	if len(str) < 2 {
		return 0, fmt.Errorf("invalid rotation string: %q", str)
	}
	rotationSide := str[0]
	i, err := strconv.Atoi(str[1:])
	if err != nil {
		return 0, err
	}
	switch rotationSide {
	case 'R':
		return i, nil
	case 'L':
		return -i, nil
	default:
		return 0, fmt.Errorf("rotation side must be 'L' or 'R', got %q", rotationSide)
	}
}

func secretEntrance(documentPath string) (int, error) {
	file, err := os.Open(documentPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	dial := 50
	password := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		line := scanner.Text()
		nextRotation, err := getNextRotation(line)
		if err != nil {
			return 0, err
		}
		dial += nextRotation
		dial = (dial%100 + 100) % 100
		if dial == 0 {
			password++
		}
	}
	return password, nil
}

func main() {
	filePath := "input.txt"
	password, err := secretEntrance(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("The password is: %d\n", password)
}
