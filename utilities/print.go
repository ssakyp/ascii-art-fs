package utilities

import (
	"fmt"
	"strings"
)

func PrintAscii(inSplit []string, slice []string, letterSlice []rune, color string, isColor bool) {
	// ANSI escape codes for color
	reset := "\033[0m"
	// Start iterating over a slice of strings separated by newline
	for _, c := range inSplit {
		// Create j and k counters for eliminating extra newlines
		j := 0
		k := 0
		// Start an inner loop for printing line by line
		for i := 0; i <= 7; i++ {
			// Start iterating over each character of a string slice element
			for _, e := range c {
				a := strings.Split(string(slice[e-32]), "\n")
				// Handle the case of first line of space character
				if i == 0 && e == ' ' {
					if isColor && ContainsLet(e, letterSlice) {
						fmt.Print(color + (a[1]) + reset)
					} else {
						fmt.Print(a[1])
					}

				} else {
					if isColor && ContainsLet(e, letterSlice) {
						fmt.Print(color + (a[i]) + reset)
					} else {
						fmt.Print(a[i])
					}
				}
			}
			// Handling the case of extra new lines
			if c != "" {
				fmt.Println()
			}
			k++

		}
		j++
		// Handling the case of extra new lines
		if c == "" || (j == len(inSplit)-1 && k == 7) {
			fmt.Println()
		}

	}
}
