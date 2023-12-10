package main

import (
	"fmt"

	"github.com/dkwgit/aoc2023/day01"
	"github.com/dkwgit/aoc2023/day02"
	"github.com/dkwgit/aoc2023/day03"
	"github.com/dkwgit/aoc2023/day04"
	"github.com/dkwgit/aoc2023/day05"

	util "github.com/dkwgit/aoc2023/utility"
)

func main() {
	fmt.Println("Welcome to AOC2023!")

	dcDay01 := util.NewDayContext("day01", nil, nil)
	fmt.Println(day01.Puzzle01(dcDay01))
	fmt.Println(day01.Puzzle02(dcDay01))

	dcDay02 := util.NewDayContext("day02", nil, nil)
	fmt.Println(day02.Puzzle01(dcDay02))
	fmt.Println(day02.Puzzle02(dcDay02))

	dcDay03 := util.NewDayContext("day03", nil, nil)
	fmt.Println(day03.Puzzle01(dcDay03))
	fmt.Println(day03.Puzzle02(dcDay03))

	dcDay04 := util.NewDayContext("day04", nil, []string{"10", "25"})
	fmt.Println(day04.Puzzle01(dcDay04))
	fmt.Println(day04.Puzzle02(dcDay04))

	dcDay05 := util.NewDayContext("day05", nil, nil)
	fmt.Println(day05.Puzzle01(dcDay05))
	fmt.Println(day05.Puzzle02(dcDay05))
}
