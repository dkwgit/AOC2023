package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	util "dkwgit/aoc2023/utility"

	strrv "4d63.com/strrev"
)

func Puzzle01(dc *util.DayContext) string {
	var values = make([]int, len(dc.InputLines))

	for _, line := range dc.InputLines {
		var digits string = ""
		digits = firstDigitInStringForward(line)
		digits += firstDigitInStringBackward(line)

		value, err := strconv.Atoi(digits)
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", digits, err))
		}
		values = append(values, value)
	}

	result := sum(values)

	return fmt.Sprintf("%s-01: %d", dc.Name, result)
}

func Puzzle02(dc *util.DayContext) string {
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9", "eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")

	var values = make([]int, len(dc.InputLines))

	for _, line := range dc.InputLines {
		matches := re.FindAllString(line, -1)
		reverseMatches := re.FindAllString(strrv.Reverse(line), -1)

		var digits string = ""
		digits = replacer.Replace(matches[0])
		digits += replacer.Replace(reverseMatches[0])

		value, err := strconv.Atoi(digits)
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", digits, err))
		}
		values = append(values, value)
	}

	result := sum(values)

	return fmt.Sprintf("%s-02: %d", dc.Name, result)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func firstDigitInStringForward(s string) string {
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			return string(s[i])
		}
	}
	panic("No digit found in string in firstDigitInStringForward, for line " + s)
	return ""
}

func firstDigitInStringBackward(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return string(s[i])
		}
	}
	panic("No digit found in string in firstDigitInStringForward, for line " + s)
	return ""
}
