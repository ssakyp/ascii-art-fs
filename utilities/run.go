package utilities

func Run(args []string) bool {
	end := false
	isColor := false
	isOnly := false
	color := ""
	var slice []string
	isBanner := false
	banner := ""
	isColor, isOnly, isBanner, end, color, banner = ArgumentsChecker(args, isColor, isOnly)

	if end {
		return end
	}

	// Bool variable for stopping the main function and displaying the color option
	if isColor {
		color, end = ColorMatch(color, isOnly)
		if end {
			return end
		}
	}

	// Variables for letters to be colored and the IN text to be colored
	in, letterSlice := TextFinder(args, isColor, isOnly, isBanner)

	// Handles the single and many new lines cases in the argument || Handles the single empty string case in the argument || Handles cases when not Latin Alphabet letters were intered
	if OnlyNewLinePrinter(in, letterSlice) {
		return true
	}

	slice, end = BannerPicker(banner)
	if end {
		return true
	}

	// Split the input argument variable in with new lines if there are any
	inSplit := SplitByNewLine(in)

	// Prints ASCII
	PrintAscii(inSplit, slice, letterSlice, color, isColor)
	return end
}
