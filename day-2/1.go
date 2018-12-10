package main

import (
    "log"
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    f, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    two := 0
    three := 0

    for scanner.Scan() {
        found_two := false
        found_three := false
        d := scanner.Text()
        for i := 0; i < len(d);i++ {
            c := string(d[i])
            count := strings.Count(d, c)
            if count == 2  && !found_two {
                two++
                found_two = true
            } else if count == 3 && !found_three {
                three++
                found_three = true
            }
            if found_two && found_three {
                break
            }
        }
        //fmt.Printf("Two: %d; Three %d .\n", two, three)
    }
    fmt.Printf("Total: %d", two*three)
}
