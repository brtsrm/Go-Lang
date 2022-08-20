package main

import "fmt"

func main() {

	a := 10

	b := 20

	total := a + b
	total = total - 5
	total *= total
	total = a / b
	total = 10 / 22
	total = a % b

	total++
	total--

	fmt.Println(total)

}
