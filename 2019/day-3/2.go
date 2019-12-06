package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord struct {
	X    int
	Y    int
	Step int
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// first line
	_ = scanner.Scan()
	start := coord{0, 0, 0}
	s := strings.Split(scanner.Text(), ",")
	l1 := step(s, start)
	// second line
	_ = scanner.Scan()
	start = coord{0, 0, 0}
	s = strings.Split(scanner.Text(), ",")
	l2 := step(s, start)

	var intersections []int

	sort.SliceStable(l1, func(i, j int) bool {
		return l1[i].X < l1[j].X
	})
	sort.SliceStable(l2, func(i, j int) bool {
		return l2[i].X < l2[j].X
	})
	for _, x := range l1 {
		for _, y := range l2 {
			if y.X > x.X {
				break
			}
			if x.X == y.X && x.Y == y.Y {
				intersections = append(intersections, x.Step+y.Step)
			}
		}
	}
	min := 999999999999
	for _, m := range intersections {
		if m < min {
			min = m
		}
	}
	fmt.Println(min)
}

func step(q []string, s coord) []coord {
	var l []coord
	for _, x := range q {
		direction := string(x[0])
		steps, _ := strconv.Atoi(x[1:])
		for i := 0; i < steps; i++ {
			switch direction {
			case "L":
				s.X = s.X - 1
				s.Step++
				l = append(l, s)
			case "R":
				s.X = s.X + 1
				s.Step++
				l = append(l, s)
			case "U":
				s.Y = s.Y + 1
				s.Step++
				l = append(l, s)
			case "D":
				s.Y = s.Y - 1
				s.Step++
				l = append(l, s)
			default:
				fmt.Println("Fail somewhere")
			}
		}
	}
	return l
}
