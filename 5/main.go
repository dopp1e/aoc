package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

type SortBy [][2]int
	
func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	} else {
		return a[i][0] < a[j][0]
	}
}

func isFresh(a int, rangeArr [][2]int) bool {
	for i := range len(rangeArr) {
		if a >= rangeArr[i][0] && a <= rangeArr[i][1] {
			return true
		}
	}

	return false
}

func countAllIDs(rangeArr [][2]int) int {
	sort.Sort(SortBy(rangeArr))

	count := 0
	prevEnd := 0
	for i := range len(rangeArr) {
		start := rangeArr[i][0]
		end := rangeArr[i][1] + 1

		if end > prevEnd {
			count += end - max(start, prevEnd)
			prevEnd = end
		}
	}

	return count
}

func countSpoiledIngredients(c []string) (int, [][2]int) {
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

	return sum, rangeArr
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	sum, r := countSpoiledIngredients(contents)

	log.Printf("%d ingredients are fresh.", sum)

	all := countAllIDs(r)

	log.Printf("%d IDs have been declared.", all)
}