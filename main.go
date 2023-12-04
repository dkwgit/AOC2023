package main

import (
	"dkwgit/aoc2023/day01"
	"dkwgit/aoc2023/day02"
	"dkwgit/aoc2023/day03"
	"fmt"

	util "dkwgit/aoc2023/utility"
)

func main() {
	fmt.Println("Welcome to AOC2023!")

	dcDay01 := util.NewDayContext("day01", nil)
	fmt.Println(day01.Puzzle01(dcDay01))
	fmt.Println(day01.Puzzle02(dcDay01))

	dcDay02 := util.NewDayContext("day02", nil)
	fmt.Println(day02.Puzzle01(dcDay02))
	fmt.Println(day02.Puzzle02(dcDay02))

	dcDay03 := util.NewDayContext("day03", nil)
	fmt.Println(day03.Puzzle01(dcDay03))
	fmt.Println(day03.Puzzle02(dcDay03))
}
