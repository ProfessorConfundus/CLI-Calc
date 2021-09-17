package main

import (
	"fmt"
)

func arithmicsInit() {
	var parsedContent string = arithmicsParse("beep")
	fmt.Println(parsedContent)
}

func arithmicsParse(value string) string {
	fmt.Println("haha beep boop")
	return value
}
