package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "strings"
    "strconv"
)

var locations []*Coordinate
var minX, maxX, minY, maxY int

type Coordinate struct {
    X int
    Y int
    Size int
    Infinite bool
}

func manhattan_distance(x, y int, c *Coordinate) int {
    s := abs(x-c.X) + abs(y-c.Y)
    return s
}

func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}

func toInt(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        log.Fatal("error during conversion: ", err)
    }
    return i
}

func neighbour(x,y int) (res *Coordinate) {
    shortest := 1000
    for _, c := range(locations) {
        d := manhattan_distance(x, y, c)
        if d == shortest {
            res = nil
        }
        if d < shortest {
            shortest = d
            res = c
        }
    }
    
    return res
}

func main() {
    f, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    minX, maxX, minY, maxY = 1000, -1000, 1000, -1000
    lines := strings.Split(string(f), "\n")
    lines = lines[:len(lines)-1]
    for _, l := range(lines) {
        coord := strings.Split(l, ", ")
        c := Coordinate{X: toInt(coord[0]), Y: toInt(coord[1])}
        locations = append(locations, &c)
        if c.X < minX {
            minX = c.X
        } else if c.X > maxX {
            maxX = c.X
        }
        if c.Y < minY {
            minY = c.Y
        } else if c.Y > maxY {
            maxY = c.Y
        }
    }
    for x := minX; x < maxX+1; x++ {
        for y := minY; y < maxY+1; y++ {
            n := neighbour(x, y)
            if n != nil {
                if x == minX || x == maxX || y == minY || y == maxY {
                    n.Infinite = true
                }
                n.Size += 1
            }
        }
    }
    max := 0
    for _, c := range(locations) {
        if c.Size > max && !c.Infinite {
            max = c.Size
        }
    }
    fmt.Println("Largest size is: ", max)
}
