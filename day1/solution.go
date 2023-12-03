package day1

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/churmd/adventofcode2023/common"
	"github.com/churmd/higherorder"
)

func Solve() {
	lines := common.SplitNewLines(input)
	sum := 0
	for _, line := range lines {
		runes := []rune(line)
		firstDigit := findFirstDigit(runes)
		lastDigit := findFirstDigit(higherorder.Reverse(runes))
		number := fmt.Sprintf("%c%c", firstDigit, lastDigit)
		value, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		sum += value
	}	

	fmt.Printf("What is the sum of all of the calibration values?\n%d\n", sum)
}

func findFirstDigit(runes []rune) rune {
	digits := higherorder.Filter(unicode.IsDigit, runes)
	return digits[0]
}