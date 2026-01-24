package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	common "doppie.com/aoc-common"
)


func parseInput(i string) int {
    d := i[0]
    if d != 'R' && d != 'L' {
        return 0
    }
    n := i[1:]

    ni, err := strconv.Atoi(n)
    if err != nil {
        log.Fatal(err)
    }
    if d != 'R' {
        ni = -ni
    }

    return ni
}

func fancyRingModulo(a, b int) int {
    a = a % b

    for a < 0 {
        a = (a % b + b) % b
    }

    return a
}

func rotate(r int, i string) int {
    a := parseInput(i)

    t := fancyRingModulo(r + a, 100)
    
    return t
}

func unlockSafe(i []string) int {
    r := 50
    c := 0
    fmt.Printf("The dial starts pointing at %d.\n", r)
    
    for _, element := range i {
        if len(element) == 0 {
            continue
        }
        r = rotate(r, element)
        fmt.Printf("The dial is rotated %s to point at %d.\n", element, r)
        if r == 0 {
            c++
        }
    } 

    return c
}

func partOne(f string) int {
    contents := common.ReadFile(f)

    return unlockSafe(contents)
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    passOne := partOne(f)

    fmt.Printf("Password is %d.", passOne)
}
