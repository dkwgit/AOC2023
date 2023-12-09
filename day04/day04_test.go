package day04

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
			"sample input 04-01",
			util.NewDayContext("day04", []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			}, []string{"5", "8"}),
			"day04-01: 13",
		},
		{
			"result 04-01",
			util.NewDayContext("day04", nil, []string{"10", "25"}),
			"day04-01: 23028",
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
			"sample input 04-02",
			util.NewDayContext("day04", []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			}, []string{"5", "8"}),
			"day04-02: 30",
		},
		{
			"result 04-02",
			util.NewDayContext("day04", nil, []string{"10", "25"}),
			"day04-02: 9236992",
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
