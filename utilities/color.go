package utilities

import "fmt"

func ColorMatch(color string, isOnly bool) (string, bool) {
	// Extract color and string && and assign it to ANSI escape codes
	end := false
	switch color {
	case "black":
		color = "\033[30m"
	case "red":
		color = "\033[31m"
	case "green":
		color = "\033[32m"
	case "yellow":
		color = "\033[33m"
	case "blue":
		color = "\033[34m"
	case "magenta":
		color = "\033[35m"
	case "cyan":
		color = "\033[36m"
	case "white":
		color = "\033[37m"
	case "orange":
		color = "\033[91m"
	default:
		if !isOnly {
			fmt.Println("INVALID COLOR\nCHOOSE FROM THE LIST OF COLORS:\nblack\nred\ngreen\nyellow\nblue\nmagenta\ncyan\nwhite\norange")
			end = true
		}
	}
	return color, end
}
