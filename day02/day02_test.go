package day02

import (
	util "dkwgit/aoc2023/utility"
	"testing"
)

func TestPuzzle01(t *testing.T) {
	type args struct {
		dc *util.DayContext
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"sample input 02-01",
			args{util.NewDayContext("day02", []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			})}, "day02-01: 8",
		},
		{"result 02-01", args{util.NewDayContext("day02", nil)}, "day02-01: 2528"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle01(tt.args.dc); got != tt.want {
				t.Errorf("Puzzle01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle02(t *testing.T) {
	type args struct {
		dc *util.DayContext
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"sample input 02-02",
			args{util.NewDayContext("day02", []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			})}, "day02-02: 2286",
		},
		{"result 02-02", args{util.NewDayContext("day02", nil)}, "day02-02: 67363"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle02(tt.args.dc); got != tt.want {
				t.Errorf("Puzzle02() = %v, want %v", got, tt.want)
			}
		})
	}
}
