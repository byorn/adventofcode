package main

import (
	"fmt"
	util "github.com/byorn/AdventOfCode/util"
	"testing"
)

func TestDay1(t *testing.T) {

	weights := []string{"1", "1", "1", "", "2", "2"}
	maximumInGroup := getMaximumInGroup(weights)
	if maximumInGroup != 4 {
		t.Errorf("Max in Group is not 4, received %v", maximumInGroup)

	}
}

func TestReadFile(t *testing.T) {
	lines := util.ReadFileContents("day1.txt")
	fmt.Println(lines)
}

func TestShouldGetValue72478(t *testing.T) {
	lines := readFileContents("day1.txt")
	maximumInGroup := getMaximumInGroup(lines)
	if maximumInGroup != 72478 {
		t.Errorf("Max in Group is not 72478, received %v", maximumInGroup)
	}
}
