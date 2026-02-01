package main

import (
	"fmt"
	"log"
	"math"
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

func isMathCharacter(c rune) bool {
	return c == '+' || c == '*'
} 

func getRanges(s string) [][2]int {
	r := make([][2]int, 0)
	
	curStart := 0

	for i := 1; i < len(s); i++ {
		if (isMathCharacter(rune(s[i]))) {
			r = append(r, [2]int{curStart, i - 1})
			curStart = i
		}
	}

	return append(r, [2]int{curStart, len(s)})
}

func assembleNumber(c []string, row int) int {
	sum := 0
	power := 0

	for i := (len(c) - 2); i >= 0; i-- {
		r := rune(c[i][row])
		if r == ' ' {
			continue
		}
		n := int(r - '0')
		fmt.Printf("%d\n", n)
		sum += n * int(math.Pow(10, float64(power)))
		power++
	}

	return sum
}

func betterHomework(c []string) int {
	sum := 0
	opId := len(c) - 1

	r := getRanges(c[opId])

	for i := 0; i < len(r); i++ {
		op := c[opId][r[i][0]]
		v := 0
		if op == '*' {
			v = 1
		}
		numCount := r[i][1] - r[i][0]
		// instructions say left to right
		// it does not really matter for this, but for the sake of compatibility
		// with other operations (subtraction, division), it'll be done that way
		for j := 0; j < numCount; j++ {
			num := assembleNumber(c, r[i][1] - j - 1)
			fmt.Printf("Total: %d\n\n", num)
			if op == '*' {
				v *= num
			} else {
				v += num
			}
		}
		sum += v
	}

	return sum
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	a := doMathHomework(contents)

	fmt.Printf("Grand total equals %d.\n", a)

	b := betterHomework(contents)

	fmt.Printf("Better total equals %d.\n", b)
}