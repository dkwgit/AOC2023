package day01

import (
	"testing"
)

func TestPuzzle01(t *testing.T) {
	type args struct {
		inputLines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample input 01-01", args{[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}}, "142"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle01(tt.args.inputLines); got != tt.want {
				t.Errorf("Puzzle01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle02(t *testing.T) {
	type args struct {
		inputLines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sample input 01-02", args{[]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}}, "281"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Puzzle02(tt.args.inputLines); got != tt.want {
				t.Errorf("Puzzle02() = %v, want %v", got, tt.want)
			}
		})
	}
}
