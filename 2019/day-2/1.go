package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	defer f.Close()

	prog := make([]int, 129)

	scanner := bufio.NewScanner(f)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' || data[i] == '\n' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, nil, nil
	})

	i := 0
	for scanner.Scan() {
		d, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("error converting strings to ints: %s", err)
		}
		prog[i] = d
		i++
	}

	var a, b, c, d int

	for i = 0; i < len(prog); i += 4 {
		if prog[i] == 1 {
			a = prog[i+1]
			b = prog[i+2]
			c = prog[i+3]
			d = prog[a] + prog[b]
		}
		if prog[i] == 2 {
			a = prog[i+1]
			b = prog[i+2]
			c = prog[i+3]
			d = prog[a] * prog[b]
		}
		if prog[i] == 99 {
			fmt.Println(prog[0])
			return
		}
		prog[c] = d
	}
}
