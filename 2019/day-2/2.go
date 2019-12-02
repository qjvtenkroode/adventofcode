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

	var x, y int
	for x = 0; x < 99; x++ {
		for y = 0; y < 99; y++ {
			o := intcode(prog, x, y)
			if o == 19690720 {
				fmt.Println("found")
				fmt.Printf("x:%d y:%d answer:%d\n", x, y, 100*x+y)
				return
			}
		}
	}
}

func intcode(prog []int, x, y int) int {
	var a, b, c, d int

	mem := make([]int, len(prog))
	copy(mem, prog)

	mem[1] = x
	mem[2] = y

	for i := 0; i < len(mem); i += 4 {
		if mem[i] == 1 {
			a = mem[i+1]
			b = mem[i+2]
			c = mem[i+3]
			d = mem[a] + mem[b]
		}
		if mem[i] == 2 {
			a = mem[i+1]
			b = mem[i+2]
			c = mem[i+3]
			d = mem[a] * mem[b]
		}
		if mem[i] == 99 {
			break
		}
		mem[c] = d
	}
	return mem[0]
}
