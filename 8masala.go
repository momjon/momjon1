package main

import "fmt"

func main() {
	a := 695
	r1 := a / 100
	r2 := (a % 100) / 10
	r3 := a % 10
	fmt.Println(r1 + r2 + r3)

}
