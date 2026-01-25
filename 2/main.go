package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

func extractRange(idr string) (int, int) {
    ranges := strings.Split(idr, "-")
    if len(ranges) != 2 {
        log.Fatal("Incorrect range while extracting - aborting.")
    }

    start, err := strconv.Atoi(ranges[0])
    if err != nil {
        log.Fatal(err)
    }

    end, err := strconv.Atoi(ranges[1])
    if err != nil {
        log.Fatal(err)
    }

    return start, end
}

func isValid(id int) bool {
    s := strconv.Itoa(id)

    if len(s) % 2 != 0 {
        return true
    }

    first := s[0:len(s)/2]
    second := s[len(s)/2:]


    if first == second {
        fmt.Printf("%s = %s + %s\n", s, first, second)
        return false
    }

    return true
}

func isComposedOfSequences(id int) bool {
    s := strconv.Itoa(id)
    sLen := len(s)

    
    for i := 0; i < sLen / 2; i++ {
        part := s[0:i+1]
        partLen := len(part)
        doReturn := false
        for j := 0; j < sLen; j += partLen {
            if (j+partLen) > sLen {
                doReturn = false
                continue
            }
            checkPart := s[(j):(j+partLen)]
            if (checkPart != part) {
                doReturn = false
                break
            }
            doReturn = true
        }
        if (doReturn) {
            return true
        }
    }

    return false
}

func findInvalidIDs(start, end int) []int {
    a := make([]int, 0)

    for i := start; i <= end; i++ {
        if !isValid(i) {
            a = append(a, i)
        } 
    }

    return a
}

func findInvalidSequences(start, end int) []int {
    a := make([]int, 0)

    for i := start; i <= end; i++ {
        if isComposedOfSequences(i) {
            a = append(a, i)
        }
    }

    return a
}


func identify(c string) (int, int) {
    sum := 0
    sumS := 0

    split := strings.SplitSeq(c, ",")
    for idr := range split {
        start, end := extractRange(idr)

        invalid := findInvalidIDs(start, end)

        for _, v := range invalid {
            sum += v
        }

        invalidSequences := findInvalidSequences(start, end)
        
        for _, w := range invalidSequences {
            sumS += w
        }
    }

    return sum, sumS
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)

    content := strings.Join(contents, "")

    sum, sumS := identify(content)

    fmt.Printf("Sum of invalid IDs equals %d.\n", sum)
    fmt.Printf("Sum of invalid IDs (checked as sequences) equals %d.\n", sumS)
}
