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

type Event struct {
    Time time.Time
    Message string
}

type Sleep struct {
    Guard int
    From time.Time
    To time.Time
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

    roster := make(map[int][]int)
    events := make([]Event, 0)
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
        events = append(events, e)
    }
    sort.Sort(TimeSlice(events))
    sleepdata := make([]Sleep, 419) 
    counter := 0
    guardid := 0
    re = regexp.MustCompile("Guard\\s\\#(\\d+)\\s")
    re1 := regexp.MustCompile("(falls asleep)")
    re2 := regexp.MustCompile("(wakes up)")
    for i := range events {
        match := re.FindStringSubmatch(events[i].Message)
        match1 := re1.FindStringSubmatch(events[i].Message)
        match2 := re2.FindStringSubmatch(events[i].Message)
        if len(match) > 1 {
            guardid, err = strconv.Atoi(match[1])
            if err != nil {
                log.Fatal("string to int conversion failed")
            }
        }
        if len(match1) > 1 {
            sleepdata[counter] = Sleep{Guard: guardid, From: events[i].Time}
        }
        if len(match2) > 1 {
            sleepdata[counter].To = events[i].Time
            counter++
        }
    }
    for i := range sleepdata {
        if len(roster[sleepdata[i].Guard]) == 0 {
            roster[sleepdata[i].Guard] = make([]int, 60)
        }
        for j := sleepdata[i].From.Minute(); j < sleepdata[i].To.Minute(); j++ {
            roster[sleepdata[i].Guard][j] += 1
        }
    }
    selected_guard := 0
    sleep := 0
    for i := range roster {
        sum := 0
        for _, num := range roster[i]{
            sum += num
        }
        if sum > sleep {
            selected_guard = i
            sleep = sum
        }
    }
    minute := 0
    num := 0
    for i, n := range roster[selected_guard] {
        if n > num {
            minute = i
            num = n
        }
    }
    fmt.Print("Guard: ", selected_guard, " Minute: ", minute, "\n")
    fmt.Println(selected_guard*minute)
}
