package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filepath string) []string {
	lines := make([]string, 0)

	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)

	}

	return lines
}
