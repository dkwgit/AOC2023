package day03

import (
	"testing"

	util "github.com/dkwgit/aoc2023/utility"
)

func TestPuzzle01(t *testing.T) {
	tests := []struct {
		name string
		dc   *util.DayContext
		want string
	}{
		{
			"sample input 03-01",
			util.NewDayContext("day03", []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			}, nil),
			"day03-01: 4361",
		},
		{
			"result 03-01",
			util.NewDayContext("day03", nil, nil),
			"day03-01: 525181",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle01(tt.dc); got != tt.want {
				t.Errorf("Puzzle01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle02(t *testing.T) {
	tests := []struct {
		name string
		dc   *util.DayContext
		want string
	}{
		{
			"sample input 03-02",
			util.NewDayContext("day03", []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			}, nil),
			"day03-02: 467835",
		},
		{
			"result 03-02",
			util.NewDayContext("day03", nil, nil),
			"day03-02: 84289137",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle02(tt.dc); got != tt.want {
				t.Errorf("Puzzle02() = %v, want %v", got, tt.want)
			}
		})
	}
}
