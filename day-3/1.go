package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "regexp"
    "strconv"
)

func main() {
    fabric := [1000][1000]int{}
    sum := 0

    f, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    re := regexp.MustCompile("\\#\\d+\\s\\@\\s(\\d+)\\,(\\d+)\\:\\s(\\d+)x(\\d+)$")

    for scanner.Scan() {
        line := scanner.Text()
        match := re.FindStringSubmatch(line)
        i,err := strconv.Atoi(match[1])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        j,err := strconv.Atoi(match[2])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        x,err := strconv.Atoi(match[3])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        y,err := strconv.Atoi(match[4])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        h,v := i+x,j+y
        for p:=j;p<v;p++ {
            for l:=i;l<h;l++ {
                fabric[p][l] += 1
            }
        }
    }
    for q := 0; q < 1000; q++ {
        //fmt.Println(fabric[q])
        for o :=0; o < 1000; o++ {
            if fabric[q][o] > 1 {
                sum += 1
            }
        }
    }
    fmt.Println(sum)
}
