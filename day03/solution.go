package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isDigit(value rune) bool {
	return value >= 48 && value <= 57
}

func isDot(value rune) bool {
	return value == 46
}

func part1() {
	f, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var schema []string
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		formatted := "." + line + "."
		schema = append(schema, formatted)
	}

	for row := 0; row < len(schema); row++ {
		// row
		for column := 0; column < len(schema[row]); column++ {
			// column
			if isDigit(rune(schema[row][column])) {
				startIdx := column
				// go to next column
				column++
				for isDigit(rune(schema[row][column])) && column < len(schema[row])-1 {
					column++
				}

				endIdx := column
				isValid := checkIfIsValid(startIdx, endIdx, row, schema)
				if isValid {

					converted, _ := strconv.Atoi(string(schema[row][startIdx:endIdx]))

					total += converted
				}
			}
		}
	}

	fmt.Printf("Solution for Day 3 Part 1: %d\n", total)
}

func checkIfIsValid(startIdx int, endIdx int, row int, schema []string) bool {
	for i := startIdx; i < endIdx; i++ {
		// check north
		if row > 0 {
			if !isDigit(rune(schema[row-1][i])) && !isDot(rune(schema[row-1][i])) {
				return true
			}
		}
		// check north east
		if row > 0 && i < len(schema[row])-1 {
			if !isDigit(rune(schema[row-1][i+1])) && !isDot(rune(schema[row-1][i+1])) {
				return true
			}
		}
		// check north west
		if row > 0 && i > 0 {
			if !isDigit(rune(schema[row-1][i-1])) && !isDot(rune(schema[row-1][i-1])) {
				return true
			}
		}

		// check south
		if row < len(schema)-1 {
			if !isDigit(rune(schema[row+1][i])) && !isDot(rune(schema[row+1][i])) {
				return true
			}
		}
		// check south east
		if row < len(schema)-1 && i < len(schema[row])-1 {
			if !isDigit(rune(schema[row+1][i+1])) && !isDot(rune(schema[row+1][i+1])) {
				return true
			}
		}
		// check south west
		if row < len(schema)-1 && i > 0 {
			if !isDigit(rune(schema[row+1][i-1])) && !isDot(rune(schema[row+1][i-1])) {
				return true
			}
		}

		// check east
		if i < len(schema[row])-1 {
			if !isDigit(rune(schema[row][i+1])) && !isDot(rune(schema[row][i+1])) {
				return true
			}
		}
		// check west
		if i > 0 {
			if !isDigit(rune(schema[row][i-1])) && !isDot(rune(schema[row][i-1])) {
				return true
			}
		}

	}

	return false

}

type part struct {
	x       int
	yStart  int
	yEnd    int
	partnum int
}

func part2() {
	f, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var schema []string
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		formatted := "." + line + "."
		schema = append(schema, formatted)
	}

	parts := []part{}

	for row := 0; row < len(schema); row++ {
		// row
		for column := 0; column < len(schema[row]); column++ {
			// column
			if isDigit(rune(schema[row][column])) {
				startIdx := column
				// go to next column
				column++
				for isDigit(rune(schema[row][column])) && column < len(schema[row])-1 {
					column++
				}

				endIdx := column
				isValid := checkIfIsValid(startIdx, endIdx, row, schema)
				if isValid {
					converted, _ := strconv.Atoi(string(schema[row][startIdx:endIdx]))
					parts = append(parts, part{x: row, yStart: startIdx, yEnd: endIdx, partnum: converted})
				}
			}
		}
	}

	for row := 0; row < len(schema); row++ {
		// row
		for column := 0; column < len(schema[row]); column++ {
			// column
			if schema[row][column] == '*' {
				// search near parts
				tmp_parts := []part{}
				for _, singlePart := range parts {
					if row > singlePart.x+1 || row < singlePart.x-1 {
						continue
					}
					if column > singlePart.yStart+(singlePart.yEnd-singlePart.yStart) || column < singlePart.yStart-1 {
						continue
					}
					tmp_parts = append(tmp_parts, singlePart)
				}

				if len(tmp_parts) == 2 {
					total += tmp_parts[0].partnum * tmp_parts[1].partnum
				}
			}
		}
	}

	fmt.Printf("Solution for Day 3 Part 2: %d\n", total)
}

func Solution() {
	part1()
	part2()
}
