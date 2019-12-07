package main

import (
	"fmt"
	"strconv"
)

func main() {
	nrpwd := 0
	for i := 136760; i < 595731; i++ {
		str := strconv.Itoa(i)
		if checkdouble(str) {
			continue
		}
		if checknodecrease(str) {
			continue
		}
		if checklargergrp(str) {
			continue
		}
		nrpwd++
	}
	fmt.Println(nrpwd)
}

func checkdouble(pwd string) bool {
	for k := 1; k < len(pwd); k++ {
		if pwd[k-1] == pwd[k] {
			return false
		}
	}
	return true
}

func checknodecrease(pwd string) bool {
	for j := 1; j < len(pwd); j++ {
		if pwd[j-1] > pwd[j] {
			return true
		}
	}
	return false
}

func checklargergrp(pwd string) bool {
	if pwd[3] != pwd[4] && pwd[4] == pwd[5] {
		return false
	}
	if pwd[2] != pwd[3] && pwd[3] == pwd[4] && pwd[4] != pwd[5] {
		return false
	}
	if pwd[1] != pwd[2] && pwd[2] == pwd[3] && pwd[3] != pwd[4] {
		return false
	}
	if pwd[0] != pwd[1] && pwd[1] == pwd[2] && pwd[2] != pwd[3] {
		return false
	}
	if pwd[0] == pwd[1] && pwd[1] != pwd[2] {
		return false
	}

	return true
}
