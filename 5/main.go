package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

func isFresh(a int, rangeArr [][2]int) bool {
	for i := range len(rangeArr) {
		if a >= rangeArr[i][0] && a <= rangeArr[i][1] {
			return true
		}
	}

	return false
}

func countSpoiledIngredients(c []string) int {
	sum := 0
	checkMode := false
	rangeArr := make([][2]int, 0)


	for _, l := range c {
		if l == "" {
			checkMode = true
			continue
		}

		if !checkMode {
			idRange := strings.Split(l, "-")
			rangeStart, err := strconv.Atoi(idRange[0])
			if err != nil {
				log.Fatalf("Conversion (%s) failed!", idRange[0])
			}
			rangeEnd, err := strconv.Atoi(idRange[1])
			if err != nil {
				log.Fatalf("Conversion (%s) failed!", idRange[1])
			}
			rangeArr = append(rangeArr, [2]int{rangeStart, rangeEnd})
		} else {
			num, err := strconv.Atoi(l)
			if err != nil {
				log.Fatalf("Conversion (%s) failed!", l)
			}
			if isFresh(num, rangeArr) {
				sum++
			}
		}
	}

	return sum
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	sum := countSpoiledIngredients(contents)

	log.Printf("%d ingredients are fresh.", sum)
}