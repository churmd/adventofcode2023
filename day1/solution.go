package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/churmd/adventofcode2023/common"
	"github.com/churmd/higherorder"
)

func Solve1() {
	lines := common.SplitNewLines(input)
	sum := sumFirstAndLastDigitsPerLine(lines)
	fmt.Printf("What is the sum of all of the calibration values?\n%d\n", sum)
}

func Solve2() {
	replacedNumWords := replaceWordNumsWithDigits(input)
	lines := common.SplitNewLines(replacedNumWords)
	sum := sumFirstAndLastDigitsPerLine(lines)
	fmt.Printf("What is the sum of all of the calibration values?\n%d\n", sum)
}

func sumFirstAndLastDigitsPerLine(lines []string) int {
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

		fmt.Printf("%s %c %c %d\n", line, firstDigit, lastDigit, value)

		sum += value
	}
	return sum
}

func findFirstDigit(runes []rune) rune {
	digits := higherorder.Filter(unicode.IsDigit, runes)
	return digits[0]
}

func replaceWordNumsWithDigits(s string) string {
	for {
		updated, done := replaceFirstWordNum(s)
		s = updated
		if done {
			break
		}
	}

	return s
}

func replaceFirstWordNum(s string) (string, bool) {
	wordsToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var wordToReplace string
	var currentIndex int
	wordFound := false

	for word, _ := range wordsToDigit {
		index := strings.Index(s, word)
		if index == -1 {
			continue
		}
		if !wordFound {
			wordToReplace = word
			currentIndex = index
			wordFound = true
		}
		if index < currentIndex {
			wordToReplace = word
			currentIndex = index
			wordFound = true
		}
	}

	if wordFound {
		updated := strings.Replace(s, wordToReplace, wordsToDigit[wordToReplace], 1)
		return updated, false
	}

	return s, true
}
