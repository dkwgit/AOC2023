package day01

import (
	"testing"

	util "github.com/dkwgit/aoc2023/utility"
)

func TestPuzzle01(t *testing.T) {
	tests := []struct {
		name    string
		context *util.DayContext
		want    string
	}{
		{"sample input 01-01", util.NewDayContext("day01", []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}), "day01-01: 142"},
		{"result 01-01", util.NewDayContext("day01", nil), "day01-01: 54667"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle01(tt.context); got != tt.want {
				t.Errorf("Puzzle01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle02(t *testing.T) {
	tests := []struct {
		name    string
		context *util.DayContext
		want    string
	}{
		{"sample input 01-02", util.NewDayContext("day01", []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}), "day01-02: 281"},
		{"result 01-02", util.NewDayContext("day01", nil), "day01-02: 54203"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle02(tt.context); got != tt.want {
				t.Errorf("Puzzle02() = %v, want %v", got, tt.want)
			}
		})
	}
}
