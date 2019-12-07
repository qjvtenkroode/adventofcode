package main

import (
	"fmt"
	"strconv"
)

func main() {
	nrpwd := 0
	for i := 136760; i < 595730; i++ {
		str := strconv.Itoa(i)
		if checkdouble(str) {
			continue
		}
		if checknodecrease(str) {
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
