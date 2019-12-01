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
        sum += math.Floor(n/3.0)-2.0
    }
    fmt.Println(int(sum))
}
