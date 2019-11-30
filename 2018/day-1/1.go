package main

import (
    "log"
    "fmt"
    "bufio"
    "os"
)

func main() {
    f, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    sum := 0
    n := 0
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &n)
        sum += n
    }
    fmt.Println(sum)
}
