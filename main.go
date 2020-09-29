package main

import "fmt"

func main() {
	var base int64
	var altura int64

	fmt.Scanln(&base)
	fmt.Scanln(&altura)

	area := (base * altura) / 2

	fmt.Println(area)
}
