package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var testMap1 = map[string][]int64{"**": nil, "/": {3, 7}, "*": nil, "+": {1, 5, 9}, "-": nil}

func arithmicsInit(input cli.Args) {
	var parsedContent string = arithmicsParse(input.First())
	fmt.Println(parsedContent)
	arithmicsCompute(testMap1)
}

func arithmicsParse(value string) string {
	fmt.Println("haha beep boop")
	return value
}

func arithmicsCompute(equation map[string][]int64) int64 {
	fmt.Println("maths go brr")
	return 10101010
}
