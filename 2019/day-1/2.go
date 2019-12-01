package main

import (
    "log"
    "fmt"
    "bufio"
    "os"
	"math"
)

func main() {
    f, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    sum := 0.0
    n := 0.0
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%f", &n)
		_, sum = calc(n, sum)
    }
    fmt.Println(int(sum))
}

func calc(mass, sum float64) (m, s float64) {
	if mass == 0.0 {
		return mass, sum
	}
	s = sum
	m = math.Floor(mass/3.0)-2.0
	if m < 0.0 {
		m = 0.0
	}
	s += m
	return calc(m, s)
}
