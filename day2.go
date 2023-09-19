package main

import (
	"strings"
)

var part2MatchCombinations = map[string][]string{
	"AX": {"L", "Z"},
	"AY": {"D", "X"},
	"AZ": {"W", "Y"},
	"BX": {"L", "X"},
	"BY": {"D", "Y"},
	"BZ": {"W", "Z"},
	"CX": {"L", "Y"},
	"CY": {"D", "Z"},
	"CZ": {"W", "X"},
}
var matchCombinations = map[string]string{
	"AX": "D",
	"AY": "W",
	"AZ": "L",
	"BX": "L",
	"BY": "D",
	"BZ": "W",
	"CX": "W",
	"CY": "L",
	"CZ": "D",
}

var pointsOptionSelected = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var pointsGameResult = map[string]int{
	"W": 6,
	"D": 3,
	"L": 0,
}

func getTotalScores(handPlay []string) int {

	handTotal := 0
	outcomeTotal := 0
	for _, play := range handPlay {
		var replacedSpace = strings.Replace(play, " ", "", 1)
		result := getOutComeAndHandPlayedPoints(replacedSpace)
		handTotal += result[0]
		outcomeTotal += result[1]
	}

	return handTotal + outcomeTotal
}

func getTotalScoresPart2(handPlay []string) int {

	handTotal := 0
	outcomeTotal := 0
	for _, play := range handPlay {
		var replacedSpace = strings.Replace(play, " ", "", 1)
		play = replacedSpace[:1] + part2MatchCombinations[replacedSpace][1]

		result := getOutComeAndHandPlayedPoints(play)
		handTotal += result[0]
		outcomeTotal += result[1]
	}

	return handTotal + outcomeTotal
}

func getOutComeAndHandPlayedPoints(play string) []int {
	hand := 0
	var outcome = 0
	result := []int{0, 0}
	matchResult := matchCombinations[play]
	outcome = pointsGameResult[matchResult]
	hand = pointsOptionSelected[play[1:]]
	result[0] = hand
	result[1] = outcome
	return result
}
