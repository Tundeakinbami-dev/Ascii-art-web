package main

import (
	"strings"
)
func render(text []string, Ascii map[rune][]string) string {
	var output strings.Builder
	for _, line := range text {
		if line == "" {
			output.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if ch, ok := Ascii[ch]; ok{
					output.WriteString(ch[row])
				}
			}
			output.WriteString("\n")
		}
	}
	return output.String()
}