package main

import (
	"testing"
)

func TestShouldGetValue8392(t *testing.T) {
	lines := readFileContents("day2.txt")
	total := getTotalScores(lines)
	if total != 8392 {
		t.Errorf("Total score is not 8392, received %v", total)
	}
}

func TestShouldGetValue10116(t *testing.T) {
	lines := readFileContents("day2.txt")
	total := getTotalScoresPart2(lines)
	if total != 10116 {
		t.Errorf("Total score is not 10116, received %v", total)
	}
}
