package main

import "fmt"

func main() {
	var a string = "A"

	var (
		name     string = "name"
		age      int    = 24
		location string = "jakarta"
		// multiple variables in 1 line
		b, c string = "b", "c"
	)
	// alternative multiple variables in 1 line
	// b, c := "b", "c"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(location)
}
