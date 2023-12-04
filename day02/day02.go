package day02

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/dkwgit/aoc2023/utility"
)

type draw struct {
	blue  int
	green int
	red   int
}

type gameRecord struct {
	gameNumber int
	draws      []draw
	possible   bool
}

var gameRecordsMap map[*util.DayContext][]gameRecord = make(map[*util.DayContext][]gameRecord)

func Puzzle01(dc *util.DayContext) string {
	gameRecords, ok := gameRecordsMap[dc]
	if !ok {
		gameRecords = loadGameRecords(dc.InputLines)
		gameRecordsMap[dc] = gameRecords
	}

	var blueLimit = 14
	var greenLimit = 13
	var redLimit = 12

	for gameIndex, game := range gameRecords {
		for _, draw := range game.draws {
			if draw.blue > blueLimit || draw.green > greenLimit || draw.red > redLimit {
				gameRecords[gameIndex].possible = false
				break
			}
		}
	}

	var result = 0
	for _, game := range gameRecords {
		if game.possible {
			result += game.gameNumber
		}
	}

	return fmt.Sprintf("day02-01: %d", result)
}

func Puzzle02(dc *util.DayContext) string {
	gameRecords, ok := gameRecordsMap[dc]
	if !ok {
		gameRecords = loadGameRecords(dc.InputLines)
		gameRecordsMap[dc] = gameRecords
	}

	var result = 0

	for _, game := range gameRecords {
		maxDraw := draw{0, 0, 0}
		for _, draw := range game.draws {
			if draw.blue > maxDraw.blue {
				maxDraw.blue = draw.blue
			}
			if draw.green > maxDraw.green {
				maxDraw.green = draw.green
			}
			if draw.red > maxDraw.red {
				maxDraw.red = draw.red
			}
		}
		result += maxDraw.blue * maxDraw.green * maxDraw.red
	}

	return fmt.Sprintf("day02-02: %d", result)
}

func loadGameRecords(inputLines []string) []gameRecord {
	var gameRecords = make([]gameRecord, len(inputLines))

	for lineIndex, line := range inputLines {
		halves := strings.Split(line, ": ")
		gameNumberString := strings.Split(halves[0], " ")[1]
		gameNumber, err := strconv.Atoi(gameNumberString)
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", gameNumberString, err))
		}
		drawStrings := strings.Split(halves[1], "; ")
		var game = gameRecord{gameNumber, make([]draw, len(drawStrings)), true}

		for drawIndex, ds := range drawStrings {
			colorStrings := strings.Split(ds, ", ")
			var d draw
			for _, cs := range colorStrings {
				pair := strings.Split(cs, " ")
				count, err := strconv.Atoi(pair[0])
				if err != nil {
					panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", pair[0], err))
				}
				if pair[1] == "blue" {
					d.blue = count
				} else if pair[1] == "green" {
					d.green = count
				} else if pair[1] == "red" {
					d.red = count
				}
				game.draws[drawIndex] = d
			}
		}
		gameRecords[lineIndex] = game
	}
	return gameRecords
}
