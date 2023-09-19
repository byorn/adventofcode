package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileContents(fileName string) []string {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}
