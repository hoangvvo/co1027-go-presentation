package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	fmt.Println("Enter your first name and age in the format\n[name] [age]")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Hello", name, "| you are", age, "years old")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	fmt.Print("You live in " + city)
}
