package main

import "fmt"

func main() {
	a := 4
	b := 5
	c := 6
	a, b, c = c, b, a
	fmt.Println(a, b, c)
}
