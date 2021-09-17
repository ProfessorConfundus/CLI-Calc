package main

import (
	"fmt"
)

var testMap1 = map[string][]int64{"**": nil, "/": {3, 7}, "*": nil, "+": {1, 5, 9}, "-": nil}

func arithmicsInit() {
	var parsedContent string = arithmicsParse("beep")
	fmt.Println(parsedContent)
	arithmicsCompute(testMap1)
}

func arithmicsParse(value string) string {
	fmt.Println("haha beep boop")
	return value
}

func arithmicsCompute(equasion map[string][]int64) int64 {
	fmt.Println("maths go brr")
	return 10101010
}
