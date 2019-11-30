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
    fabric := [1000][1000]map[int]bool{}
    for a:=0;a<1000;a++{
        for b:=0;b<1000;b++{
            fabric[a][b] = make(map[int]bool)
        }
    }

    f, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    re := regexp.MustCompile("\\#(\\d+)\\s\\@\\s(\\d+)\\,(\\d+)\\:\\s(\\d+)x(\\d+)$")

    for scanner.Scan() {
        line := scanner.Text()
        match := re.FindStringSubmatch(line)
        id,err := strconv.Atoi(match[1])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        i,err := strconv.Atoi(match[2])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        j,err := strconv.Atoi(match[3])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        x,err := strconv.Atoi(match[4])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        y,err := strconv.Atoi(match[5])
        if err != nil {
            log.Fatal("error converting:",err)
        }
        h,v := i+x,j+y
        for p:=j;p<v;p++ {
            for l:=i;l<h;l++ {
                fabric[p][l][id] = true
            }
        }
    }
    
    for id:=1;id<1306;id++{ 
        m := make(map[int]bool)
        for a:=0;a<1000;a++{
            for b:=0;b<1000;b++{
                length := len(fabric[a][b])
                if fabric[a][b][id] {
                    m[length] = true
                }
            }
        }
        if len(m) == 1 && m[1] {
            fmt.Println(id)
        }
    }
}
