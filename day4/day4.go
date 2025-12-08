package day4

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // Up, Down, Left, Right
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonals
}

func getAdjacentRolls(input [][]byte, row, col int) int {
	total := 0
	maxRow, maxCol := len(input), len(input[0])
	for _, d := range directions {
		newRow, newCol := row+d[0], col+d[1]
		if newRow >= 0 && newRow < maxRow && newCol >= 0 && newCol < maxCol {
			if isRollOfPaper(input[newRow][newCol]) {
				total++
			}
		}
	}
	return total
}

func isRollOfPaper(char byte) bool {
	return char == '@'
}

func printingDepartment(filepath string) (int, error) {
	fileReader, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer fileReader.Close()
	var diagram [][]byte
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		line := scanner.Text()
		chars := []byte(line)
		diagram = append(diagram, chars)
	}
	var count int
	for row := 0; row < len(diagram); row++ {
		for col := 0; col < len(diagram[row]); col++ {
			if isRollOfPaper(diagram[row][col]) {
				adjacentRolls := getAdjacentRolls(diagram, row, col)
				if adjacentRolls < 4 {
					count++
				}
			}
		}
	}
	return count, nil
}

func Day4() {
	filepath := "day4/input.txt"
	solution, err := printingDepartment(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[day 4 - Printing Department] The solution is: %d\n", solution)
}
