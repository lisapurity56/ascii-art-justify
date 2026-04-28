package main

import (
	"os"
	"strings"
)

func readBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename + ".txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func generateAscii(text string, banner []string) []string {
	const charHeight = 8
	const blockSize = 9

	result := make([]string, charHeight)

	for i := 0; i < charHeight; i++ {
		result[i] = ""
	}

	for _, char := range text {

		// skip invalid characters
		if char < 32 || char > 126 {
			continue
		}

		index := int(char) - 32

		for row := 0; row < charHeight; row++ {
			pos := index*blockSize + 1 + row

			if pos < len(banner) {
				result[row] += banner[pos]
			}
		}
	}

	return result
}