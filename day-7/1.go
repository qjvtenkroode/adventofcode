package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type step struct {
	Name    string
	Depends []string
	Next    []*step
	Prev    []*step
	Done    bool
}

func (s *step) Ready() bool {
	for _, d := range s.Prev {
		if !d.Done {
			return false
		}
	}
	return true
}

func getNext(list map[string]*step) (res *step) {
	for _, x := range list {
		if !x.Done && x.Ready() && (res == nil || x.Name < res.Name) {
			res = x
		}
	}
	return res
}

func finished(list map[string]*step) bool {
	for _, x := range list {
		if !x.Done {
			return false
		}
	}
	return true
}

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("error reading file: %s", err)
	}

	lines := strings.Split(string(f), "\n")
	lines = lines[:len(lines)-1]

	steps := make(map[string]*step)

	re := regexp.MustCompile("Step\\s(\\w)\\smust be finished before step (\\w)")

	for _, l := range lines {
		match := re.FindStringSubmatch(l)
		if steps[match[2]] != nil {
			steps[match[2]].Depends = append(steps[match[2]].Depends, match[1])
		} else {
			s := step{Name: match[2], Depends: []string{match[1]}}
			steps[match[2]] = &s
		}
	}

	for _, x := range steps {
		for _, dep := range x.Depends {
			prev := steps[dep]
			if prev != nil {
				prev.Next = append(prev.Next, x)
				x.Prev = append(x.Prev, prev)
			} else {
				s := step{Name: dep, Next: []*step{x}}
				steps[dep] = &s
			}
		}
	}

	var n *step
	pattern := ""
	for {
		n = getNext(steps)
		n.Done = true
		pattern += n.Name
		if finished(steps) {
			fmt.Println(pattern)
			return
		}
	}
}
