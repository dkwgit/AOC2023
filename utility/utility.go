package utility

import (
	"bufio"
	"os"
)

type DayContext struct {
	InputLines []string
	Name       string
	Args       []string
}

func NewDayContext(name string, lines []string, args []string) *DayContext {
	dc := new(DayContext)
	dc.Name = name
	if lines == nil {
		dc.loadLines()
	} else {
		dc.InputLines = lines
	}
	dc.Args = args
	return dc
}

func SumIntArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func (dc *DayContext) loadLines() {
	var path string = "/workspace/AOC2023/input/" + dc.Name + ".txt"
	file, err := os.Open(path)
	if err != nil {
		panic("Error opening input file " + path)
	}
	defer file.Close()

	dc.InputLines = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dc.InputLines = append(dc.InputLines, line)
		err = scanner.Err()
		if err != nil {
			panic("Error reading input file " + path + " for line " + line)
		}
	}
}
