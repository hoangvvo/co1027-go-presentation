package main

import "fmt"

func passInt(n int) {
	n = 10
}

func passStr(s string) {
	s += " Khoa"
}

func passArr(arr [5]int) {
	arr[1] = 100
}

type Student struct {
	Name  string
	Major string
}

func passStruct(s Student) {
	s.Major = "Chemistry"
}

func main() {
	n := 5
	passInt(n)
	fmt.Println(n)

	s := "Bach"
	passStr(s)
	fmt.Println(s)

	arr := [5]int{1, 2, 3, 4, 5}
	passArr(arr)
	fmt.Println(arr)

	st := Student{Name: "Hoang", Major: "Computer Science"}
	passStruct(st)
	fmt.Printf("%+v", st)
}
