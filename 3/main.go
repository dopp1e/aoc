package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	common "doppie.com/aoc-common"
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

func makeKey(s string, i int) string {
    return s + "," + strconv.Itoa(i)
}

func findBestDigits(s string, maxLength int) int {
    memo := make(map[string]int)

    var best func(s string, digits int) int
    best = func(s string, digits int) int {
        key := makeKey(s, digits)
        if cached, ok := memo[key]; ok {
            return cached
        }
        if digits == 0 {
            return 0
        }

        if len(s) == digits {
            v, err := strconv.Atoi(s)
            if err != nil {
                log.Fatal("Conversion failed!")
            }
            return v
        }

        a := int(s[0]-'0')*int(math.Pow(10, float64(digits-1))) + best(s[1:], digits-1)
        b := best(s[1:], digits)
        v := max(a, b)

        memo[key] = v
        return v
    }
    
    return best(s, maxLength)
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

func findLargeJoltage(c []string) int {
    sum := 0

    for _, v := range c {
        set := findBestDigits(v, 12)
        fmt.Printf("Found set: %d.\n", set)
        sum += set
    }

    return sum
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)

    j := findMaxJoltage(contents)

    fmt.Printf("Found max joltage: %d.\n", j)

    newJ := findLargeJoltage(contents[1:])

    fmt.Printf("Found max new joltage: %d.\n", newJ)
}
