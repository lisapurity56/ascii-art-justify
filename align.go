package main

import "strings"

// Align lines
func alignText(lines []string, align string, width int) []string {
	var result []string

	for _, line := range lines {
		lineLength := len(line)

		switch align {

		case "left":
			result = append(result, line)

		case "right":
			spaces := width - lineLength
			if spaces < 0 {
				spaces = 0
			}
			result = append(result, strings.Repeat(" ", spaces)+line)

		case "center":
			spaces := (width - lineLength) / 2
			if spaces < 0 {
				spaces = 0
			}
			result = append(result, strings.Repeat(" ", spaces)+line)

		default:
			result = append(result, line)
		}
	}

	return result
}

func justifyText(input string, banner []string, width int) []string {
	words := strings.Fields(input)

	wordAscii := [][]string{}
	for _, w := range words {
		wordAscii = append(wordAscii, buildWordAscii(w, banner))
	}

	result := make([]string, 8)

	totalWordWidth := 0
	for _, w := range wordAscii {
		totalWordWidth += len(w[0])
	}

	gaps := len(words) - 1
	if gaps <= 0 {
		return wordAscii[0]
	}

	totalSpaces := width - totalWordWidth
	spacePerGap := totalSpaces / gaps

	for row := 0; row < 8; row++ {
		line := ""
		extra := totalSpaces % gaps

		for i := 0; i < len(wordAscii); i++ {
			line += wordAscii[i][row]

			if i < gaps {
				spaces := spacePerGap
				if extra > 0 {
					spaces++
					extra--
				}
				line += strings.Repeat(" ", spaces)
			}
		}

		result[row] = line
	}

	return result
}

func buildWordAscii(word string, banner []string) []string {
	result := make([]string, 8)

	for i := 0; i < 8; i++ {
		line := ""
		for _, ch := range word {
			index := (int(ch)-32)*9 + i + 1
			line += banner[index]
		}
		result[i] = line
	}

	return result
}