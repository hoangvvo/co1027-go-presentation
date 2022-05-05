package main

import "fmt"

type Vertex struct {
	Name string
	X    int
	Y    int
}

func main() {
	var v0 Vertex // Zeroth struct with zeroth fields
	fmt.Printf("%+v\n", v0)

	v1 := Vertex{"a", 1, 2} // Implicit
	fmt.Printf("%+v\n", v1)

	v2 := Vertex{Name: "b", Y: 4, X: 2} // Explicit
	fmt.Printf("%+v\n", v2)
}
