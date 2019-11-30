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


    dup := make(map[int]bool)
    sum := 0
    n := 0
    not_found := true
    dup[sum] = true
    for not_found {
        f.Seek(0,0)
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            fmt.Sscanf(scanner.Text(), "%d", &n)
            sum += n
            if dup[sum] {
                fmt.Println("Found it:", sum)
                not_found = false
                break
            }
            dup[sum] = true
        }
    }
}
