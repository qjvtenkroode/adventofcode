package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "strings"
)

func destroy_polymer(s string) (string, bool) {
    b := false
    // TODO destroy first Cc combo in the polymer string
    // B and b differ 32
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

    b := true
    for b {
        polymer, b = destroy_polymer(polymer)
    }
    fmt.Println(len(polymer)-1)
}
