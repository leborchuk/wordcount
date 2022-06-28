// tool for count the number of words in a string.
// Should be launched as ./wordcount 'go is awesome'.
package main

import (
	"fmt"
	"os"
	"strings"
)

func countWords(wordString string) (int, error) {
	return len(strings.Fields(wordString)), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./wordcount 'go is awesome'")
		os.Exit(1)
	}
	var wordsCount, err = countWords(os.Args[1])
	if err != nil {
		fmt.Printf("An error occured %v\n", err)
		os.Exit(1)
	}
	fmt.Println(wordsCount)
}

