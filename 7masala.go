package main

import "fmt"

func main() {
	a := 268

	r2 := (a / 10) % 10
	r3 := a % 10
	fmt.Println(r3, r2)
}
