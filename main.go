package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// get terminal width using stty
func getTerminalWidth() (int, error) {
	// run command safely using terminal
	cmd := exec.Command("sh", "-c", "stty size < /dev/tty")

	// get output (example: "24 80")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// clean and split output
	parts := strings.Fields(string(output))
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid stty output")
	}

	// second value is width
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return width, nil
}

func main() {
	args := os.Args

	// check arguments
	if len(args) != 4 {
		fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("Example: go run . --align=right  something  standard")
		return
	}

	// get alignment type
	alignType := getAlign(args[1])

	// get text and banner file
	text := args[2]
	bannerFile := args[3]

	// read banner file
	banner, err := readBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// align ASCII art (width = 80)
	width, err := getTerminalWidth()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var output []string

	// special case for justify
	if alignType == "justify" {
		output = justifyText(text, banner, width)
	} else {
		ascii := generateAscii(text, banner)
		output = alignText(ascii, alignType, width)
	}

	// print result
	for _, line := range output {
		fmt.Println(line)
	}
}
