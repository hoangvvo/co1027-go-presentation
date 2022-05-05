package main

import (
	"fmt"
	"math"
)

func powWithLimit(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		powWithLimit(3, 2, 10),
		powWithLimit(3, 3, 20),
	)
}
