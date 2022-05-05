package main

import (
	"fmt"
)

type Student struct {
	Name      string
	ClassYear int
	GPA       float32
}

func main() {
	var name = "Hoang"
	var age = 19
	fmt.Printf("Hello, my name is %s. I am %d years old\n", name, age)
	var val = 4
	fmt.Printf("Binary: %b\\%b\n", val, 5)
	fmt.Printf("Pad with spaces, left justified: %-4d|\n", 15)
	fmt.Printf("Pad with spaces: %04d\n", 15)
	fmt.Printf("Round to precision 2: %.2f\n", 12.3456)
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Print array: %v\n", arr)
	s := Student{Name: "Hoang", ClassYear: 2021, GPA: 5.0}
	fmt.Printf("Student Info: %+v\n", s)
	var hello = "Annyeonghaseyo"
	str := fmt.Sprintf("%s, %s!", hello, name) // set to string
	fmt.Println(str)
}
