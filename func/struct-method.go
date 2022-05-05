package main

import "fmt"

type Vertex struct {
	Name string
	X    int
	Y    int
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) GetX() int {
	return v.X
}

func main() {
	v := Vertex{Name: "A", X: 2, Y: 4}
	v.Scale(5)
	fmt.Println(v.GetX())
}
