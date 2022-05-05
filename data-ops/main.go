package main

import "fmt"

func main() {
	// Arithmetic Operators
	println("Arithmetic Operators")
	var width = 4.5
	var height = 5.5
	var area = width * height
	fmt.Println("area = ", area)

	// Relational Operators
	println("Relational Operators")
	println("area > 10 is ", area > 10)

	// Assignment Operators
	println("Assignment Operators")
	num := 10 // note: assignment shorthand
	num = num + 2
	println("num = ", num)
	num = 4
	println("num = ", num)

	// Logical Operators
	println("Logical Operators")
	isBetween6And9 := num >= 6 && num <= 9
	println("6 <= num <= 9", isBetween6And9)

	// Bitwise Operators
	println("Bitwise Operators")
	fmt.Printf("0011 & 0101: %b\n", 0011&0101)
	//          0001

	// Misc Operators
	println("Misc Operators")
	var val int = 0
	println("val = ", val)
	var valAddr *int = &val
	println("valAddr = ", valAddr)
	*valAddr = 10
	println("val = ", val)
}
