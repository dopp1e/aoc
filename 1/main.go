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

func moduloWithQuotient(a, d int) (int, int) {
    r := a % d
    b := a - r
    return r, (b / d)
}

func fancyRingModulo(a, b int) int {
    a = a % b

    for a < 0 {
        a = (a % b + b) % b
    }

    return a
}

func rotate(r int, i string, alt bool) (int, int) {
    a := parseInput(i)

    c := 0

    if alt {
        step := 1
        if a < 0 {
            step = -1
        }

        for j := r; j != r + a; j += step {
            if fancyRingModulo(j, 100) == 0 {
                c++
            }
        }
    }

    t := fancyRingModulo(r + a, 100)
    
    return t, c
}

func unlockSafe(i []string, altPass bool) int {
    r := 50
    c := 0
    fmt.Printf("The dial starts pointing at %d.\n", r)
    
    for _, element := range i {
        if len(element) == 0 {
            continue
        }
        x, y := rotate(r, element, altPass)
        r = x
        fmt.Printf("The dial is rotated %s to point at %d.\n", element, r)
        if altPass {
            c += y
        } else {
            if r == 0 {
                c++
            }
        }
    } 

    return c
}

func partOne(f string) int {
    contents := common.ReadFile(f)

    return unlockSafe(contents, false)
}

func partTwo(f string) int {
    contents := common.ReadFile(f)

    return unlockSafe(contents, true)
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    passOne := partOne(f)

    fmt.Printf("Password is %d.", passOne)

    passTwo := partTwo(f)

    fmt.Printf("Alternative password is %d.", passTwo)
}
