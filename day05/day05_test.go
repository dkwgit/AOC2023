package day05

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
			"result 05-01",
			util.NewDayContext("day05", nil, nil),
			"day05-01: 0",
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
			"result 05-02",
			util.NewDayContext("day05", nil, nil),
			"day05-02: 0",
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
