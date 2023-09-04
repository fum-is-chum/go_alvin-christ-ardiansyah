package main

import (
	"fmt"
	"time"
)

func kelipatan(x int) {
	multiplier := 1
	for true {
		fmt.Println(x * multiplier)
		multiplier++
	}
}

func main() {
	go kelipatan(3)
	time.Sleep(3 * time.Second)
}