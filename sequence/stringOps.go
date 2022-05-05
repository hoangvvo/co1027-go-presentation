package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Viet Nam"
	fmt.Printf(`== Get length ==
lenght of Viet Nam: %d`, len(str))

	fmt.Printf(`
== Concatenate =="
%s`, "Ja"+"pan")

	strCat := "Cat says Meow"
	fmt.Printf(`
== Get substring ==
strCat[1:5] = %s
`, strCat[1:5])

	fmt.Println("== Search for substring ==")
	japan := "Japan"
	fmt.Printf("Is abc in %s? %t\n", japan, strings.Contains(japan, "abc"))
	fmt.Printf("Non-overlapping instances of na in banana: %d\n", strings.Count("Banana", "na"))
	vn := "Viet Nam"
	fmt.Printf("Index of first 'Nam' substring in %s: %d\n", vn, strings.Index(vn, "Nam"))
	can := "Canada"
	fmt.Printf("Index of last 'a' substring in %s: %d\n", can, strings.Index(can, "a"))

	fmt.Println("== Replace substring ==")
	bk := "Bách Khoa Bách Khoa Bách Khoa"
	bk = strings.Replace(bk, "Khoa", "Trúng", 2) // Use -1 to replace all
	fmt.Printf("Replace first two “Khoa” with “Trúng”: %s\n", bk)

	fmt.Println("== Lowercase/Uppercase ==")
	up := "Uppercase"
	lo := "Lowecase"
	fmt.Printf("strings.ToUpper(%s) = %s, strings.ToLower(%s) = %s\n",
		up, strings.ToUpper(up), lo, strings.ToLower(lo))

	fmt.Println("== Split/Join ==")
	fmt.Printf("Split “1|2|3|4|5” into slice by “|”: %v\n", strings.Split("1|2|3|4|5", "|"))
	fmt.Printf("Join slice by “&”: %s\n", strings.Join([]string{"a", "b", "c"}, "&"))
}
