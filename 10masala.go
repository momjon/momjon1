package main

import "fmt"

func main() {
	a := 467
	r1 := a / 100
	r2 := (a / 10) % 10
	r3 := a % 10
	fmt.Println(r2*100 + r1*10 + r3)
}
