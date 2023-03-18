package utils

import (
	"fmt"
	"strings"
)

func Logger(msg string) {
	title := createTitle("public", len(msg))
	logTemplate(title, msg)
}

func privateLogger(msg string) {
	title := createTitle("private", len(msg))
	logTemplate(title, msg)
}

func logTemplate(title string, msg string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println(msg)
	fmt.Println(strings.Repeat("-", len(msg)))
	fmt.Println()
}

func createTitle(title string, msgLen int) string {
	baseStr := strings.Repeat("-", msgLen)
	return strings.Map(transform(msgLen, title), baseStr)
}

func transform(msgLen int, title string) func(r rune) rune {
	start := (msgLen - len(title)) / 2
	end := msgLen - 5
	titleIndex := 0
	count := 0

	return func(r rune) rune {
		count += 1

		if titleIndex >= len(title) {
			return r
		}

		if count > start && count < end {
			char := title[titleIndex]
			titleIndex += 1
			return rune(char)
		}

		return r
	}
}
