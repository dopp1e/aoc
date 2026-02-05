package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	common "doppie.com/aoc-common"
)

// type for keeping distances between junction boxes
type distance struct {
	s1 string
	s2 string
	dist float64
}

// helper type for sorting distances
type DistanceSort []distance

func (a DistanceSort) Len() int           { return len(a) }
func (a DistanceSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DistanceSort) Less(i, j int) bool { return a[i].dist < a[j].dist }

// given string "x,y,z" unpack and turn into ints
func unpackCoordinates(s string) (int, int, int) {
	spl := strings.Split(s, ",")
	x, err := strconv.Atoi(spl[0])
	if err != nil {
		log.Fatalf("Failed converting first element of %s to int!\n", s)
	}
	y, err := strconv.Atoi(spl[1])
	if err != nil {
		log.Fatalf("Failed converting second element of %s to int!\n", s)
	}
	z, err := strconv.Atoi(spl[2])
	if err != nil {
		log.Fatalf("Failed converting third element of %s to int!\n", s)
	}
	return x, y, z
}

// calculate distance between two sets of x,y,z coordinates
func euclideanDistance(x1, y1, z1, x2, y2, z2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2) - float64(x1), 2) + math.Pow(float64(y2) - float64(y1), 2)+ math.Pow(float64(z2) - float64(z1), 2))
}

func putTogether(s1, s2 string) string {
	return s1 + ";" + s2
}

// func isInString(s1, s2 string, check distance) bool {
// 	firstOrder := s1 == check.s1 && s2 == check.s2
// 	secondOrder := s1 == check.s2 && s2 == check.s1
// 	return firstOrder || secondOrder  
// }

func areAlreadyInList(s1, s2 string, passed map[string]bool) bool {
	_, ok1 := passed[putTogether(s1, s2)]
	_, ok2 := passed[putTogether(s2, s1)]
	return ok1 || ok2
}

func countIDs(m map[string]int) int {
	arr := make([]int, 0)

	for k := range m {
		if slices.Contains(arr, m[k]) {
			continue
		}

		arr = append(arr, m[k])
	}

	return len(arr)
}

func connectBoxes(c []string, connections int, ext bool) int {
	m := 1

	// calculate all distances between points...
	list := make([]distance, 0)
	passedCache := make(map[string]bool)
	
	for i := 0; i < len(c); i++ {
		x1, y1, z1 := unpackCoordinates(c[i])
		for j := 0; j < len(c); j++ {
			// ...except between the same points
			if i == j {
				continue
			}
			// if already calculated, skip
			if areAlreadyInList(c[i], c[j], passedCache) {
				continue
			}

			x2, y2, z2 := unpackCoordinates(c[j])
			dist := euclideanDistance(x1, y1, z1, x2, y2, z2)
			list = append(list, distance{s1: c[i], s2: c[j], dist: dist})
			passedCache[putTogether(c[i], c[j])] = true
		}
	}

	// we have a list of distances between junction boxes sorted by distance
	sort.Sort(DistanceSort(list))

	// keeping connectionIDs
	connectionMap := make(map[string]int)
	// keep track of which ones have been assigned
	assignMap := make(map[string]bool)
	currentIndex := 0

	for i := 0; connections > 0 || ext; i++ {
		if i >= len(list) {
			break
		}
		cur := list[i]
		c1, ok1 := connectionMap[cur.s1]
		c2, ok2 := connectionMap[cur.s2]
		if ok1 && ok2 {
			// replace all IDs in second one with the values of the first one
			for k := range connectionMap {
				if connectionMap[k] == c2 {
					connectionMap[k] = c1
				}
			}
			connections--
			if (ext && len(assignMap) == len(c) && countIDs(connectionMap) == 1) {
				x1, _, _ := unpackCoordinates(cur.s1)
				x2, _, _ := unpackCoordinates(cur.s2)
				return x1 * x2 
			}
			continue
		}
		// if only one has been connected
		if ok1 != ok2 {
			// add whichever one is not connected to where the other one is connected
			if ok1 {
				assignMap[cur.s2] = true
				connectionMap[cur.s2] = c1
				//fmt.Printf("Adding %s to %d.\n", cur.s2, c1)
			} else if ok2 {
				connectionMap[cur.s1] = c2
				assignMap[cur.s1] = true
				//fmt.Printf("Adding %s to %d.\n", cur.s1, c2)
			}
			connections--
			if (ext && len(assignMap) == len(c) && countIDs(connectionMap) == 1) {
				x1, _, _ := unpackCoordinates(cur.s1)
				x2, _, _ := unpackCoordinates(cur.s2)
				return x1 * x2 
			}
			continue
		}
		// if neither is connected
		// give them indices
		connectionMap[cur.s1] = currentIndex
		connectionMap[cur.s2] = currentIndex
		// set their assignments
		assignMap[cur.s1] = true
		assignMap[cur.s2] = true
		currentIndex++
		//fmt.Printf("New connection between %s and %s made.\n", cur.s1, cur.s2)
		connections--
		if (ext && len(assignMap) == len(c) && countIDs(connectionMap) == 1) {
			x1, _, _ := unpackCoordinates(cur.s1)
			x2, _, _ := unpackCoordinates(cur.s2)
			return x1 * x2 
		}
	}

	// count instances of numbers
	countMap := make(map[int]int)

	for k := range connectionMap {
		_, ok := countMap[connectionMap[k]]
		if ok {
			countMap[connectionMap[k]]++
		} else {
			countMap[connectionMap[k]] = 1
		}
	}

	allConnections := make([]int, 0)

	for k := range countMap {
		//fmt.Printf("%d = %d\n", k, countMap[k])
		allConnections = append(allConnections, countMap[k])
	}

	slices.Sort(allConnections)

	for i := 0; i < min(3, len(allConnections)); i++ {
		id := len(allConnections) - 1 - i
		//fmt.Printf("Multiplying %d by %d.\n", m, allConnections[id])
		m *= allConnections[id]
	}

	return m
}

func main() {
    argsWithoutProg := os.Args[1:]
    f := argsWithoutProg[0]

    contents := common.ReadFile(f)[1:]

	c := connectBoxes(contents, 10, false)

	fmt.Printf("The size of the circuit with 10 connections equals %d!\n", c)

	d := connectBoxes(contents, 1000, false)

	fmt.Printf("The size of the circuit with 1000 connections equals %d!\n", d)

	e := connectBoxes(contents, 1000, true)

	fmt.Printf("The actual size of the circuit equals %d!\n", e)
}