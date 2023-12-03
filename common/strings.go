package common

import "strings"

func SplitNewLines(s string) []string {
	return strings.Split(s, "\n")
}