package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"flag"
)

var DICT string = "/usr/share/dict/dict.txt"

func lineIsWord(line string) bool {
	return strings.ToUpper(line) == line
}

func cleanWord(line string) string {
	return strings.Trim(line, "\r\n \t")
}

func FindWord(search string) {
	file, err := os.Open(DICT)
	if err != nil {
		fmt.Println("Error opening file", DICT)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	found := false
	lastWasNil := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break;
		}

		if lastWasNil && line[0] != '\r' {
			if lineIsWord(line) {
				if found {
					return
				} else if search == cleanWord(line) {
					found = true
					fmt.Print(line)
				}
			} else if found {
				fmt.Print(line)
			}
		} else if found {
			fmt.Print(line)
		}

		lastWasNil = line[0] == '\r'
	}
}

func main() {
	flag.Parse()
	words := flag.Args()

	for _, word := range words {
		FindWord(strings.ToUpper(word))
	}
}
