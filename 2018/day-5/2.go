package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "strings"
)

func destroy_polymer(s string) (string, bool) {
    b := false
    length := len(s)
    for i := 0; i < length-1; i++ {
        diff := s[i] - s[i+1]
        diff2 := s[i+1] - s[i]
        if diff == 32 || diff2 == 32 {
            var builder strings.Builder
            b = true
            s1 := ""
            if(i > 0) {
                s1 = s[:i]
            }
            s2 := s[i+2:]
            builder.WriteString(s1)
            builder.WriteString(s2)
            s = builder.String()
            break
        }
    }

    return s, b
}

func main() {
    f, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }

    polymer := string(f)
    abc := "abcdefghijklmnopqrstuwvxyz"
    shortest := 10000000000000
    for i := range(abc) {
        fmt.Println("removing ", string(abc[i]), string(abc[i]-32))

        p := strings.Replace(polymer, string(abc[i]), "", -1)
        p = strings.Replace(p, string(abc[i]-32), "", -1)

        b := true
        for b {
            p, b = destroy_polymer(p)
        }
        length := len(p)-1
        fmt.Println(length)
        if length < shortest {
            shortest = length
        }
    }
    fmt.Println("Shortest polymer: ", shortest)
}
