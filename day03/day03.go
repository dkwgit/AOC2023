package day03

import (
	"fmt"
	"strconv"
	"unicode"

	set "github.com/deckarep/golang-set/v2"

	util "github.com/dkwgit/aoc2023/utility"
)

type point struct {
	x     int
	y     int
	value string
}

type item struct {
	value          rune
	isDigit        bool
	isGear         bool
	symbolAdjacent bool
	partId         int
}

type part struct {
	start          point
	end            point
	value          int
	id             int
	symbolAdjacent bool
}

type schematic struct {
	rows    int
	columns int
	data    [][]item
	parts   []part
}

var schematicMap map[*util.DayContext]*schematic = make(map[*util.DayContext]*schematic)

func Puzzle01(dc *util.DayContext) string {
	schematic, ok := schematicMap[dc]
	if !ok {
		schematic = createSchematic(dc.InputLines)
		schematicMap[dc] = schematic
	}
	var realParts = make([]int, 0)

	for _, part := range schematic.parts {
		var adjacency bool = false
		for i := part.start.x; i < part.end.x; i++ {
			if schematic.data[part.start.y][i].symbolAdjacent {
				adjacency = true
			}
		}

		if adjacency {
			realParts = append(realParts, part.value)
		}
	}

	var result = util.SumIntArray(realParts)

	return fmt.Sprintf("day03-01: %d", result)
}

func Puzzle02(dc *util.DayContext) string {
	schematic, ok := schematicMap[dc]
	if !ok {
		schematic = createSchematic(dc.InputLines)
		schematicMap[dc] = schematic
	}

	var result = 0
	for row := 0; row < schematic.rows; row++ {
		for column := 0; column < schematic.columns; column++ {
			set := set.NewSet[int]()
			if schematic.data[row][column].isGear {
				for y := row - 1; y <= row+1; y++ {
					for x := column - 1; x <= column+1; x++ {
						if schematic.data[y][x].partId != -1 {
							set.Add(schematic.data[y][x].partId)
						}
					}
				}
				if set.Cardinality() == 2 {
					var slice = set.ToSlice()
					id1 := slice[0]
					id2 := slice[1]
					var value1 int
					var value2 int
					for i := 0; i < len(schematic.parts); i++ {
						if schematic.parts[i].id == id1 {
							value1 = schematic.parts[i].value
						}
						if schematic.parts[i].id == id2 {
							value2 = schematic.parts[i].value
						}
					}
					result += value1 * value2
				}
			}
		}
	}

	return fmt.Sprintf("day03-02: %d", result)
}

func createSchematic(inputLines []string) *schematic {
	var parts []part = make([]part, 0)
	// Pad the input with empty rows and columns
	var rows = len(inputLines) + 2
	var columns = len(inputLines[0]) + 2

	var emptyRow = make([]item, columns)
	for i := 0; i < columns; i++ {
		emptyRow[i] = item{'.', false, false, false, -1}
	}

	var data = make([][]item, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]item, columns)
		copy(data[i], emptyRow)
	}

	var nextPartId = 1
	for row, line := range inputLines {
		var inPart bool = false
		var partEnd bool = false
		line += "."
		var p part
		var partId int = -1
		for column, char := range line {
			partEnd = false
			isDigit := false
			isGear := false
			if unicode.IsDigit(char) {
				isDigit = true
				if !inPart {
					partId = nextPartId
					nextPartId++
					p = part{
						point{x: column + 1, y: row + 1, value: string(char)},
						point{},
						-1,
						partId,
						false,
					}
					inPart = true
				}
			} else {
				if char == '*' {
					isGear = true
				}
				if inPart {
					partEnd = true
					inPart = false
					partId = -1
				}
			}
			data[row+1][column+1] = item{char, isDigit, isGear, false, partId}
			if partEnd {
				p.end = point{x: column + 1, y: row + 1, value: string(char)}
				var partString string = ""
				for i := p.start.x; i < p.end.x; i++ {
					partString += string(data[p.start.y][i].value)
				}
				p.value, _ = strconv.Atoi(partString)
				parts = append(parts, p)
			}
		}
	}

	schematic := schematic{rows: rows, columns: columns, data: data, parts: parts}
	for row := 1; row < rows-1; row++ {
		for column := 1; column < columns-1; column++ {
			schematic.setSymbolAdjacency(row, column)
		}
	}

	return &schematic
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
