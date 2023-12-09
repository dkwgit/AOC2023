package day04

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	set "github.com/deckarep/golang-set/v2"

	util "github.com/dkwgit/aoc2023/utility"
)

type card struct {
	id      int
	wins    set.Set[string]
	numbers set.Set[string]
	count   int
}

func (c *card) score() int {
	score := 0
	intersect := c.numbers.Intersect(c.wins)
	count := intersect.Cardinality()
	if count >= 0 {
		score = int(math.Pow(2, float64(count-1)))
	}
	return score
}

var cardsMap map[*util.DayContext][]card = make(map[*util.DayContext][]card)

func Puzzle01(dc *util.DayContext) string {
	cards, ok := cardsMap[dc]
	if !ok {
		winningCardCount, err := strconv.Atoi(dc.Args[0])
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", dc.Args[0], err))
		}
		cardCount, err := strconv.Atoi(dc.Args[1])
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", dc.Args[1], err))
		}
		cards = loadCards(dc.InputLines, winningCardCount, cardCount)
		cardsMap[dc] = cards
	}

	var result = 0
	for _, card := range cards {
		result += card.score()
	}
	return fmt.Sprintf("day04-01: %d", result)
}

func Puzzle02(dc *util.DayContext) string {
	cards, ok := cardsMap[dc]
	if !ok {
		winningCardCount, err := strconv.Atoi(dc.Args[0])
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", dc.Args[0], err))
		}
		cardCount, err := strconv.Atoi(dc.Args[1])
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in Puzzle01: %s", dc.Args[1], err))
		}
		cards = loadCards(dc.InputLines, winningCardCount, cardCount)
		cardsMap[dc] = cards
	}

	for cardIndex := 0; cardIndex < len(cards); cardIndex++ {
		intersect := cards[cardIndex].numbers.Intersect(cards[cardIndex].wins)
		wins := intersect.Cardinality()
		if wins >= 0 {
			maxIndex := int(math.Min((float64)(cardIndex+wins), (float64)(len(cards)-1)))
			for forwardIndex := cardIndex + 1; forwardIndex <= maxIndex; forwardIndex++ {
				cards[forwardIndex].count += cards[cardIndex].count
			}
		}
	}

	var result = 0
	for cardIndex := 0; cardIndex < len(cards); cardIndex++ {
		result += cards[cardIndex].count
	}

	return fmt.Sprintf("day04-02: %d", result)
}
func loadCards(inputLines []string, winCount int, numberCount int) []card {
	pattern := fmt.Sprintf("^Card +(\\d{1,3}): +((?:\\d{1,2} +){%d}\\d{1,2}) +\\| +((?:\\d{1,2} +){%d}\\d{1,2}).*$", winCount-1, numberCount-1)
	reLine := regexp.MustCompile(pattern)
	reMultiSpaceSeparator := regexp.MustCompile(" +")
	var cards = make([]card, len(inputLines))

	for i, line := range inputLines {
		matches := reLine.FindAllStringSubmatch(line, -1)
		matchSet := matches[0]
		cardString := matchSet[1]
		winningNumbers := reMultiSpaceSeparator.Split(matchSet[2], -1)
		numbers := reMultiSpaceSeparator.Split(matchSet[3], -1)
		cardNumber, err := strconv.Atoi(cardString)
		if err != nil {
			panic(fmt.Sprintf("Error converting string %s to int in loadCards: %s", cardString, err))
		}
		cards[i] = card{
			id:      cardNumber,
			wins:    set.NewSet[string](winningNumbers...),
			numbers: set.NewSet[string](numbers...),
			count:   1,
		}
	}

	return cards
}
