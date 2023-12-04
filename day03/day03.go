package day03

import (
	"fmt"
	"strconv"
	"unicode"

	util "github.com/dkwgit/aoc2023/utility"
)

type point struct {
	x int
	y int
}

type item struct {
	value          rune
	isDigit        bool
	symbolAdjacent bool
}

type schematic struct {
	rows       int
	columns    int
	data       [][]item
	partStarts []point
	partEnds   []point
}

func Puzzle01(dc *util.DayContext) string {
	var schematic = createSchematic(dc.InputLines)
	var parts = make([]int, 0)

	for i := 0; i < len(schematic.partStarts); i++ {
		start := schematic.partStarts[i]
		end := schematic.partEnds[i]

		var partString = ""
		var adjacency bool = false
		for x := start.x; x < end.x; x++ {
			if schematic.data[start.y][x].isDigit {
				partString += string(schematic.data[start.y][x].value)
				if schematic.data[start.y][x].symbolAdjacent {
					adjacency = true
				}
			} else {
				break
			}
		}

		if adjacency {
			part, _ := strconv.Atoi(partString)
			parts = append(parts, part)
		}
	}

	var result = util.SumIntArray(parts)

	return fmt.Sprintf("day03-01: %d", result)
}

func Puzzle02(dc *util.DayContext) string {
	var result = 0

	return fmt.Sprintf("day03-02: %d", result)
}

func createSchematic(inputLines []string) schematic {
	var partStarts []point = make([]point, 0)
	var partEnds []point = make([]point, 0)
	// Pad the input with empty rows and columns
	var rows = len(inputLines) + 2
	var columns = len(inputLines[0]) + 2

	var emptyRow = make([]item, columns)
	for i := 0; i < columns; i++ {
		emptyRow[i] = item{'.', false, false}
	}

	var data = make([][]item, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]item, columns)
		copy(data[i], emptyRow)
	}

	for row, line := range inputLines {
		row += 1
		var inPart bool = false
		var partEnd bool = false
		for column, char := range line {
			partEnd = false
			column += 1
			data[row][column] = item{char, false, false}
			if unicode.IsDigit(char) {
				data[row][column].isDigit = true
				if !inPart {
					partStarts = append(partStarts, point{x: column, y: row})
					inPart = true
				}
			} else {
				if inPart {
					partEnd = true
					inPart = false
				}
			}
			if partEnd {
				partEnds = append(partEnds, point{x: column, y: row})
			}
		}
		if len(partStarts) != len(partEnds) {
			fmt.Printf("Problem after row %d\n", row-1)
		}
	}

	schematic := schematic{rows: rows, columns: columns, data: data, partStarts: partStarts, partEnds: partEnds}
	for row := 1; row < rows-1; row++ {
		for column := 1; column < columns-1; column++ {
			schematic.setSymbolAdjacency(row, column)
		}
	}

	return schematic
}

func (s *schematic) setSymbolAdjacency(row int, column int) {
	if unicode.IsDigit(s.data[row][column].value) {
		for y := row - 1; y <= row+1; y++ {
			for x := column - 1; x <= column+1; x++ {
				if s.data[y][x].value != '.' && !unicode.IsDigit(s.data[y][x].value) {
					s.data[row][column].symbolAdjacent = true
					break
				}
			}
		}
	}
}
