package main

import (
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func getAlign(option string) string {
	if strings.HasPrefix(option, "--align=") {
		return strings.TrimPrefix(option, "--align=")
	}
	return "left"
}
