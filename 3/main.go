package main

import (
	"fmt"
	"os"

	"doppie.com/aoc-common"
)

func findBestBank(s string) int {
    best := 0

    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s); j++ {
            firstDigit := s[i]
            secondDigit := s[j]

            number := int(firstDigit - '0') * 10 + int(secondDigit - '0')

            if number > best {
                best = number
            }
        }
    }

    return best
}

func findMaxJoltage(c []string) int {
    sum := 0

    for _, v := range c {
        bank := findBestBank(v)
        fmt.Printf("Found bank: %d.\n", bank)

        sum += bank
    }

    return sum
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)

    j := findMaxJoltage(contents)

    fmt.Printf("Found max joltage: %d.", j)
}
