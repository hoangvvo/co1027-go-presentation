package main

import (
	"fmt"
	"runtime"
)

func whoCreatedTheoryOfNaturalSelection() string {
	return "darwin" // code name core system run Apple OS
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case whoCreatedTheoryOfNaturalSelection():
		fmt.Println("🍎 OS X.")
	case "linux":
		fmt.Println("🐧 Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
