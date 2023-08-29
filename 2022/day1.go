package main

import (
	"fmt"
	"strconv"
)

func main() {

}
func getMaximumInGroup(elfs []string) int {
	var maximumFromAllGroups int
	currentElf := 1
	var maximumInCurrentGroup = 0
	for index, w := range elfs {
		if w == "" {
			if maximumInCurrentGroup > maximumFromAllGroups {
				maximumFromAllGroups = maximumInCurrentGroup
			}
			currentElf++
			maximumInCurrentGroup = 0

		} else {
			elfWeight, err := strconv.Atoi(w)
			if err == nil {
				maximumInCurrentGroup += elfWeight
			} else {
				fmt.Println("Couldnt convert ", w, " to an int")
			}
			//if last category, calculate maximum
			if index == len(elfs)-1 {
				if maximumInCurrentGroup > maximumFromAllGroups {
					maximumFromAllGroups = maximumInCurrentGroup
				}
			}
		}
	}
	return maximumFromAllGroups
}
