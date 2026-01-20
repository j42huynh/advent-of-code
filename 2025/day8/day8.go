package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatal("missing input file argument")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		coordStrings := strings.Split(line, ",")

		coord := []int{}
		for _, c := range coordStrings {
			coord = append(coord, getInt(c))
		}
		coords = append(coords, coord)
	}

	fmt.Print("Part1: ")
	fmt.Println(part1(coords))
	fmt.Print("Part2: ")
	fmt.Println(part2(coords))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return num
}

func part1(coords [][]int) int {
	distancesHeap := buildDistancesHeap(coords)

	circuits := []int{}
	coordsIndToCircuitsInd := map[int]int{}
	iterations := 0
	for len(*distancesHeap) > 0 {
		if iterations >= 1000 {
			break
		}
		iterations += 1

		d := heap.Pop(distancesHeap)
		dist, ok := d.(Distance)
		if !ok {
			log.Fatal("popped something not a Distance")
			return -1
		}

		coord1CircuitsInd, ok1 := coordsIndToCircuitsInd[dist.coord1Index]
		coord2CircuitsInd, ok2 := coordsIndToCircuitsInd[dist.coord2Index]
		if ok1 && ok2 {
			if coord1CircuitsInd == coord2CircuitsInd {
				continue
			}

			// merge circuits if coords are part of different circuits
			maxCircuitsInd := max(coord1CircuitsInd, coord2CircuitsInd)
			minCircuitsInd := min(coord1CircuitsInd, coord2CircuitsInd)
			circuits[minCircuitsInd] += circuits[maxCircuitsInd]
			circuits = merge(circuits, maxCircuitsInd, coordsIndToCircuitsInd, coord1CircuitsInd, coord2CircuitsInd)
		} else if ok1 {
			circuits[coord1CircuitsInd] += 1
			coordsIndToCircuitsInd[dist.coord2Index] = coord1CircuitsInd
		} else if ok2 {
			circuits[coord2CircuitsInd] += 1
			coordsIndToCircuitsInd[dist.coord1Index] = coord2CircuitsInd
		} else {
			circuits = append(circuits, 2)
			coordsIndToCircuitsInd[dist.coord1Index] = len(circuits) - 1
			coordsIndToCircuitsInd[dist.coord2Index] = len(circuits) - 1
		}
	}

	slices.Sort(circuits)
	result := 1
	for i := len(circuits) - 1; i >= 0 && i >= len(circuits)-3; i-- {
		result *= circuits[i]
	}
	return result
}

func part2(coords [][]int) int {
	distancesHeap := buildDistancesHeap(coords)

	circuits := []int{}
	coordsIndToCircuitsInd := map[int]int{}
	for len(*distancesHeap) > 0 {
		d := heap.Pop(distancesHeap)
		dist, ok := d.(Distance)
		if !ok {
			log.Fatal("popped something not a Distance")
			return -1
		}

		coord1CircuitsInd, ok1 := coordsIndToCircuitsInd[dist.coord1Index]
		coord2CircuitsInd, ok2 := coordsIndToCircuitsInd[dist.coord2Index]
		if ok1 && ok2 {
			if coord1CircuitsInd == coord2CircuitsInd {
				continue
			}

			// merge circuits if coords are part of different circuits
			maxCircuitsInd := max(coord1CircuitsInd, coord2CircuitsInd)
			minCircuitsInd := min(coord1CircuitsInd, coord2CircuitsInd)
			circuits[minCircuitsInd] += circuits[maxCircuitsInd]
			if circuits[minCircuitsInd] == len(coords) {
				return dist.coord1[0] * dist.coord2[0]
			}
			circuits = merge(circuits, maxCircuitsInd, coordsIndToCircuitsInd, coord1CircuitsInd, coord2CircuitsInd)
		} else if ok1 {
			circuits[coord1CircuitsInd] += 1
			coordsIndToCircuitsInd[dist.coord2Index] = coord1CircuitsInd
			if circuits[coord1CircuitsInd] == len(coords) {
				return dist.coord1[0] * dist.coord2[0]
			}
		} else if ok2 {
			circuits[coord2CircuitsInd] += 1
			coordsIndToCircuitsInd[dist.coord1Index] = coord2CircuitsInd
			if circuits[coord2CircuitsInd] == len(coords) {
				return dist.coord1[0] * dist.coord2[0]
			}
		} else {
			circuits = append(circuits, 2)
			coordsIndToCircuitsInd[dist.coord1Index] = len(circuits) - 1
			coordsIndToCircuitsInd[dist.coord2Index] = len(circuits) - 1
		}
	}

	return -1
}

func buildDistancesHeap(coords [][]int) *MinDistanceHeap {
	distancesHeap := &MinDistanceHeap{}
	heap.Init(distancesHeap)

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			distance := Distance{
				distance:    calculateDistance(coords[i], coords[j]),
				coord1:      coords[i],
				coord1Index: i,
				coord2:      coords[j],
				coord2Index: j,
			}
			heap.Push(distancesHeap, distance)
		}
	}
	return distancesHeap
}

func calculateDistance(coord1 []int, coord2 []int) float64 {
	xDiff := math.Pow(float64(coord1[0]-coord2[0]), 2)
	yDiff := math.Pow(float64(coord1[1]-coord2[1]), 2)
	zDiff := math.Pow(float64(coord1[2]-coord2[2]), 2)
	return math.Sqrt(xDiff + yDiff + zDiff)
}

func merge(circuits []int, indToRemove int, coordsIndToCircuitsInd map[int]int, coord1CircuitsInd int, coord2CircuitsInd int) []int {
	result := slices.Delete(circuits, indToRemove, indToRemove+1)
	for k, v := range coordsIndToCircuitsInd {
		if v == indToRemove {
			if indToRemove == coord1CircuitsInd {
				coordsIndToCircuitsInd[k] = coord2CircuitsInd
			} else {
				coordsIndToCircuitsInd[k] = coord1CircuitsInd
			}
		} else if v > indToRemove {
			coordsIndToCircuitsInd[k] -= 1
		}
	}
	return result
}
