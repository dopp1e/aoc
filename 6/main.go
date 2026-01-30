package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

func replaceSpaces(s string, n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteString(" ")
	}

	return strings.ReplaceAll(s, sb.String(), " ")
}

func splitLine(s string) []string {
	for i := 20; i > 0; i-- {
		s = replaceSpaces(s, i)
	} 
	if (s[0] == ' ') {
		s = s[1:]
	}
	// fmt.Printf("%s-\n", s)
	a := strings.Split(strings.TrimSpace(s), " ")
	// for i := range len(a) {
	// 	fmt.Printf("%s-", a[i])
	// }
	// fmt.Print("\n")

	return a
}

func doMathHomework(c []string) int {
	splitRows := make([][]string, 0)
	for i := range len(c) {
		splitRows = append(splitRows, splitLine(c[i]))
	}

	sum := 0
	
	opCount := len(splitRows[0])
	numberCount := len(splitRows) - 1
	for i := range opCount {
		operation := splitRows[numberCount][i]
		localSum := 0
		if operation == "*" {
			localSum = 1
		}
		for j := range numberCount {
			num, err := strconv.Atoi(splitRows[j][i])
			if err != nil {
				log.Fatalf("Error converting %s to an integer!", splitRows[j][i])
			}
			switch operation {
			case "+":
				localSum += num
			case "*":
				localSum *= num
			}
		}
		sum += localSum
	}

	return sum
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	a := doMathHomework(contents)

	fmt.Printf("Grand total equals %d", a)
}