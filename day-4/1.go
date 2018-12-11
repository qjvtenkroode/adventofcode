package main

import (
    "fmt"
    "regexp"
    "os"
    "bufio"
    "log"
    "time"
    "sort"
    "strconv"
)

/* 
TODO create struct which combines GuardID with shift start, falls asleep, wakes up, duration sleeping
TODO create map with GuardID, sum of sleep per min from 00:00 until 00:59 (array size 60)
TODO write routine to calculate the map
*/

type Event struct {
    Time time.Time
    Message string
    Guard int
}

type TimeSlice []Event

func (p TimeSlice) Len() int {
    return len(p)
}

func (p TimeSlice) Less(i, j int) bool {
    return p[i].Time.Before(p[j].Time)
}

func (p TimeSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func main() {
    f, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal("error reading file: %s", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    events := make([]Event, 1127)
    format := "2006-01-02 15:04"

    re := regexp.MustCompile("\\[([^\\]]+)\\](.*)$")

    for scanner.Scan() {
        line := scanner.Text()
        match := re.FindStringSubmatch(line)
        t, err := time.Parse(format, match[1])
        if err != nil {
            log.Fatal("error parsing time")
        }
        m := match[2]
        e := Event{Time: t, Message: m}
        fmt.Println(e)
        events = append(events, e)
    }
    sort.Sort(TimeSlice(events))
    guardid := 0
    re = regexp.MustCompile("Guard\\s\\#(\\d+)\\s")
    for i := range events {
        match := re.FindStringSubmatch(events[i].Message)
        if len(match) > 1 {
            guardid, err = strconv.Atoi(match[1])
            if err != nil {
                log.Fatal("string to int conversion failed")
            }
        }
        events[i].Guard = guardid
        fmt.Println(events[i])
    }
}
