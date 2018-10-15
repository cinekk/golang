package main

import "fmt"

func Sum(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	fmt.Println(Sum(10, 3))
	fmt.Println("Hello Dupa")
}