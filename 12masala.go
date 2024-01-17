package main

import "fmt"

func main() {
	a := 396
	r1 := a / 100
	r2 := (a % 100) / 10
	r3 := a % 10
	fmt.Println(r2*100 + r1*10 + r3)

}
