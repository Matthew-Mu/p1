package main

import (
	"strings"
)

func findDigitString(s string) int {
	var digit int
	switch {
	case strings.HasPrefix(s, "one"):
		digit = 1
	case strings.HasPrefix(s, "two"):
		digit = 2
	case strings.HasPrefix(s, "three"):
		digit = 3
	case strings.HasPrefix(s, "four"):
		digit = 4
	case strings.HasPrefix(s, "five"):
		digit = 5
	case strings.HasPrefix(s, "six"):
		digit = 6
	case strings.HasPrefix(s, "seven"):
		digit = 7
	case strings.HasPrefix(s, "eight"):
		digit = 8
	case strings.HasPrefix(s, "nine"):
		digit = 9
	}
	return digit

}
