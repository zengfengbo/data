package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func isName(line string) bool {
	return line[0] == '>'
}

func printMap(m map[rune]int, line string) {
	for c, n := range m {
		fmt.Printf("%s\t%c\t%d\n", line, c, n)
	}
}

func countLine(i map[rune]int, line string) map[rune]int {
	o := i
	for _, c := range line {
		o[c] += 1
	}
	return o
}

func parseFile(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	// scanner.Split(bufio.ScanLines)
	out := make(map[rune]int)
	line := ""
	name := ""

	for scanner.Scan() {
		line = strings.ToUpper(scanner.Text())
		if isName(line) {
			printMap(out, name)
			name = line
			out = make(map[rune]int)
		} else {
			out = countLine(out, line)
		}
	}
	printMap(out, name)
}

func main() {
	flag.Parse()
	f := flag.Arg(0)
	// fmt.Println(f)
	parseFile(f)
}
