package main

import (
	"ascii-art-color/utilities"
	"os"
)

func main() {
	args := os.Args
	end := utilities.Run(args)
	if end {
		return
	}
}
