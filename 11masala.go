package main

import "fmt"

func main() {
	a := 378
	r1 := a / 100
	r2 := (a / 10) % 10
	r3 := a % 10
	fmt.Println(r1*100 + r3*10 + r2)
}
