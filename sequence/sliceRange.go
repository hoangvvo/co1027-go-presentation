package main

import "fmt"

var arr = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
}
