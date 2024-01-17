package main

import "fmt"

func main() {
	a := 674
	r1 := a / 100
	r2 := (a / 10) % 10
	r3 := (a % 10)
	fmt.Println(r3*100 + r2*10 + r1)
}
