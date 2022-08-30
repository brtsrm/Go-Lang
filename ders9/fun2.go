package main

import "fmt"

func main() {

	fmt.Println(33)
	numberTerms, sum := add2(1, 2)
	fmt.Println(numberTerms, sum)

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add2(terms ...int) (numberTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numberTerms = len(terms)
	return
}
