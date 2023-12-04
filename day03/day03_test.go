package day03

import (
	util "dkwgit/aoc2023/utility"
	"testing"
)

func TestPuzzle01(t *testing.T) {
	tests := []struct {
		name string
		dc   *util.DayContext
		want string
	}{
		{"sample input 03-01", nil, "day03-01: 0"},
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
		{"sample input 03-02", nil, "day03-02: 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle02(tt.dc); got != tt.want {
				t.Errorf("Puzzle02() = %v, want %v", got, tt.want)
			}
		})
	}
}
