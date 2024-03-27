package utilities

import (
	"fmt"
	"os"
	"strings"
)

// Handles cases of visible and unvisible newlines
func SplitByNewLine(str string) []string {
	word := ""
	var final []string
	slice := strings.Split(str, "\\n")
	for _, el := range slice {
		for _, char := range el {
			if char == 10 {
				final = append(final, word)
				word = ""
				continue
			} else {
				word += string(char)
			}
		}
		if word == "" {
			final = append(final, word)
		} else {
			final = append(final, word)
			word = ""
		}
	}
	return final
}

// Function for finding the letters in the input text when color flag is used
func ContainsLet(char rune, slice []rune) bool {
	counter := 0
	for _, c := range slice {
		if c == char {
			counter++
		}
	}
	return counter >= 1
}

// Checks if the string contains only Latin Alphabet
func IsAlpha(in string) bool {
	is := true
	for _, c := range in {
		if (c < 32 || c > 126) && c != 10 {
			is = false
		}
	}
	return is
}

// Function handles when the input is empty or only newlines
func OnlyNewLinePrinter(in string, letterSlice []rune) bool {
	count := strings.Count(in, "\\n")
	if in == "\\n" {
		fmt.Println()
		return true
	}

	if in == "" {
		return true
	}

	if len(in) == count*2 {
		for count > 0 {
			fmt.Println()
			count--
		}
		return true
	}
	return false
}

// Validator function for banners
func isBannerStr(banner string) bool {
	return banner == "thinkertoy" || banner == "shadow" || banner == "standard"
}

// Function for checking
func ArgumentsChecker(args []string, isColor bool, isOnly bool) (bool, bool, bool, bool, string, string) {
	color := ""
	end := false
	length := len(args)
	banner := "standard"
	isBanner := false

	for _, el := range args {
		if !IsAlpha(el) {
			fmt.Println("Invalid input, only Latin Alphabet is allowed.")
			end = true
			break
		}
	}
	// Define the color flag
	// Case when go run . --color=red ab abc || go run . --color=red ab
	if (length == 3 || length == 4 && (!isBannerStr(args[3]) && OnlyAlpha(args))) && strings.HasPrefix(args[1], "--color=") {
		isColor = true
		for i := 8; i < len(args[1]); i++ {
			color += string(args[1][i])
		}

		// Case when go run . --color=red ab abc shadow
	} else if (length == 5) && (strings.HasPrefix(args[1], "--color=") && isBannerStr(args[4])) && OnlyAlpha(args) {
		isColor = true
		isBanner = true
		for i := 8; i < len(args[1]); i++ {
			color += string(args[1][i])
		}
		banner = args[4]

		// Case when go run . --color=red ab shadow
	} else if length == 4 && strings.HasPrefix(args[1], "--color=") && isBannerStr(args[3]) && OnlyAlpha(args) {
		isColor = true
		isBanner = true
		for i := 8; i < len(args[1]); i++ {
			color += string(args[1][i])
		}
		banner = args[3]

		// Case when go run . abc shadow
	} else if length == 3 && isBannerStr(args[2]) && OnlyAlpha(args) {
		isBanner = true
		banner = args[2]
	}

	// Make sure the only one argument is to be entered
	if length == 2 {
		isOnly = true
	} else if length == 3 && !isBanner && !isColor && OnlyAlpha(args) {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		end = true
	} else if length == 4 && !isColor && OnlyAlpha(args) {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> 'something'")
		end = true
	} else if length == 1 && OnlyAlpha(args) {
		fmt.Println("Usage: go run . [STRING] \n\nEX: go run . 'something'")
		end = true
	} else if ( (length == 5 && OnlyAlpha(args)) && ( !strings.HasPrefix(args[1], "--color=") || (!isBannerStr(args[4])) ) ) || (length > 5 && OnlyAlpha(args)){
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --color=<color> <letters to be colored> 'something' standard")
		end = true
	}

	return isColor, isOnly, isBanner, end, color, banner
}

// Finds the input text that needs to be printed for FS
func TextFinder(args []string, isColor bool, isOnly bool, isBanner bool) (string, []rune) {
	letters := ""
	in := ""
	if (len(args) == 4 && isColor && !isBanner) || (len(args) == 5 && isColor && isBanner) {
		letters = args[2]
		in = args[3]
	} else if (len(args) == 3 && isColor) || (len(args) == 4 && isColor && isBanner) {
		letters = args[2]
		in = args[2]
	} else if isOnly {
		in = args[1]
	} else if len(args) == 3 && isBanner && isColor {
		in = args[2]
	} else if len(args) == 3 && isBanner {
		in = args[1]
	}

	// Convert string to a slice of runes
	letterSlice := []rune(letters)
	return in, letterSlice
}

// Picks the correct banner from txt
func BannerPicker(banner string) ([]string, bool) {
	finalBanner := banner + ".txt"
	end := false
	// Extract text from standard.txt in a byteSlice type
	byteSlice, err := os.ReadFile(finalBanner)
	if err != nil {
		fmt.Println("Wrong banner path")
		end = true
	}
	// Convert the the byteslice into string type
	text := string(byteSlice)
	text = strings.ReplaceAll(text, "\r", "")
	// Split the string into each ascii element
	slice := strings.Split(text, "\n\n")
	return slice, end
}

func OnlyAlpha(args []string) bool {
	counter := 0
	for _, el := range args {
		if IsAlpha(el) {
			counter++
		}
	}
	return counter == len(args)
}