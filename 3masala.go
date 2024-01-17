package main

import "fmt"

func main() {
	a := 34
	r1 := a / 10
	var r2 = a % 10

	fmt.Println(r2*10 + r1)
}
