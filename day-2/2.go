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
    var ids [250]string
    i := 0

    for scanner.Scan() {
        d := scanner.Text()
        ids[i] = string(d)
        i++
    }
    length := len(ids)
    index := 0
    var s1 string
    for i = 0; i < length; i++ {
        missing := false
        missing_twice := false
        for j := 1; j < length; j++ {
            x := ids[i]
            y := ids[j]
            s1 = ids[i]
            missing = false
            missing_twice = false
            fmt.Println(x, y)
            for o := 0; o < len(x); o++ {
                c := strings.Compare(string(x[o]),string(y[o]))
                fmt.Println(c)
                if c != 0 {
                    if missing {
                        missing_twice = true
                        break
                    }
                    missing = true
                    index = o
                }
            }
            fmt.Println(missing, missing_twice)
            if missing && !missing_twice {
                break
            }
        }
        if missing && !missing_twice {
            break
        }
    }
    for i = 0; i < len(s1); i++ {
        if i != index {
            fmt.Print(string(s1[i]))
        }
    }
}
